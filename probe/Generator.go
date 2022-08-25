package probe

import (
	log "github.com/sirupsen/logrus"
)

const GeneratorType = "Generator"

type ProbeGenerator func() Probe

type Generator struct {
	function ProbeGenerator
	value    Probe
}

func (p *Generator) Up(input *Input) (bool, string) {
	return p.GetValue().Up(input)
}

func (p *Generator) GetValue() Probe {
	if p.value == nil {
		p.value = p.function()
		log.Debugf("Generated probe '%v'", p.value.GetType())
	}
	log.Debugf("Got generated probe '%v'", p.value.GetType())
	return p.value
}

func (*Generator) GetType() string {
	return GeneratorType
}

func NewGenerator(value ProbeGenerator) *Generator {
	return &Generator{
		function: value,
	}
}
