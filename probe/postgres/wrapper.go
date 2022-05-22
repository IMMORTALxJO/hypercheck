package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"

	log "github.com/sirupsen/logrus"
)

type postgresWrapper struct {
	Address string
	err     error
	conn    *pgx.Conn
}

func (w *postgresWrapper) GetOnline() bool {
	err := w.getError()
	if err != nil {
		log.Error(err)
	}
	return err == nil
}

func (w *postgresWrapper) getError() error {
	if w.conn == nil {
		w.load()
	}
	return w.err
}

func (w *postgresWrapper) getConn() *pgx.Conn {
	if w.conn == nil {
		w.load()
	}
	return w.conn
}

func (w *postgresWrapper) load() {
	w.conn, w.err = pgx.Connect(context.Background(), w.Address)
	log.Debugf("db.conn loaded %s", w.Address)
}

var postgresCache = map[string]*postgresWrapper{}

func getPostgresWrapper(address string) *postgresWrapper {
	_, ok := postgresCache[address]
	if !ok {
		postgresCache[address] = &postgresWrapper{
			Address: address,
		}
	} else {
		log.Debugf("got postgresWrapper for %s from cache", address)
	}
	return postgresCache[address]
}
