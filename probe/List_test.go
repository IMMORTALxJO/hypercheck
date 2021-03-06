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

}

func TestListErr(t *testing.T) {
	probe := NewList()
	probe.Add(NewNumber(2, "int"))
	assert.Assert(t, GetProbeMsg(probe, "", "", "badop", "0") != "")
	assert.Assert(t, GetProbeMsg(probe, "", "badaggr", ">=", "0") != "")
	probe.Add(NewString("test"))
	assert.Assert(t, GetProbeMsg(probe, "", "sum", ">=", "0") != "")
	probe = NewList()
	probe.Add(NewNumber(2, "int"))
	probe.Add(NewNumber(2, "bytes"))
	assert.Assert(t, GetProbeMsg(probe, "", "sum", ">=", "0") != "")
}
