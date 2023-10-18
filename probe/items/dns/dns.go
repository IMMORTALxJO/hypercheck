package dns

import (
	"fmt"
	t "hypercheck/probe/types"
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
)

const tableName = "dns"
const checkIsUpMessage = "is exists"
const notCheckedMessage = "not checked"

type Item struct {
	Address     string `gorm:"primaryKey"`
	RType       string
	Checked     bool
	Failed      bool
	LatencyNano uint64
	Message     string
	Value       string
}

func (Item) TableName() string {
	return tableName
}

func (i Item) IsFailed() bool {
	return i.Failed
}

func (i Item) GetMessage() string {
	return fmt.Sprintf("%s:%s - %s", i.Address, i.RType, i.Message)
}

func (i *Item) Enrich() {
	i.Checked = true
	i.Message = checkIsUpMessage
	i.Failed = false
	if i.RType == "A" {
		lookup, err := net.LookupIP(i.Address)
		if err != nil {
			i.Message = err.Error()
			log.Debugf("DNS for %s failed: %s", i.Address, err)
			i.Failed = true
			return
		}
		var ips []string
		for _, ip := range lookup {
			ips = append(ips, ip.String())
		}
		i.Value = strings.Join(ips, "\n")
		log.Debugf("DNS for %s:%s resolved to %v", i.Address, i.RType, lookup)
	} else if i.RType == "CNAME" {
		lookup, err := net.LookupCNAME(i.Address)
		if err != nil {
			i.Message = err.Error()
			log.Debugf("DNS for %s failed: %s", i.Address, err)
			i.Failed = true
			return
		}
		i.Value = lookup
		log.Debugf("DNS for %s:%s resolved to %v", i.Address, i.RType, lookup)
	} else if i.RType == "NS" {
		lookup, err := net.LookupNS(i.Address)
		if err != nil {
			i.Message = err.Error()
			log.Debugf("DNS for %s failed: %s", i.Address, err)
			i.Failed = true
			return
		}
		var ns []string
		for _, ip := range lookup {
			ns = append(ns, ip.Host)
		}
		i.Value = strings.Join(ns, "\n")
		log.Debugf("DNS for %s:%s resolved to %v", i.Address, i.RType, lookup)
	} else if i.RType == "MX" {
		lookup, err := net.LookupMX(i.Address)
		if err != nil {
			i.Message = err.Error()
			log.Debugf("DNS for %s failed: %s", i.Address, err)
			i.Failed = true
			return
		}
		var mx []string
		for _, ip := range lookup {
			mx = append(mx, ip.Host)
		}
		i.Value = strings.Join(mx, "\n")
		log.Debugf("DNS for %s:%s resolved to %v", i.Address, i.RType, lookup)
	} else if i.RType == "TXT" {
		lookup, err := net.LookupTXT(i.Address)
		if err != nil {
			i.Message = err.Error()
			log.Debugf("DNS for %s failed: %s", i.Address, err)
			i.Failed = true
			return
		}
		i.Value = strings.Join(lookup, "\n")
		log.Debugf("DNS for %s:%s resolved to %v", i.Address, i.RType, lookup)
	} else {
		i.Message = "Unknown RType"
		i.Failed = true
	}
}

func NewItem(address string) t.Item {
	rType := "A"
	if strings.Contains(address, ":") {
		splittedAddress := strings.Split(address, ":")
		rType = splittedAddress[1]
		address = splittedAddress[0]
	}
	return &Item{
		Address: address,
		RType:   rType,
		Message: notCheckedMessage,
	}
}
