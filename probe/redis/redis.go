package redis

import (
	"fmt"
	"hypercheck/probe/types"
)

const Name = "redis"

func GenerateProbe(input string) (types.Probe, string) {
	data := getRedisWrapper(input)
	description := "Check redis kv storage"
	if len(input) > 0 {
		description = fmt.Sprintf("Redis %s", input)
	}
	redisProbe := types.NewMap(description, len(input) > 0)
	redisProbe.Add("online", types.NewGenerator("PING-PONG success", types.BoolType, func() types.Probe {
		return types.NewBool("", data.getPing())
	}))
	redisProbe.Add("offline", types.NewGenerator("PING-PONG failed", types.BoolType, func() types.Probe {
		return types.NewBool("", !data.getPing())
	}))

	return redisProbe, ""
}
