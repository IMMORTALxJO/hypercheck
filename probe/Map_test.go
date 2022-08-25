package probe

import (
	"testing"

	"gotest.tools/assert"
)

func TestMap(t *testing.T) {
	probe := NewMap()
	assert.Equal(t, probe.GetType(), MapType)

	probeNumbersList := NewList()
	probe.Add("number", NewNumber(2, "int"))
	probe.Add("string", NewString("teststring"))
	probeNumbersList.Add(NewNumber(10, "int"))
	probeNumbersList.Add(NewNumber(100, "int"))
	probe.Add("numbers", probeNumbersList)
	assert.Assert(t, GetProbeResult(probe, "number", "", ">", "1"))
	assert.Assert(t, GetProbeResult(probe, "number", "", "<", "10"))
	assert.Assert(t, GetProbeResult(probe, "string", "", "==", "teststring"))
	assert.Assert(t, GetProbeResult(probe, "numbers", "", ">=", "10"))
	assert.Assert(t, GetProbeResult(probe, "numbers", "", "<=", "100"))

	assert.Assert(t, !GetProbeResult(probe, "numbers", "", "<", "100"))
	assert.Assert(t, !GetProbeResult(probe, "number", "", "==", "1"))

	assert.Assert(t, GetProbeMsg(probe, "notfound", "", ">", "0") != "")
}
