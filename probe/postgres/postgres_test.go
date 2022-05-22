package postgres

import (
	"probe/probe"
	"testing"

	"gotest.tools/assert"
)

func getProbe(pattern string) probe.Probe {
	probe, _ := GenerateProbe(pattern)
	return probe
}

func TestPostgres(t *testing.T) {
	// listen postgres port
	assert.Assert(t, probe.GetProbeResult(getProbe("postgres://user:password@localhost:5432/postgres"), "online", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(getProbe("postgres://user:password@localhost:5432/postgres"), "offline", "", "", ""))

	// listen port but not redis
	assert.Assert(t, probe.GetProbeResult(getProbe("postgres://user:password@localhost:8080/postgres"), "offline", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(getProbe("postgres://user:password@localhost:8080/postgres"), "online", "", "", ""))

	// not listen port but not redis
	assert.Assert(t, probe.GetProbeResult(getProbe("postgres://user:password@localhost:8081/postgres"), "offline", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(getProbe("postgres://user:password@localhost:8081/postgres"), "online", "", "", ""))

}
