package options

import (
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

type Common struct {
	Name string
}

func (c *Common) GetType() string {
	return "Common"
}
func (c *Common) GetOptions() string {
	return "None"
}
func (c *Common) ParseOptions(option string) {
	if c.Name != option {
		log.Fatalf("Common.ParseOptions: check has wrong format '%s'", option)
	}
}
func (c *Common) GetOperation() string {
	return ""
}
func (c *Common) GetValue() string {
	return ""
}
func (c *Common) GetValueBytes() uint64 {
	return 0
}
func (c *Common) GetValueInt() uint64 {
	return 0
}

func (c *Common) GetCopy() Option {
	copy := &Common{}
	copier.Copy(&copy, &c)
	return copy
}
