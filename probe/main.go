package probe

import (
	"fmt"

	"hypercheck/probe/items/tcp"

	log "github.com/sirupsen/logrus"

	t "hypercheck/probe/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Probe struct {
	db     *gorm.DB
	tables map[string]ItemLink
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
	for driverName, _ := range p.tables {
		var items []tcp.Item
		p.db.Find(&items)
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
		tables: map[string]ItemLink{
			"tcp": {
				newItem: tcp.NewItem,
			},
		},
	}
}
