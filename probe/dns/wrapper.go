package dns

import (
	"net"

	log "github.com/sirupsen/logrus"
)

type dnsWrapper struct {
	Domain    string
	rec_A     []string
	rec_NS    []string
	rec_MX    []string
	rec_TXT   []string
	rec_CNAME string
}

func (w *dnsWrapper) GetA() []string {
	if w.rec_A == nil {
		w.rec_A = []string{}
		lookup, err := net.LookupIP(w.Domain)
		if err != nil {
			log.Error(err)
		}
		for _, ip := range lookup {
			w.rec_A = append(w.rec_A, ip.String())
			log.Debugf("%s resolved to %v", w.Domain, ip)
		}
	}
	return w.rec_A
}

func (w *dnsWrapper) GetCNAME() string {
	var err error
	if w.rec_CNAME == "" {
		w.rec_CNAME, err = net.LookupCNAME(w.Domain)
		if err != nil {
			log.Error(err)
		}
		log.Debugf("%s CNAME is %v", w.Domain, w.rec_CNAME)
	}
	return w.rec_CNAME
}

func (w *dnsWrapper) GetNS() []string {
	if w.rec_NS == nil {
		lookup, err := net.LookupNS(w.Domain)
		if err != nil {
			log.Error(err)
		}
		for _, ip := range lookup {
			w.rec_NS = append(w.rec_NS, ip.Host)
			log.Debugf("%s NS to %s", w.Domain, ip.Host)
		}
	}
	return w.rec_NS
}

func (w *dnsWrapper) GetMX() []string {
	if w.rec_MX == nil {
		lookup, err := net.LookupMX(w.Domain)
		if err != nil {
			log.Error(err)
		}
		for _, ip := range lookup {
			w.rec_MX = append(w.rec_MX, ip.Host)
			log.Debugf("%s MX to %s", w.Domain, ip.Host)
		}
	}
	return w.rec_MX
}

func (w *dnsWrapper) GetTXT() []string {
	var err error
	if w.rec_MX == nil {
		w.rec_TXT, err = net.LookupTXT(w.Domain)
		if err != nil {
			log.Error(err)
		}
		log.Debugf("%s TXT to %v", w.Domain, w.rec_TXT)
	}
	return w.rec_TXT
}

func (w *dnsWrapper) GetOnline() bool {
	return len(w.GetA()) > 0
}

var dnsCache = map[string]*dnsWrapper{}

func getDnsWrapper(domain string) *dnsWrapper {
	_, ok := dnsCache[domain]
	if !ok {
		dnsCache[domain] = &dnsWrapper{
			Domain: domain,
		}
	} else {
		log.Debugf("got dnsWrapper for %s from cache", domain)
	}
	return dnsCache[domain]
}
