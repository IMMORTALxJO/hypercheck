package tcp

import (
	"net"
	"time"

	t "hypercheck/probe/types"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Driver struct{}

func (Driver) Name() string {
	return "tcp"
}

func (d *Driver) Initialize(db *gorm.DB) {
	if err := db.Migrator().CreateTable(&Item{}); err != nil {
		log.Error(err)
	}
}

func (d *Driver) GenerateProbe(db *gorm.DB, input string) {
	if err := db.Create(&Item{Address: input}).Error; err != nil {
		log.Error(err)
		return
	}
	log.Debugf("TCP probe generated for %s", input)
}

func (d *Driver) Enrich(db *gorm.DB) {
	var items []Item
	if err := db.Where("checked = ?", false).Find(&items).Error; err != nil {
		log.Error(err)
		return
	}
	for _, item := range items {
		item.Checked = true
		item.Message = "is online"
		item.Failed = false
		netd := net.Dialer{Timeout: time.Duration(1) * time.Second}
		startTime := time.Now()
		log.Debugf("checking TCP for %s", item.Address)
		conn, err := netd.Dial("tcp", item.Address)
		if err != nil {
			item.Message = err.Error()
			log.Debugf("TCP for %s failed: %s", item.Address, err)
			item.Failed = true
		}
		item.LatencyNano = uint64(time.Since(startTime).Nanoseconds())
		log.Debugf("TCP for %s took %d nanoseconds", item.Address, item.LatencyNano)
		if err != nil {
			item.Message = err.Error()
			item.Failed = true
		} else {
			defer conn.Close()
		}
		if err := db.Save(&item).Error; err != nil {
			log.Error(err)
		}
	}
}

func (d *Driver) GetItems(db *gorm.DB) []t.Item {
	var items []Item
	if err := db.Find(&items).Error; err != nil {
		log.Error(err)
		return nil
	}

	result := make([]t.Item, len(items))
	for i, item := range items {
		result[i] = item
	}
	return result
}
