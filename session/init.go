package session

import "fmt"

import "errors"

var (
	sessionMgr SessionManager
)

//provider:
//1. memory， 返回一个内存的session管理类
//2. redis, 返回一个redis的session管理类
func Init(provider string, addr string, options ...string) (err error) {

	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionManager()
	case "redis":
		sessionMgr = NewRedisSessionManager()
	default:
		err = fmt.Errorf("not support")
		return
	}
	err = sessionMgr.Init(addr, options...)
	return
}

func Get(sessionId string) (session Session, err error) {
	if sessionMgr == nil {
		err = errors.New("sessionManager doesn't init")
		return
	}
	return sessionMgr.Get(sessionId)
}

func CreateSession() (session Session, err error) {
	if sessionMgr == nil {
		err = errors.New("sessionManager doesn't init")
		return
	}
	return sessionMgr.CreateSession()
}
