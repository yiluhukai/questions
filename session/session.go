package session

// 定义一个session接口、内存session和mysql session 或者redis session都会实现这个接口

type Session interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Del(key string) error
	Save() error
	IsModify() bool
	Id() string
}
