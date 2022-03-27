package probe

import (
	"testing"

	"gotest.tools/assert"
)

func TestNumber(t *testing.T) {
	probe := NewNumber(1, "int")
	assert.Equal(t, probe.GetType(), "Number")

	assert.Assert(t, GetProbeResult(probe, "", "", "==", "1"))
	assert.Assert(t, GetProbeResult(probe, "", "", "=", "1"))
	assert.Assert(t, GetProbeResult(probe, "", "", ">", "0"))
	assert.Assert(t, GetProbeResult(probe, "", "", ">=", "1"))
	assert.Assert(t, GetProbeResult(probe, "", "", "<", "2"))
	assert.Assert(t, GetProbeResult(probe, "", "", "<=", "1"))
	assert.Assert(t, GetProbeResult(probe, "", "", "!=", "0"))

	assert.Assert(t, GetProbeMsg(probe, "", "", "badop", "0") != "")
	assert.Assert(t, GetProbeMsg(probe, "", "", "!=", "10badparse") != "")
}
