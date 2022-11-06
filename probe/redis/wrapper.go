package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

type redisWrapper struct {
	Address string
	conn    redis.Conn
	err     error
}

func (w *redisWrapper) getPing() bool {
	conn, err := w.getConn()
	if err != nil {
		log.Error(err)
		return false
	}
	answer, err := redis.DoWithTimeout(conn, time.Second, "ping")
	if answer != "PONG" {
		log.Errorf("redis %s ping answer is not PONG", w.Address)
		return false
	}
	return true
}

func (w *redisWrapper) getConn() (redis.Conn, error) {
	if w.conn == nil {
		w.load()
	}
	return w.conn, w.err
}

func (w *redisWrapper) load() {
	w.conn, w.err = redis.Dial("tcp", w.Address)
	log.Debugf("redis.conn loaded %s", w.Address)
}

var redisCache = map[string]*redisWrapper{}

func getRedisWrapper(address string) *redisWrapper {
	_, ok := redisCache[address]
	if !ok {
		redisCache[address] = &redisWrapper{
			Address: address,
		}
	} else {
		log.Debugf("got redisWrapper for %s from cache", address)
	}
	return redisCache[address]
}
