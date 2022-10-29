package types

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

const TestListDescription = "List description"

func NewListTest() *List {
	return NewList(TestListDescription)
}

func TestList(t *testing.T) {
	probe := NewListTest()
	assert.Equal(t, probe.GetType(), ListType)
	probe.Add(NewNumberTest(2, "int"))
	probe.Add(NewGeneratorTest(NumberType, func() Probe {
		return NewNumberTest(10, "int")
	}))

	assert.Assert(t, getProbeResult(probe, "", "", ">", "0"))
	assert.Assert(t, getProbeResult(probe, "", "", ">=", "2"))
	assert.Assert(t, getProbeResult(probe, "", "", "<=", "10"))

	assert.Assert(t, !getProbeResult(probe, "", "", "<", "10"))
	assert.Assert(t, !getProbeResult(probe, "", "", ">", "2"))
	assert.Assert(t, !getProbeResult(probe, "", "", "==", "100"))

	assert.Assert(t, !getProbeResult(probe, "", "", "==", "2"))
	assert.Assert(t, getProbeResult(probe, "", "any", "==", "2"))
	assert.Assert(t, getProbeResult(probe, "", "count", "==", "2"))
	assert.Assert(t, getProbeResult(probe, "", "sum", "==", "12"))

	// parametrized elements in list quite tricky
	probe.Add(NewParametrized(NewNumberTest(20, "int"), NewProbeInput("", "", "==", "20"))) // probe always true
	assert.Assert(t, getProbeResult(probe, "", "sum", "==", "32"))                          // includes in aggregations
	assert.Assert(t, getProbeResult(probe, "", "", "!=", "20"))                             // but still ignores list input

	probe.Add(NewParametrized(NewNumberTest(30, "int"), NewProbeInput("", "", "==", "20"))) // probe always false
	assert.Assert(t, getProbeResult(probe, "", "sum", "==", "62"))                          // includes in aggregations
	assert.Assert(t, !getProbeResult(probe, "", "", "<", "100"))                            // but still ignores list input and fails

	probe.Add(NewStringTest("stringtest"))
	assert.Assert(t, getProbeResult(probe, "", "any", "==", "stringtest"))
}

func TestListErr(t *testing.T) {
	probe := NewListTest()
	probe.Add(NewNumberTest(2, "int"))
	assert.Assert(t, getProbeMsg(probe, "", "", "badop", "0") != "")
	assert.Assert(t, getProbeMsg(probe, "", "badaggr", ">=", "0") != "")
	probe.Add(NewStringTest("test"))
	assert.Assert(t, getProbeMsg(probe, "", "sum", ">=", "0") != "")
	probe = NewListTest()
	probe.Add(NewNumberTest(2, "int"))
	probe.Add(NewNumberTest(2, "bytes"))
	assert.Assert(t, getProbeMsg(probe, "", "sum", ">=", "0") != "")
}

func TestListMeta(t *testing.T) {
	probe := NewListTest()
	assert.Equal(t, probe.GetType(), ListType)
	assert.Equal(t, probe.GetDescription(), fmt.Sprintf("%s ( %s )", TestListDescription, ListType))
}
