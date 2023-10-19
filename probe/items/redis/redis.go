package redis

import (
	"fmt"
	t "hypercheck/probe/types"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

const tableName = "redis"
const checkIsUpMessage = "online"
const notCheckedMessage = "not checked"

type Item struct {
	Address     string `gorm:"primaryKey"`
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
	return fmt.Sprintf("%s - %s", i.Address, i.Message)
}

func (i *Item) Enrich() {
	i.Checked = true
	i.Message = checkIsUpMessage
	i.Failed = false
	conn, err := redis.Dial("tcp", i.Address)
	if err != nil {
		i.Message = err.Error()
		log.Debugf("redis for %s failed: %s", i.Address, err)
		i.Failed = true
		return
	}
	answer, err := redis.DoWithTimeout(conn, time.Second, "ping")
	if err != nil {
		i.Message = err.Error()
		log.Debugf("redis for %s failed: %s", i.Address, err)
		i.Failed = true
		return
	}
	log.Debugf("redis answer for %s is %s", i.Address, answer)
	if answer != "PONG" {
		i.Message = "ping answer is not PONG"
		i.Failed = true
	}
}

func NewItem(address string) t.Item {
	addressWithPort := address
	if !strings.Contains(address, ":") {
		addressWithPort = address + ":6379"
	}
	return &Item{
		Address: addressWithPort,
		Message: notCheckedMessage,
	}
}
