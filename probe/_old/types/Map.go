package types

import (
	"fmt"
	"strings"
)

const MapType = "Map"

type Map struct {
	description        string
	descriptionIsShort bool
	value              map[string]Probe
	keys               []string
}

func (p *Map) Up(input *Input) (bool, string) {
	key := input.Key
	probe := p.value[key]
	if probe == nil {
		return false, fmt.Sprintf("Unknow check name '%s' ( allowed '%s' )", key, strings.Join(p.keys, "','"))
	}
	return probe.Up(input)
}

func (p *Map) Add(key string, probe Probe) {
	p.value[key] = probe
	p.keys = append(p.keys, key)
}

func (*Map) GetType() string {
	return MapType
}

func (p *Map) GetDescription() string {
	if p.descriptionIsShort {
		return p.description
	}
	description := p.description
	for key, probe := range p.value {
		description += fmt.Sprintf("\n\t%s - %s", key, probe.GetDescription())
	}
	return description
}

func NewMap(description string, descriptionIsShort bool) *Map {
	return &Map{
		description:        description,
		descriptionIsShort: descriptionIsShort,
		value:              map[string]Probe{},
		keys:               []string{},
	}
}
