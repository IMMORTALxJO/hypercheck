package types

import "gorm.io/gorm"

type Driver interface {
	Initialize(db *gorm.DB)
	GenerateProbe(db *gorm.DB, input string)
	Enrich(db *gorm.DB)
	GetItems(db *gorm.DB) []Item
	Name() string
}
