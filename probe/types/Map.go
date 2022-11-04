package types

import (
	"fmt"
	"strings"
)

const MapType = "Map"

type Map struct {
	description string
	value       map[string]Probe
	keys        []string
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
	description := p.description
	description += fmt.Sprintf(", attributes:")
	for key, probe := range p.value {
		description += fmt.Sprintf("\n\t%s - %s", key, probe.GetDescription())
	}
	return description
}

func NewMap(description string) *Map {
	return &Map{
		description: description,
		value:       map[string]Probe{},
		keys:        []string{},
	}
}
