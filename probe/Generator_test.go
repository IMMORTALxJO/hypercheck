package probe

import (
	"testing"

	"gotest.tools/assert"
)

func TestGenerator(t *testing.T) {
	assert.Equal(t, NewGenerator(func() Probe {
		return NewNumber(1, "int")
	}).GetType(), GeneratorType)
	assert.Assert(t, GetProbeResult(NewGenerator(func() Probe {
		return NewNumber(1, "int")
	}), "", "", ">", "0"))
	assert.Assert(t, GetProbeResult(NewGenerator(func() Probe {
		return NewNumber(2, "int")
	}), "", "", "!=", "1"))

}
