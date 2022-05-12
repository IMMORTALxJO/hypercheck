package probe

import (
	"strconv"

	"code.cloudfoundry.org/bytefmt"
	log "github.com/sirupsen/logrus"
)

type Parser interface {
	Parse(string) (uint64, string)
	GetType() string
}

type ParserInt struct{}

func (*ParserInt) Parse(input string) (uint64, string) {
	value, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return 0, "Input parse error"
	}
	log.Debugf("ParserInt('%s') => %d", input, value)
	return value, ""
}

func (*ParserInt) GetType() string {
	return "int"
}

type ParserBytes struct{}

func (*ParserBytes) Parse(input string) (uint64, string) {
	parsed, err := bytefmt.ToBytes(input)
	if err != nil {
		return 0, "Input parse error"
	}
	value := uint64(parsed)
	log.Debugf("ParserBytes('%s') => %d", input, value)
	return value, ""
}

func (*ParserBytes) GetType() string {
	return "bytes"
}
