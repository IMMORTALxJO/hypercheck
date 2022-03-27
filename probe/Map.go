package probe

type Map struct {
	value map[string]Probe
}

func (p *Map) Up(input *Input) (bool, string) {
	key := input.Key
	probe := p.value[key]
	if probe == nil {
		return false, "Map key not found"
	}
	return probe.Up(input)
}

func (p *Map) Add(key string, probe Probe) {
	p.value[key] = probe
}

func (p *Map) GetType() string {
	return "Map"
}

func NewMap() *Map {
	return &Map{
		value: map[string]Probe{},
	}
}
