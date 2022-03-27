package probe

import (
	"testing"

	"gotest.tools/assert"
)

type TestParser struct {
	Input  string
	Answer uint64
}

func getParserResult(parser Parser, input string) uint64 {
	result, _ := parser.Parse(input)
	return result
}

func TestParserInt(t *testing.T) {
	parserInt := &ParserInt{}
	assert.Equal(t, getParserResult(parserInt, "1"), uint64(1))
	assert.Equal(t, getParserResult(parserInt, "0"), uint64(0))
	assert.Equal(t, getParserResult(parserInt, "999999999"), uint64(999999999))

	parserBytes := &ParserBytes{}
	assert.Equal(t, getParserResult(parserBytes, "1b"), uint64(1))
	assert.Equal(t, getParserResult(parserBytes, "10b"), uint64(10))
	assert.Equal(t, getParserResult(parserBytes, "1Kb"), uint64(1024))
	assert.Equal(t, getParserResult(parserBytes, "1Mb"), uint64(1024*1024))
}
