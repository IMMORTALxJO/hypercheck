package redis

import "probe/probe"

func GenerateProbe(input string) (probe.Probe, string) {
	data := getRedisWrapper(input)
	redisProbe := probe.NewMap()
	redisProbe.Add("online", probe.NewGenerator(func() probe.Probe {
		return probe.NewBool(data.GetPing())
	}))
	redisProbe.Add("offline", probe.NewGenerator(func() probe.Probe {
		return probe.NewBool(!data.GetPing())
	}))

	return redisProbe, ""
}
