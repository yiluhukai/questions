package session

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"sync"
)

const (
	SessionFlagNone = iota
	SessionFlagModify
	SessionFlagLoad
)

type RedisSession struct {
	sessionId string
	pool      *redis.Pool
	rw        sync.RWMutex
	// 存放session的键值对
	maps map[string]interface{}
	flag int
}

func NewRedisSession(sessionId string, pool *redis.Pool) Session {
	return &RedisSession{
		sessionId: sessionId,
		pool:      pool,
		maps:      make(map[string]interface{}, 8),
		flag:      SessionFlagNone,
	}
}

func (rs *RedisSession) Set(key string, value interface{}) (err error) {
	rs.rw.Lock()
	defer rs.rw.Unlock()
	rs.maps[key] = value
	rs.flag = SessionFlagModify
	return
}

// 从redis中取出session
func (rs *RedisSession) loadFromRedis() (err error) {
	conn := rs.pool.Get()
	reply, err := conn.Do("GET")
	replyStr, err := redis.String(reply, err)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(replyStr), &rs.maps)
	return
}

func (rs *RedisSession) Get(key string) (session interface{}, err error) {
	rs.rw.RLock()
	defer rs.rw.RUnlock()
	// 实现懒加载
	if rs.flag == SessionFlagNone {
		err = rs.loadFromRedis()
		if err != nil {
			return
		}
	}
	session, ok := rs.maps[key]
	if !ok {
		err = ErrorSessionKeyNotExist
		return
	}
	return
}

// 删除session中的一项
func (rs *RedisSession) Del(key string) (err error) {
	rs.rw.Lock()
	defer rs.rw.Lock()
	delete(rs.maps, key)
	rs.flag = SessionFlagModify
	return
}

func (rs *RedisSession) Save() (err error) {
	rs.rw.Lock()
	defer rs.rw.Unlock()

	if rs.flag != SessionFlagModify {
		return
	}
	conn := rs.pool.Get()
	session, err := json.Marshal(rs.maps)
	if err != nil {
		return
	}
	_, err = conn.Do("SET", rs.sessionId, string(session))
	if err != nil {
		rs.flag = SessionFlagNone
	}
	return
}

func (ms *RedisSession) IsModify() bool {
	return ms.flag == SessionFlagModify
}

// return sessionID
func (ms *RedisSession) Id() string {
	return ms.sessionId
}
