package db

import "hypercheck/probe/types"

const Name = "DB"

func GenerateProbe(input string) (types.Probe, string) {
	data := getDbWrapper(input)
	dbProbe := types.NewMap("Check database ( pgsql, mysql )")
	dbProbe.Add("online", types.NewGenerator("is reachable", types.BoolType, func() types.Probe {
		return types.NewBool("", data.GetOnline())
	}))
	dbProbe.Add("offline", types.NewGenerator("is unreachable", types.BoolType, func() types.Probe {
		return types.NewBool("", !data.GetOnline())
	}))

	return dbProbe, ""
}
