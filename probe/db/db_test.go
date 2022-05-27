package db

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
	// listen port
	assert.Assert(t, probe.GetProbeResult(getProbe("postgres://user:password@localhost:5432/postgres?sslmode=disable"), "online", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(getProbe("postgres://user:password@localhost:5432/postgres?sslmode=disable"), "offline", "", "", ""))
	assert.Assert(t, probe.GetProbeResult(getProbe("postgres://baduser:password@localhost:5432/postgres?sslmode=disable"), "offline", "", "", ""))
	assert.Assert(t, probe.GetProbeResult(getProbe("postgres://user:badpassword@localhost:5432/postgres?sslmode=disable"), "offline", "", "", ""))

	// listen port but not redis
	assert.Assert(t, probe.GetProbeResult(getProbe("postgres://user:password@localhost:8080/postgres?sslmode=disable"), "offline", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(getProbe("postgres://user:password@localhost:8080/postgres?sslmode=disable"), "online", "", "", ""))

	// not listen port but not redis
	assert.Assert(t, probe.GetProbeResult(getProbe("postgres://user:password@localhost:8081/postgres?sslmode=disable"), "offline", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(getProbe("postgres://user:password@localhost:8081/postgres?sslmode=disable"), "online", "", "", ""))

}

func TestMysql(t *testing.T) {
	// listen port
	assert.Assert(t, probe.GetProbeResult(getProbe("mysql://user:password@localhost:3306/database"), "online", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(getProbe("mysql://user:password@localhost:3306/database"), "offline", "", "", ""))
	assert.Assert(t, probe.GetProbeResult(getProbe("mysql://baduser:password@localhost:3306/database"), "offline", "", "", ""))
	assert.Assert(t, probe.GetProbeResult(getProbe("mysql://user:badpassword@localhost:3306/database"), "offline", "", "", ""))

	// listen port but not redis
	assert.Assert(t, probe.GetProbeResult(getProbe("mysql://user:password@localhost:8080/database"), "offline", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(getProbe("mysql://user:password@localhost:8080/database"), "online", "", "", ""))

	// not listen port but not redis
	assert.Assert(t, probe.GetProbeResult(getProbe("mysql://user:password@localhost:8081/database"), "offline", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(getProbe("mysql://user:password@localhost:8081/database"), "online", "", "", ""))

}
