package cli

import (
	"testing"

	"gotest.tools/assert"
)

func check(t *testing.T, arguments string, size int, id int, key string, aggregator string, operator string, value string) {
	result := ParseArguments(arguments)
	assert.Equal(t, len(result), size)
	assert.Equal(t, result[id].Key, key)
	assert.Equal(t, result[id].Operator, operator)
	assert.Equal(t, result[id].Aggregator, aggregator)
	assert.Equal(t, result[id].Value, value)
}

func TestUtilsParseArguments(t *testing.T) {
	check(t, "key", 1, 0, "key", "", "", "")
	check(t, "key:aggr", 1, 0, "key", "aggr", "", "")
	check(t, "key:aggr>2", 1, 0, "key", "aggr", ">", "2")
	check(t, "key:aggr=~test", 1, 0, "key", "aggr", "=~", "test")

	// not used but possible cases
	check(t, "key:aggr:test", 1, 0, "key", "aggr", ":", "test")
	check(t, "size:sum>10Mb", 1, 0, "size", "sum", ">", "10Mb")
	check(t, "size>10Mb", 1, 0, "size", "", ">", "10Mb")
	check(t, "size:sum>10Mb,count>10", 2, 0, "size", "sum", ">", "10Mb")
	check(t, "size:sum>10Mb,count>10", 2, 1, "count", "", ">", "10")
}
