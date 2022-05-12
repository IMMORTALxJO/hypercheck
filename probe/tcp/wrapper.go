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

func (w *tcpWrapper) GetOnline() bool {
	netd := net.Dialer{Timeout: time.Duration(1) * time.Second}
	startTime := time.Now().UnixNano()
	conn, err := netd.Dial("tcp", w.Address)
	w.latency = uint64(time.Now().UnixNano() - startTime)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func (w *tcpWrapper) GetLatency() uint64 {
	w.GetOnline()
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
