package probe

import (
	"testing"

	"gotest.tools/assert"
)

func TestList(t *testing.T) {
	probe := NewList()
	assert.Equal(t, probe.GetType(), ListType)
	probe.Add(NewNumber(2, "int"))
	probe.Add(NewGenerator(func() Probe {
		return NewNumber(10, "int")
	}))

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

	// parametrized elements in list quite tricky
	probe.Add(NewParametrized(NewNumber(20, "int"), NewProbeInput("", "", "==", "20"))) // probe always true
	assert.Assert(t, GetProbeResult(probe, "", "sum", "==", "32"))                      // includes in aggregations
	assert.Assert(t, GetProbeResult(probe, "", "", "!=", "20"))                         // but still ignores list input

	probe.Add(NewParametrized(NewNumber(30, "int"), NewProbeInput("", "", "==", "20"))) // probe always false
	assert.Assert(t, GetProbeResult(probe, "", "sum", "==", "62"))                      // includes in aggregations
	assert.Assert(t, !GetProbeResult(probe, "", "", "<", "100"))                        // but still ignores list input and fails
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
