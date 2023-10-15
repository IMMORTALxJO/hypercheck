package types

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

const TestBoolDescription = "Bool description"

func NewBoolTest(value bool) *Bool {
	return NewBool(TestBoolDescription, value)
}

func getProbeResult(probe Probe, key string, agg string, op string, target string) bool {
	result, _ := probe.Up(NewProbeInput(key, agg, op, target))
	return result
}

func getProbeMsg(probe Probe, key string, agg string, op string, target string) string {
	_, msg := probe.Up(NewProbeInput(key, agg, op, target))
	return msg
}

func TestBool(t *testing.T) {
	probeTrue := NewBoolTest(true)
	probeFalse := NewBoolTest(false)
	assert.Equal(t, probeTrue.GetType(), BoolType)
	assert.Assert(t, getProbeResult(probeTrue, "", "", "", ""))
	assert.Assert(t, !getProbeResult(probeFalse, "", "", "", ""))
}

func TestBoolMeta(t *testing.T) {
	probe := NewBoolTest(false)
	assert.Equal(t, probe.GetType(), BoolType)
	assert.Equal(t, probe.GetDescription(), fmt.Sprintf("%s ( %s )", TestBoolDescription, BoolType))
}
