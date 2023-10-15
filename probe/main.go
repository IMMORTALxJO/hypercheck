package probe

import (
	"database/sql"
	"fmt"
	"strings"

	"hypercheck/probe/drivers/tcp"

	log "github.com/sirupsen/logrus"

	t "hypercheck/probe/types"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Probe struct {
	db      *gorm.DB
	drivers map[string]*DriverLink
}

type DriverLink struct {
	link        t.Driver
	initialized bool
}

func (p *Probe) GetDB() *gorm.DB {
	return p.db
}

func (p *Probe) Add(driverName string, input string) {
	if driver, ok := p.drivers[driverName]; ok {
		if !driver.initialized {
			driver.link.Initialize(p.db)
			driver.initialized = true
		}
		driver.link.GenerateProbe(p.db, input)
	} else {
		panic("unknown driver " + driverName)
	}
}

func (p *Probe) Run() {
	for _, driver := range p.drivers {
		driver.link.Enrich(p.db)
	}
}

func (p *Probe) Validate() {
	for _, driver := range p.drivers {
		driverName := strings.ToUpper(driver.link.Name())
		for _, item := range driver.link.GetItems(p.db) {
			log.Debugf("%s probe: %+v", driverName, item)
			emoji := "✅"
			if item.IsFailed() {
				emoji = "❌"
			}
			fmt.Printf("%s %s %s\n", emoji, driverName, item.GetMessage())
		}
	}
}

func New() *Probe {
	ramdb, _ := sql.Open("ramsql", "hypercheck")
	db, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: ramdb,
	}), &gorm.Config{})

	return &Probe{
		db: db,
		drivers: map[string]*DriverLink{
			"tcp": {
				link:        &tcp.Driver{},
				initialized: false,
			},
		},
	}
}
