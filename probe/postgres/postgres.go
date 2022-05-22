package postgres

import "probe/probe"

func GenerateProbe(input string) (probe.Probe, string) {
	data := getPostgresWrapper(input)
	postgresProbe := probe.NewMap()
	postgresProbe.Add("online", probe.NewGenerator(func() probe.Probe {
		return probe.NewBool(data.GetOnline())
	}))
	postgresProbe.Add("offline", probe.NewGenerator(func() probe.Probe {
		return probe.NewBool(!data.GetOnline())
	}))

	return postgresProbe, ""
}
