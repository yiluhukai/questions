package account

import "yiluhuakai/questions/session"

// InitSession is used to init a SessionManager
func InitSession(provider string, addrss string, options ...string) (err error) {
	return session.Init(provider, addrss, options...)
}
