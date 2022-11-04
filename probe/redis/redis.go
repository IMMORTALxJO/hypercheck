package redis

import "hypercheck/probe/types"

const Name = "redis"

func GenerateProbe(input string) (types.Probe, string) {
	data := getRedisWrapper(input)
	redisProbe := types.NewMap("Test redis kv database")
	redisProbe.Add("online", types.NewGenerator("PING-PONG success", types.BoolType, func() types.Probe {
		return types.NewBool("", data.GetPing())
	}))
	redisProbe.Add("offline", types.NewGenerator("PING-PONG failed", types.BoolType, func() types.Probe {
		return types.NewBool("", !data.GetPing())
	}))

	return redisProbe, ""
}
