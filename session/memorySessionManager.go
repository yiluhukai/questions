package session

import (
	"github.com/satori/go.uuid"
	"sync"
)

type MemorySessionManager struct {
	sessions map[string]Session
	rw       sync.RWMutex
}

func (sm *MemorySessionManager) Init(addr string, options ...string) error {
	return nil
}

func NewMemorySessionManager() SessionManager {
	return &MemorySessionManager{
		sessions: make(map[string]Session, 1024),
	}
}

func (sm *MemorySessionManager) Get(sessionId string) (session Session, err error) {
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	session, ok := sm.sessions[sessionId]
	if !ok {
		err = ErrorSessionNotExist
		return
	}
	return
}

func (sm *MemorySessionManager) CreateSession() (session Session, err error) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	id, err := uuid.NewV4()
	if err != nil {
		return
	}
	sessionId := id.String()
	session = NewMemorySession(sessionId)
	sm.sessions[sessionId] = session
	return
}
