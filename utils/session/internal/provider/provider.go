package provider

import "session/pkg"

var provides = make(map[string]Provider)

type Provider interface {
	SessionInit(sid string) (pkg.Session, error) // 初始化session
	SessionRead(sid string) (pkg.Session, error) // 读取session
	SessionDestroy(sid string) error             // 注销session
	SessionGC(maxLifeTime int64)                 // 超时回收session
	SessionUpdate(sid string) error              // 更新session超时时间
}

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider" + name)
	}
	provides[name] = provider
}

func GetProvider(name string) (Provider, bool) {
	provider, ok := provides[name]
	return provider, ok
}
