package types

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

const TestStringDescription = "String description"

func NewStringTest(value string) *String {
	return NewString(TestStringDescription, value)
}
func TestString(t *testing.T) {
	probe := NewStringTest("teststring")
	assert.Assert(t, getProbeResult(probe, "", "", "==", "teststring"))
	assert.Assert(t, getProbeResult(probe, "", "", "!=", "abasd"))
	assert.Assert(t, getProbeResult(probe, "", "", "~", "tst"))
	assert.Assert(t, getProbeResult(probe, "", "", "~=", "^teststring$"))
	assert.Assert(t, getProbeResult(probe, "", "", "~=", "\\w+"))
	assert.Assert(t, getProbeResult(probe, "", "", "~=", ".*test.*"))
	assert.Assert(t, getProbeResult(probe, "", "", "~!", ".*abasd.*"))
	assert.Assert(t, getProbeResult(probe, "", "length", "==", "10"))

	assert.Assert(t, !getProbeResult(probe, "", "", "!=", "teststring"))
	assert.Assert(t, !getProbeResult(probe, "", "", "==", "abasd"))
	assert.Assert(t, !getProbeResult(probe, "", "", "~", "abasd"))
	assert.Assert(t, !getProbeResult(probe, "", "", "~=", "^abasd$"))
	assert.Assert(t, !getProbeResult(probe, "", "", "~=", ".*abasd.*"))
	assert.Assert(t, !getProbeResult(probe, "", "", "~!", ".*test.*"))

	assert.Assert(t, getProbeMsg(probe, "", "", "badop", "") != "")
	assert.Assert(t, getProbeMsg(probe, "", "", "~=", `\`) != "")
	assert.Assert(t, getProbeMsg(probe, "", "", "~!", `\`) != "")
}

func TestStringErr(t *testing.T) {
	probe := NewStringTest("teststring")
	assert.Assert(t, getProbeMsg(probe, "", "badaggr", ">=", "0") != "")
}

func TestStringMeta(t *testing.T) {
	probe := NewStringTest("teststring")
	assert.Equal(t, probe.GetType(), StringType)
	assert.Equal(t, probe.GetDescription(), fmt.Sprintf("%s ( %s )", TestStringDescription, StringType))
}
