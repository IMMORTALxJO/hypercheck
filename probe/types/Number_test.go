package types

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

const TestNumberDescription = "Number description"

func NewNumberTest(value uint64, parser string) *Number {
	return NewNumber(TestNumberDescription, value, parser)
}

func TestNumber(t *testing.T) {
	probe := NewNumberTest(1, "int")
	probeBytes := NewNumberTest(1024, "bytes")
	probeDuration := NewNumberTest(301, "duration")
	assert.Equal(t, probe.GetParserType(), "int")
	assert.Equal(t, probeBytes.GetParserType(), "bytes")
	assert.Equal(t, probeDuration.GetParserType(), "duration")

	assert.Assert(t, getProbeResult(probe, "", "", "==", "1"))
	assert.Assert(t, getProbeResult(probe, "", "", "!=", "2"))
	assert.Assert(t, getProbeResult(probe, "", "", "=", "1"))
	assert.Assert(t, getProbeResult(probe, "", "", ">", "0"))
	assert.Assert(t, getProbeResult(probe, "", "", ">=", "1"))
	assert.Assert(t, getProbeResult(probe, "", "", "<", "2"))
	assert.Assert(t, getProbeResult(probe, "", "", "<=", "1"))
	assert.Assert(t, getProbeResult(probe, "", "", "<=", "1"))
	assert.Assert(t, getProbeResult(probeBytes, "", "", "==", "1Kb"))
	assert.Assert(t, getProbeResult(probeBytes, "", "", "<", "1Mb"))
	assert.Assert(t, getProbeResult(probeBytes, "", "", ">", "1B"))
	assert.Assert(t, getProbeResult(probeDuration, "", "", ">", "5m"))
}

func TestNumberErr(t *testing.T) {
	probe := NewNumberTest(1, "int")
	probeBytes := NewNumberTest(1024, "bytes")
	assert.Assert(t, getProbeMsg(probe, "", "", "badop", "0") != "")
	assert.Assert(t, getProbeMsg(probe, "", "", "!=", "10badparse") != "")
	assert.Assert(t, getProbeMsg(probeBytes, "", "", "!=", "1badparse") != "")
}

func TestNumberMeta(t *testing.T) {
	probe := NewNumberTest(1, "int")
	assert.Equal(t, probe.GetType(), NumberType)
	assert.Equal(t, probe.GetDescription(), fmt.Sprintf("%s ( %s )", TestNumberDescription, NumberType))
}
