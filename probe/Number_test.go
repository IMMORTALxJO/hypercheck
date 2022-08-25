package probe

import (
	"testing"

	"gotest.tools/assert"
)

func TestNumber(t *testing.T) {
	probe := NewNumber(1, "int")
	probeBytes := NewNumber(1024, "bytes")
	probeDuration := NewNumber(301, "duration")
	assert.Equal(t, probe.GetType(), NumberType)
	assert.Equal(t, probeBytes.GetType(), NumberType)
	assert.Equal(t, probeDuration.GetType(), NumberType)
	assert.Equal(t, probe.GetParserType(), "int")
	assert.Equal(t, probeBytes.GetParserType(), "bytes")
	assert.Equal(t, probeDuration.GetParserType(), "duration")

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
	assert.Assert(t, GetProbeResult(probeDuration, "", "", ">", "5m"))
}

func TestNumberErr(t *testing.T) {
	probe := NewNumber(1, "int")
	probeBytes := NewNumber(1024, "bytes")
	assert.Assert(t, GetProbeMsg(probe, "", "", "badop", "0") != "")
	assert.Assert(t, GetProbeMsg(probe, "", "", "!=", "10badparse") != "")
	assert.Assert(t, GetProbeMsg(probeBytes, "", "", "!=", "1badparse") != "")
}
