package probe

import (
	"testing"

	"gotest.tools/assert"
)

func TestList(t *testing.T) {
	probe := NewList()
	assert.Equal(t, probe.GetType(), "List")
	probe.Add(NewNumber(2, "int"))
	probe.Add(NewNumber(10, "int"))
	assert.Assert(t, GetProbeResult(probe, "", "", ">", "0"))
	assert.Assert(t, GetProbeResult(probe, "", "", ">=", "2"))
	assert.Assert(t, GetProbeResult(probe, "", "", "<=", "10"))

	assert.Assert(t, !GetProbeResult(probe, "", "", "<", "10"))
	assert.Assert(t, !GetProbeResult(probe, "", "", ">", "2"))
	assert.Assert(t, !GetProbeResult(probe, "", "", "==", "100"))

	assert.Assert(t, !GetProbeResult(probe, "", "", "==", "2"))
	assert.Assert(t, GetProbeResult(probe, "", "any", "==", "2"))
	assert.Assert(t, GetProbeResult(probe, "", "count", "==", "2"))
	assert.Assert(t, GetProbeResult(probe, "", "sum", "==", "12"))

	assert.Assert(t, GetProbeMsg(probe, "", "", "badop", "0") != "")
}
