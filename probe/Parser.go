package probe

import (
	"strconv"

	"code.cloudfoundry.org/bytefmt"
	log "github.com/sirupsen/logrus"
)

type Parser interface {
	Parse(string) (uint64, error)
	GetType() string
}

type ParserInt struct{}

func (p *ParserInt) Parse(input string) (uint64, error) {
	parsed, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	value := uint64(parsed)
	log.Debugf("ParserInt('%s') => %d", input, value)
	return value, nil
}

func (p *ParserInt) GetType() string {
	return "int"
}

type ParserBytes struct{}

func (p *ParserBytes) Parse(input string) (uint64, error) {
	parsed, err := bytefmt.ToBytes(input)
	if err != nil {
		return 0, err
	}
	value := uint64(parsed)
	log.Debugf("ParserBytes('%s') => %d", input, value)
	return value, nil
}

func (p *ParserBytes) GetType() string {
	return "bytes"
}
