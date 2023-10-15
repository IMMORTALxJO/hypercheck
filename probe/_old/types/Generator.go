package types

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

const GeneratorType = "Generator"

type ProbeGenerator func() Probe

type Generator struct {
	description   string
	generatedType string
	function      ProbeGenerator
	value         Probe
}

func (p *Generator) Up(input *Input) (bool, string) {
	result, msg := p.GetValue().Up(input)
	return result, msg
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

func (p *Generator) GetGeneratedType() string {
	return p.generatedType
}

func (p *Generator) GetDescription() string {
	return fmt.Sprintf("%s ( %s )", p.description, p.GetGeneratedType())
}

func NewGenerator(description string, generatedType string, value ProbeGenerator) *Generator {
	return &Generator{
		generatedType: generatedType,
		description:   description,
		function:      value,
	}
}
