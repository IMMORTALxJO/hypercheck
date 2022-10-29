package types

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

const TestMapDescription = "Map description"

func NewMapTest() *Map {
	return NewMap(TestMapDescription)
}

func TestMap(t *testing.T) {
	probe := NewMapTest()
	probeNumbersList := NewListTest()
	probe.Add("number", NewNumberTest(2, "int"))
	probe.Add("string", NewStringTest("teststring"))
	probeNumbersList.Add(NewNumberTest(10, "int"))
	probeNumbersList.Add(NewNumberTest(100, "int"))
	probe.Add("numbers", probeNumbersList)
	assert.Assert(t, getProbeResult(probe, "number", "", ">", "1"))
	assert.Assert(t, getProbeResult(probe, "number", "", "<", "10"))
	assert.Assert(t, getProbeResult(probe, "string", "", "==", "teststring"))
	assert.Assert(t, getProbeResult(probe, "numbers", "", ">=", "10"))
	assert.Assert(t, getProbeResult(probe, "numbers", "", "<=", "100"))

	assert.Assert(t, !getProbeResult(probe, "numbers", "", "<", "100"))
	assert.Assert(t, !getProbeResult(probe, "number", "", "==", "1"))

	assert.Assert(t, getProbeMsg(probe, "notfound", "", ">", "0") != "")
}

func TestMapMeta(t *testing.T) {
	probe := NewMapTest()
	numberProbe := NewNumberTest(10, "int")
	probe.Add("number", numberProbe)
	assert.Equal(t, probe.GetType(), MapType)
	assert.Equal(t, probe.GetDescription(), fmt.Sprintf("%s\n\tnumber - %s", TestMapDescription, numberProbe.GetDescription()))
}
