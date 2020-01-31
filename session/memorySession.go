package session

import "sync"

type MemorySession struct {
	data map[string]interface{}
	id   string
	rw   sync.RWMutex
	flag int
}

func NewMemorySession(key string) Session {
	return &MemorySession{
		data: map[string]interface{}{},
		id:   key,
		rw:   sync.RWMutex{},
		flag: SessionFlagNone,
	}
}

func (ms *MemorySession) Set(key string, session interface{}) (err error) {
	ms.rw.Lock()
	defer ms.rw.Unlock()
	ms.data[key] = session
	ms.flag = SessionFlagModify
	return
}

func (ms *MemorySession) Get(key string) (session interface{}, err error) {
	ms.rw.RLock()
	defer ms.rw.RLock()
	session, ok := ms.data[key]
	if !ok {
		err = ErrorSessionKeyNotExist
		return
	}
	return
}

func (ms *MemorySession) Del(key string) (err error) {
	ms.rw.Lock()
	defer ms.rw.Unlock()
	delete(ms.data, key)
	ms.flag = SessionFlagModify
	return
}

func (ms *MemorySession) Save() (err error) {
	ms.flag = SessionFlagNone
	return
}

func (ms *MemorySession) IsModify() bool {
	return ms.flag == SessionFlagModify
}

// return sessionID
func (ms *MemorySession) Id() string {
	return ms.id
}
