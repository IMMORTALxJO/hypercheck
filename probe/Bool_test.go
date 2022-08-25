package probe

import (
	"testing"

	"gotest.tools/assert"
)

func TestBool(t *testing.T) {
	probeTrue := NewBool(true)
	probeFalse := NewBool(false)
	assert.Equal(t, probeTrue.GetType(), BoolType)
	assert.Assert(t, GetProbeResult(probeTrue, "", "", "", ""))
	assert.Assert(t, !GetProbeResult(probeFalse, "", "", "", ""))
}
