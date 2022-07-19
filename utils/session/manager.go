package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"sync"
)

type Manager struct {
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

func NewManager(provideType Type, maxLifeTime int64) (*Manager, error) {
	provide, ok := GetProvider(provideType)
	if !ok {
		return nil, fmt.Errorf("session: unknown provider %q (forgotten import?)", provideType)
	}
	return &Manager{provider: provide, maxLifeTime: maxLifeTime}, nil
}

// 生成 unique session id
func (m *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// SessionStart 如果 sid 为""，则创建新 session；如果 session 不为""，则获取该 session
func (m *Manager) SessionStart() (s Session, err error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	sid := m.sessionId()
	s, err = m.provider.SessionInit(sid)
	return
}

func (m *Manager) SessionGet(sid string) (s Session, err error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	s, err = m.provider.SessionRead(sid)
	return s, err
}

// SessionDestroy 销毁 session，主动注销
func (m *Manager) SessionDestroy(sid string) (err error) {
	if sid == "" {
		return
	} else {
		m.lock.Lock()
		defer m.lock.Unlock()
		err = m.provider.SessionDestroy(sid)
	}
	return
}

// GC 超时回收 session
func (m *Manager) GC() {
	m.lock.Lock()
	defer m.lock.Lock()
	m.provider.SessionGC(m.maxLifeTime)
}
