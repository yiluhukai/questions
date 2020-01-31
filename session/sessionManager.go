package session

type SessionManager interface {
	Init(addr string, options ...string) error
	Get(sessionId string) (Session, error)
	CreateSession() (session Session, err error)
}
