package types

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestParametrized(t *testing.T) {
	stringProbeWrapped := NewParametrized(NewStringTest("teststring"), NewProbeInput("", "", "==", "teststring"))
	assert.Equal(t, stringProbeWrapped.GetType(), ParametrizedType)
	assert.Assert(t, getProbeResult(stringProbeWrapped, "does", "not", "matter", "=)"))

	numberProbeWrapped := NewParametrized(NewNumberTest(12, "int"), NewProbeInput("", "", "==", "12"))
	assert.Equal(t, numberProbeWrapped.GetType(), ParametrizedType)
	assert.Assert(t, getProbeResult(numberProbeWrapped, "does", "not", "matter", "=)"))
	numberProbeWrapped = NewParametrized(NewNumberTest(12, "int"), NewProbeInput("", "", "!=", "12"))
	assert.Assert(t, !getProbeResult(numberProbeWrapped, "does", "not", "matter", "=)"))

	boolProbeWrapped := NewParametrized(NewBoolTest(true), NewProbeInput("", "", "", ""))
	assert.Equal(t, boolProbeWrapped.GetType(), ParametrizedType)
	assert.Assert(t, getProbeResult(boolProbeWrapped, "does", "not", "matter", "=)"))
	boolProbeWrapped = NewParametrized(NewBoolTest(false), NewProbeInput("", "", "", ""))
	assert.Assert(t, !getProbeResult(boolProbeWrapped, "does", "not", "matter", "=)"))

	listProbe := NewListTest()
	listProbe.Add(NewNumberTest(12, "int"))
	listProbe.Add(NewNumberTest(14, "int"))
	listProbeWrapped := NewParametrized(listProbe, NewProbeInput("", "", ">", "10"))
	assert.Equal(t, listProbeWrapped.GetType(), ParametrizedType)
	assert.Assert(t, getProbeResult(listProbeWrapped, "does", "not", "matter", "=)"))
	listProbeWrapped = NewParametrized(listProbe, NewProbeInput("", "", "<", "12"))
	assert.Assert(t, !getProbeResult(listProbeWrapped, "does", "not", "matter", "=)"))

	generatorProbeWrapped := NewParametrized(NewGeneratorTest(NumberType, func() Probe {
		return NewNumberTest(1, "int")
	}), NewProbeInput("", "", "==", "1"))
	assert.Equal(t, generatorProbeWrapped.GetType(), ParametrizedType)
	assert.Assert(t, getProbeResult(generatorProbeWrapped, "does", "not", "matter", "=)"))
	generatorProbeWrapped = NewParametrized(NewGeneratorTest(NumberType, func() Probe {
		return NewNumberTest(1, "int")
	}), NewProbeInput("", "", "!=", "1"))
	assert.Assert(t, !getProbeResult(generatorProbeWrapped, "does", "not", "matter", "=)"))

	mapProbe := NewMapTest(false)
	mapProbe.Add("true", NewBoolTest(true))
	mapProbe.Add("false", NewBoolTest(false))
	mapProbeWrapped := NewParametrized(mapProbe, NewProbeInput("true", "", "", ""))
	assert.Equal(t, mapProbeWrapped.GetType(), ParametrizedType)
	assert.Assert(t, getProbeResult(mapProbeWrapped, "does", "not", "matter", "=)"))
	mapProbeWrapped = NewParametrized(mapProbe, NewProbeInput("false", "", "", ""))
	assert.Assert(t, !getProbeResult(mapProbeWrapped, "does", "not", "matter", "=)"))
}

func TestParametrizedMeta(t *testing.T) {
	probe := NewParametrized(NewBoolTest(true), NewProbeInput("does", "not", "matter", "=)"))
	assert.Equal(t, probe.GetType(), ParametrizedType)
	assert.Equal(t, probe.GetDescription(), fmt.Sprintf("%s ( %s )", TestBoolDescription, BoolType))
}
