package session

type Session interface {
	Set(key, value interface{}) error // 设置session value
	Get(key interface{}) interface{}  // 获取session value
	Delete(key interface{}) error     // 删除session value
	SessionID() string                // 返回当前session id
}
