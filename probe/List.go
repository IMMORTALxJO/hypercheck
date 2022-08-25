package probe

import (
	"fmt"
)

const ListType = "List"

type List struct {
	value []Probe
}

func (p *List) Up(input *Input) (bool, string) {
	aggregator := input.Aggregator
	if aggregator == "" {
		aggregator = "all"
	}
	if aggregator == "count" {
		probe := NewNumber(uint64(len(p.GetValue())), "int")
		return probe.Up(input)
	}
	if aggregator == "sum" {
		sum := uint64(0)
		parserType := ""
		for _, probe := range p.GetValue() {
			if probe.GetType() == GeneratorType {
				probe = probe.(*Generator).GetValue()
			}
			if probe.GetType() == ParametrizedType {
				probe = probe.(*Parametrized).GetValue()
			}
			if probe.GetType() != NumberType {
				return false, "Sum aggregation is for Numbers only"
			}
			sum = sum + probe.(*Number).GetValue()
			if probe.(*Number).GetParserType() != parserType {
				if parserType != "" {
					return false, "Multiple numbers type in single list"
				}
				parserType = probe.(*Number).GetParserType()
			}
		}
		probe := NewNumber(sum, parserType)
		return probe.Up(input)
	}
	for _, probe := range p.GetValue() {
		result, msg := probe.Up(input)
		if result && aggregator == "any" {
			return true, msg
		}
		if !result && aggregator == "all" {
			return false, msg
		}
	}
	if aggregator == "all" {
		return true, ""
	}
	return false, fmt.Sprintf("Unknown aggregator '%s'", aggregator)
}

func (p *List) Add(probe Probe) {
	p.value = append(p.value, probe)
}

func (*List) GetType() string {
	return ListType
}

func (p *List) GetValue() []Probe {
	return p.value
}

func NewList() *List {
	return &List{
		value: []Probe{},
	}
}
