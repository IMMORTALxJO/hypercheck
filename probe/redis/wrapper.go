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

func (w *redisWrapper) GetOnline() bool {
	err := w.getError()
	if err != nil {
		log.Error(err)
	}
	return err == nil
}

func (w *redisWrapper) GetPing() bool {
	err := w.getError()
	var answer interface{}
	if err == nil {
		answer, err = redis.DoWithTimeout(w.getConn(), time.Second, "ping")
		if answer != "PONG" {
			log.Errorf("redis %s ping answer is not PONG", w.Address)
			return false
		}
	}
	if err != nil {
		log.Error(err)
	}
	return err == nil
}

func (w *redisWrapper) getError() error {
	if w.conn == nil {
		w.load()
	}
	return w.err
}

func (w *redisWrapper) getConn() redis.Conn {
	if w.conn == nil {
		w.load()
	}
	return w.conn
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
