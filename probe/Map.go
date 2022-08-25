package probe

import "fmt"

const MapType = "Map"

type Map struct {
	value map[string]Probe
}

func (p *Map) Up(input *Input) (bool, string) {
	key := input.Key
	probe := p.value[key]
	if probe == nil {
		return false, fmt.Sprintf("Unknow check name '%s'", key)
	}
	return probe.Up(input)
}

func (p *Map) Add(key string, probe Probe) {
	p.value[key] = probe
}

func (*Map) GetType() string {
	return MapType
}

func NewMap() *Map {
	return &Map{
		value: map[string]Probe{},
	}
}
