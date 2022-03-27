package probe

import (
	"testing"

	"gotest.tools/assert"
)

func TestNumber(t *testing.T) {
	probe := NewNumber(1, "int")
	probeBytes := NewNumber(1024, "bytes")
	assert.Equal(t, probe.GetType(), "Number")
	assert.Equal(t, probeBytes.GetType(), "Number")
	assert.Equal(t, probe.GetParserType(), "int")
	assert.Equal(t, probeBytes.GetParserType(), "bytes")

	assert.Assert(t, GetProbeResult(probe, "", "", "==", "1"))
	assert.Assert(t, GetProbeResult(probe, "", "", "!=", "2"))
	assert.Assert(t, GetProbeResult(probe, "", "", "=", "1"))
	assert.Assert(t, GetProbeResult(probe, "", "", ">", "0"))
	assert.Assert(t, GetProbeResult(probe, "", "", ">=", "1"))
	assert.Assert(t, GetProbeResult(probe, "", "", "<", "2"))
	assert.Assert(t, GetProbeResult(probe, "", "", "<=", "1"))
	assert.Assert(t, GetProbeResult(probe, "", "", "<=", "1"))
	assert.Assert(t, GetProbeResult(probeBytes, "", "", "==", "1Kb"))
	assert.Assert(t, GetProbeResult(probeBytes, "", "", "<", "1Mb"))
	assert.Assert(t, GetProbeResult(probeBytes, "", "", ">", "1B"))
}

func TestNumberErr(t *testing.T) {
	probe := NewNumber(1, "int")
	probeBytes := NewNumber(1024, "bytes")
	assert.Assert(t, GetProbeMsg(probe, "", "", "badop", "0") != "")
	assert.Assert(t, GetProbeMsg(probe, "", "", "!=", "10badparse") != "")
	assert.Assert(t, GetProbeMsg(probeBytes, "", "", "!=", "1badparse") != "")
}
