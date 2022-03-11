package options

import (
	"regexp"
	"strconv"
	"strings"

	"code.cloudfoundry.org/bytefmt"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

type Condition struct {
	Name            string
	OpsList         []string
	ParserType      string
	checkOp         string
	checkValue      string
	checkValueBytes uint64
	checkValueInt   uint64
}

func (c *Condition) GetType() string {
	return "Conditional"
}
func (c *Condition) GetOptions() string {
	options := "op='" + c.checkOp + "',value='" + c.checkValue + "'"
	if c.ParserType == "bytes" {
		options = options + " ( " + strconv.FormatUint(c.checkValueBytes, 10) + " bytes )"
	}
	return options
}
func (c *Condition) ParseOptions(arg string) {
	chars := regexp.MustCompile("[!<>=]+")
	splitted := chars.Split(arg, -1)
	if c.Name != string(splitted[0]) || len(splitted) != 2 {
		log.Fatalf("Condition.ParseOptions: check argument has wrong format '%s'", arg)
	}
	c.checkValue = splitted[1]
	if c.ParserType == "Bytes" {
		parsedBytes, err := bytefmt.ToBytes(c.checkValue)
		if err != nil {
			log.Fatalf("Condition.ParseOptions: wrong bytes format '%s'", c.checkValue)
		}
		c.checkValueBytes = parsedBytes
	}
	if c.ParserType == "PosInt" {
		parsedInt, err := strconv.Atoi(c.checkValue)
		if err != nil || parsedInt < 0 {
			log.Fatalf("Condition.ParseOptions: couldn't parse positive integer '%s'", c.checkValue)
		}
		c.checkValueInt = uint64(parsedInt)
	}
	c.checkOp = strings.TrimSuffix(strings.TrimPrefix(arg, splitted[0]), splitted[1])

	log.Debugf("Condition.ParseOptions: parsed %s to options %s", arg, c.GetOptions())
}

func (c *Condition) GetOperation() string {
	return c.checkOp
}
func (c *Condition) GetValue() string {
	return c.checkValue
}
func (c *Condition) GetValueBytes() uint64 {
	return c.checkValueBytes
}
func (c *Condition) GetValueInt() uint64 {
	return c.checkValueInt
}

func (c *Condition) GetCopy() Option {
	copy := &Condition{}
	copier.Copy(&copy, &c)
	return copy
}
