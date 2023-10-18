package tcp

import (
	"fmt"
	t "hypercheck/probe/types"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
)

const tableName = "tcp"
const checkIsUpMessage = "is online"
const notCheckedMessage = "not checked"

type Item struct {
	Address     string `gorm:"primaryKey"`
	Checked     bool
	Failed      bool
	LatencyNano uint64
	Message     string
}

func (Item) TableName() string {
	return tableName
}

func (i Item) IsFailed() bool {
	return i.Failed
}

func (i Item) GetMessage() string {
	return fmt.Sprintf("%s - %s", i.Address, i.Message)
}

func (i *Item) Enrich() {
	i.Checked = true
	i.Message = checkIsUpMessage
	i.Failed = false
	netd := net.Dialer{Timeout: time.Duration(1) * time.Second}
	startTime := time.Now()
	log.Debugf("checking TCP for %s", i.Address)
	conn, err := netd.Dial("tcp", i.Address)
	if err != nil {
		i.Message = err.Error()
		log.Debugf("TCP for %s failed: %s", i.Address, err)
		i.Failed = true
	}
	i.LatencyNano = uint64(time.Since(startTime).Nanoseconds())
	log.Debugf("TCP for %s took %d nanoseconds", i.Address, i.LatencyNano)
	if err != nil {
		i.Message = err.Error()
		i.Failed = true
		return
	}
	conn.Close()
}

func NewItem(address string) t.Item {
	return &Item{
		Address: address,
		Message: notCheckedMessage,
	}
}
