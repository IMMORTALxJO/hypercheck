package probe

import (
	"fmt"

	"hypercheck/probe/items/dns"
	"hypercheck/probe/items/tcp"

	log "github.com/sirupsen/logrus"

	t "hypercheck/probe/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Probe struct {
	db     *gorm.DB
	tables map[string]*ItemLink
}

type ItemLink struct {
	newItem     func(input string) t.Item
	initialized bool
}

func (p *Probe) GetDB() *gorm.DB {
	return p.db
}

func (p *Probe) Add(driverName string, input string) {
	if driver, ok := p.tables[driverName]; ok {
		item := driver.newItem(input)
		if !driver.initialized {
			p.db.AutoMigrate(item)
			driver.initialized = true
		}
		item.Enrich()
		p.db.Save(item)
	} else {
		panic("unknown driver " + driverName)
	}
}

func (p *Probe) Validate() int {
	exitCode := 0
	for driverName, driver := range p.tables {
		log.Debugf("validating %s", driverName)
		if !driver.initialized {
			log.Debugf("skipping %s", driverName)
			continue
		}
		var items []t.Item
		switch driverName {
		case "tcp":
			var tcpItems []tcp.Item
			p.db.Find(&tcpItems)
			items = make([]t.Item, len(tcpItems))
			for i, tcpItem := range tcpItems {
				items[i] = t.Item(&tcpItem) // Преобразование tcp.Item в t.Item
			}
		case "dns":
			var dnsItems []dns.Item
			p.db.Find(&dnsItems)
			items = make([]t.Item, len(dnsItems))
			for i, dnsItem := range dnsItems {
				items[i] = t.Item(&dnsItem) // Преобразование dns.Item в t.Item
			}
		}
		log.Debugf("found %d items for %s", len(items), driverName)
		for _, item := range items {
			log.Debugf("%s probe: %+v", driverName, item)
			emoji := "✅"
			if item.IsFailed() {
				emoji = "❌"
				exitCode = 1
			}
			fmt.Printf("%s %s %s\n", emoji, driverName, item.GetMessage())
		}
	}
	return exitCode
}

func (p *Probe) Exec(batch []string) {
	// print queries result
	for _, query := range batch {
		log.Debugf("executing query %s", query)
		p.db.Exec(query)
	}
}

func New() *Probe {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	return &Probe{
		db: db,
		tables: map[string]*ItemLink{
			"tcp": {
				newItem: tcp.NewItem,
			},
			"dns": {
				newItem: dns.NewItem,
			},
		},
	}
}
