package types

import (
	"testing"

	"gotest.tools/assert"
)

func TestInput(t *testing.T) {
	assert.Equal(t, NewProbeInput("num", "sum", ">", "10").ToString(), "num:sum>10")
	assert.Equal(t, NewProbeInput("num", "", ">", "10").ToString(), "num>10")
}
