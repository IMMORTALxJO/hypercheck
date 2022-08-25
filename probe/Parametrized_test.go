package probe

import (
	"testing"

	"gotest.tools/assert"
)

func TestParametrized(t *testing.T) {
	stringProbeWrapped := NewParametrized(NewString("teststring"), NewProbeInput("", "", "==", "teststring"))
	assert.Equal(t, stringProbeWrapped.GetType(), ParametrizedType)
	assert.Assert(t, GetProbeResult(stringProbeWrapped, "does", "not", "matter", "=)"))

	numberProbeWrapped := NewParametrized(NewNumber(12, "int"), NewProbeInput("", "", "==", "12"))
	assert.Equal(t, numberProbeWrapped.GetType(), ParametrizedType)
	assert.Assert(t, GetProbeResult(numberProbeWrapped, "does", "not", "matter", "=)"))
	numberProbeWrapped = NewParametrized(NewNumber(12, "int"), NewProbeInput("", "", "!=", "12"))
	assert.Assert(t, !GetProbeResult(numberProbeWrapped, "does", "not", "matter", "=)"))

	boolProbeWrapped := NewParametrized(NewBool(true), NewProbeInput("", "", "", ""))
	assert.Equal(t, boolProbeWrapped.GetType(), ParametrizedType)
	assert.Assert(t, GetProbeResult(boolProbeWrapped, "does", "not", "matter", "=)"))
	boolProbeWrapped = NewParametrized(NewBool(false), NewProbeInput("", "", "", ""))
	assert.Assert(t, !GetProbeResult(boolProbeWrapped, "does", "not", "matter", "=)"))

	listProbe := NewList()
	listProbe.Add(NewNumber(12, "int"))
	listProbe.Add(NewNumber(14, "int"))
	listProbeWrapped := NewParametrized(listProbe, NewProbeInput("", "", ">", "10"))
	assert.Equal(t, listProbeWrapped.GetType(), ParametrizedType)
	assert.Assert(t, GetProbeResult(listProbeWrapped, "does", "not", "matter", "=)"))
	listProbeWrapped = NewParametrized(listProbe, NewProbeInput("", "", "<", "12"))
	assert.Assert(t, !GetProbeResult(listProbeWrapped, "does", "not", "matter", "=)"))

	generatorProbeWrapped := NewParametrized(NewGenerator(func() Probe {
		return NewNumber(1, "int")
	}), NewProbeInput("", "", "==", "1"))
	assert.Equal(t, generatorProbeWrapped.GetType(), ParametrizedType)
	assert.Assert(t, GetProbeResult(generatorProbeWrapped, "does", "not", "matter", "=)"))
	generatorProbeWrapped = NewParametrized(NewGenerator(func() Probe {
		return NewNumber(1, "int")
	}), NewProbeInput("", "", "!=", "1"))
	assert.Assert(t, !GetProbeResult(generatorProbeWrapped, "does", "not", "matter", "=)"))

	mapProbe := NewMap()
	mapProbe.Add("true", NewBool(true))
	mapProbe.Add("false", NewBool(false))
	mapProbeWrapped := NewParametrized(mapProbe, NewProbeInput("true", "", "", ""))
	assert.Equal(t, mapProbeWrapped.GetType(), ParametrizedType)
	assert.Assert(t, GetProbeResult(mapProbeWrapped, "does", "not", "matter", "=)"))
	mapProbeWrapped = NewParametrized(mapProbe, NewProbeInput("false", "", "", ""))
	assert.Assert(t, !GetProbeResult(mapProbeWrapped, "does", "not", "matter", "=)"))
}
