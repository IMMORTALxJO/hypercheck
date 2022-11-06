package tcp

import (
	"net"
	"time"

	log "github.com/sirupsen/logrus"
)

type tcpWrapper struct {
	Address string
	latency uint64
}

func (w *tcpWrapper) getOnline() bool {
	netd := net.Dialer{Timeout: time.Duration(1) * time.Second}
	startTime := time.Now().Unix()
	conn, err := netd.Dial("tcp", w.Address)
	w.latency = uint64(time.Now().Unix() - startTime)
	if err != nil {
		log.Error(err)
		return false
	}
	defer conn.Close()
	return true
}

func (w *tcpWrapper) getLatency() uint64 {
	w.getOnline()
	return w.latency
}

var tcpCache = map[string]*tcpWrapper{}

func getTcpWrapper(address string) *tcpWrapper {
	_, ok := tcpCache[address]
	if !ok {
		tcpCache[address] = &tcpWrapper{
			Address: address,
		}
	} else {
		log.Debugf("got tcpWrapper for %s from cache", address)
	}
	return tcpCache[address]
}
