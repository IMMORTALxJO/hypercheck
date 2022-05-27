package db

import "probe/probe"

func GenerateProbe(input string) (probe.Probe, string) {
	data := getDbWrapper(input)
	dbProbe := probe.NewMap()
	dbProbe.Add("online", probe.NewGenerator(func() probe.Probe {
		return probe.NewBool(data.GetOnline())
	}))
	dbProbe.Add("offline", probe.NewGenerator(func() probe.Probe {
		return probe.NewBool(!data.GetOnline())
	}))

	return dbProbe, ""
}
