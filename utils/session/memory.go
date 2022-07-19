/*
	The memory package practice a memory provider of session.
	memory 实现了一个内存存储的session provider。
	该provider由一个map和一个双向单链表构成，map用于快速访问，双向单链表用于快速gc(越接近超时的session排在链尾，创建、更新和使用某一session
	时，将其放置到链首，每次gc时，判断链尾元素是否过期)。
*/
package session

import (
	"container/list"
	"sync"
	"time"
)

var pder = &memProvider{list: list.New()}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	Register(Memory, pder)
}

// memSession 内存session
type memSession struct {
	sid          string                      // unique session id
	timeAccessed time.Time                   // last access time
	value        map[interface{}]interface{} // session value
}

func (s *memSession) Set(key, value interface{}) error {
	s.value[key] = value
	return pder.SessionUpdate(s.sid)
}

func (s *memSession) Get(key interface{}) interface{} {
	_ = pder.SessionUpdate(s.sid)
	if v, ok := s.value[key]; ok {
		return v
	} else {
		return nil
	}
}

func (s *memSession) Delete(key interface{}) error {
	delete(s.value, key)
	return pder.SessionUpdate(s.sid)
}

func (s *memSession) SessionID() string {
	return s.sid
}

// memProvider 内存存储session
type memProvider struct {
	lock     sync.Mutex               // lock
	sessions map[string]*list.Element // session 存储
	list     *list.List               // gc
}

// SessionInit 初始化session, 将新的session放置在双向链表的头部
func (p *memProvider) SessionInit(sid string) (Session, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newSess := &memSession{
		sid:          sid,
		timeAccessed: time.Now(),
		value:        v,
	}
	element := p.list.PushFront(newSess)
	p.sessions[sid] = element
	return newSess, nil
}

// SessionRead 获取session
func (p *memProvider) SessionRead(sid string) (Session, error) {
	if element, ok := p.sessions[sid]; ok {
		return element.Value.(*memSession), nil
	} else {
		return nil, nil
	}
}

// SessionDestroy 注销session
func (p *memProvider) SessionDestroy(sid string) error {
	if element, ok := p.sessions[sid]; ok {
		delete(p.sessions, sid)
		p.list.Remove(element)
		return nil
	}
	return nil
}

// SessionGC 回收链尾超时session
func (p *memProvider) SessionGC(maxLifeTime int64) {
	p.lock.Lock()
	defer p.lock.Unlock()
	for {
		element := p.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*memSession).timeAccessed.Unix() + maxLifeTime) < time.Now().Unix() {
			p.list.Remove(element)
			delete(p.sessions, element.Value.(*memSession).sid)
		} else {
			break
		}
	}
}

// SessionUpdate 更新session过期时间
func (p *memProvider) SessionUpdate(sid string) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if element, ok := p.sessions[sid]; ok {
		element.Value.(*memSession).timeAccessed = time.Now()
		p.list.MoveToFront(element)
		return nil
	}
	return nil
}
