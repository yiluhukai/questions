package session

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RedisSessionManager struct {
	pool     *redis.Pool
	rw       sync.RWMutex
	password string
	addr     string
	sessions map[string]Session
}

func NewRedisSessionManager() SessionManager {
	return &RedisSessionManager{
		sessions: make(map[string]Session),
	}
}

//初始化一个pool
func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			fmt.Println("dialog redis to test")
			if err != nil {

				return nil, err
			}

			// if _, err := c.Do("AUTH", "foobared"); err != nil {
			// 	c.Close()
			// 	return nil, err
			// }
			return c, err
		},
		//connect存在时间大于一分钟，连接一次redis服务器

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

// 初始化redis 连接池
func (rs *RedisSessionManager) Init(addr string, options ...string) (err error) {

	if len(options) > 0 {
		rs.password = options[0]
	}
	rs.pool = newPool(addr, rs.password)
	_, err = rs.pool.Dial()
	if err != nil {
		return
	}
	rs.addr = addr
	return
}

func (rs *RedisSessionManager) Get(sessionId string) (session Session, err error) {
	rs.rw.RLock()
	defer rs.rw.RUnlock()
	session, ok := rs.sessions[sessionId]
	if !ok {
		err = ErrorSessionKeyNotExist
		return
	}
	return
}

func (rs *RedisSessionManager) CreateSession() (session Session, err error) {
	rs.rw.Lock()
	defer rs.rw.Unlock()
	id, err := uuid.NewV4()
	if err != nil {
		return
	}
	sessionId := id.String()
	session = NewRedisSession(sessionId, rs.pool)
	rs.sessions[sessionId] = session
	return
}
