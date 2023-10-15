package tcp

import "fmt"

const tableName = "tcp"

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
