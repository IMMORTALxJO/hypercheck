package types

import (
	"testing"

	"gotest.tools/assert"
)

func getParserResult(parser Parser, input string) uint64 {
	result, _ := parser.Parse(input)
	return result
}

func getParserMsg(parser Parser, input string) string {
	_, msg := parser.Parse(input)
	return msg
}

func TestParser(t *testing.T) {
	assert.Equal(t, getParserResult(&ParserInt{}, "1"), uint64(1))
	assert.Equal(t, getParserResult(&ParserInt{}, "0"), uint64(0))
	assert.Equal(t, getParserResult(&ParserInt{}, "999999999"), uint64(999999999))
	assert.Equal(t, getParserResult(&ParserBytes{}, "1b"), uint64(1))
	assert.Equal(t, getParserResult(&ParserBytes{}, "10b"), uint64(10))
	assert.Equal(t, getParserResult(&ParserBytes{}, "1Kb"), uint64(1024))
	assert.Equal(t, getParserResult(&ParserBytes{}, "1Mb"), uint64(1024*1024))
	assert.Equal(t, getParserResult(&ParserDuration{}, "1s"), uint64(1))
	assert.Equal(t, getParserResult(&ParserDuration{}, "10s"), uint64(10))
	assert.Equal(t, getParserResult(&ParserDuration{}, "1m"), uint64(60))
	assert.Equal(t, getParserResult(&ParserDuration{}, "1h1m1s"), uint64(60*60+60+1))
}

func TestParserErr(t *testing.T) {
	assert.Assert(t, getParserMsg(&ParserInt{}, "1") == "")
	assert.Assert(t, getParserMsg(&ParserInt{}, "bad") != "")
	assert.Assert(t, getParserMsg(&ParserBytes{}, "1b") == "")
	assert.Assert(t, getParserMsg(&ParserBytes{}, "bad") != "")
	assert.Assert(t, getParserMsg(&ParserDuration{}, "1s") == "")
	assert.Assert(t, getParserMsg(&ParserDuration{}, "bad") != "")
}
