package types

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

const TestGeneratorDescription = "Generator description"

func NewGeneratorTest(subType string, value ProbeGenerator) *Generator {
	return NewGenerator(TestGeneratorDescription, subType, value)
}

func TestGenerator(t *testing.T) {
	assert.Assert(t, getProbeResult(NewGeneratorTest(NumberType, func() Probe {
		return NewNumberTest(1, "int")
	}), "", "", ">", "0"))
	assert.Assert(t, getProbeResult(NewGeneratorTest(NumberType, func() Probe {
		return NewNumberTest(2, "int")
	}), "", "", "!=", "1"))

}

func TestGeneratorMeta(t *testing.T) {
	probe := NewGeneratorTest(NumberType, func() Probe {
		return NewNumberTest(1, "int")
	})
	assert.Equal(t, probe.GetType(), GeneratorType)
	assert.Equal(t, probe.GetGeneratedType(), NumberType)
	assert.Equal(t, probe.GetDescription(), fmt.Sprintf("%s ( %s )", TestGeneratorDescription, NumberType))
}
