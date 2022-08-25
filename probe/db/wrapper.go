package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"net/url"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

type dbWrapper struct {
	Address string
	err     error
	parsed  *url.URL
	conn    *sql.DB
}

func (w *dbWrapper) GetOnline() bool {
	err := w.getError()
	if err != nil {
		log.Error(err)
	}
	return err == nil
}

func (w *dbWrapper) getError() error {
	if w.conn == nil {
		w.load()
	}
	return w.err
}

func (w *dbWrapper) getScheme() string {
	w.parsed, w.err = url.Parse(w.Address)
	return w.parsed.Scheme
}

func (w *dbWrapper) composeAddress() string {
	if w.getScheme() == "mysql" {
		return fmt.Sprintf("%s@tcp(%s)%s?%s", w.parsed.User, w.parsed.Host, w.parsed.Path, w.parsed.RawQuery)
	}
	return w.Address
}

func (w *dbWrapper) load() {
	w.conn, w.err = sql.Open(w.getScheme(), w.composeAddress())
	if w.err == nil {
		ctxBG := context.Background()
		ctxConnTimeout, cancel := context.WithTimeout(ctxBG, 1*time.Second)
		defer cancel()
		w.err = w.conn.PingContext(ctxConnTimeout)
	}
	log.Debugf("db.conn loaded %s", w.Address)
}

var dbCache = map[string]*dbWrapper{}

func getDbWrapper(address string) *dbWrapper {
	_, ok := dbCache[address]
	if !ok {
		dbCache[address] = &dbWrapper{
			Address: address,
		}
	} else {
		log.Debugf("got dbWrapper for %s from cache", address)
	}
	return dbCache[address]
}
