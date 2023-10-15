package db

import (
	"fmt"
	"hypercheck/probe/types"
)

const Name = "DB"

func GenerateProbe(input string) (types.Probe, string) {
	data := getDbWrapper(input)
	description := "Check databases (pgsql, mysql)"
	if len(input) > 0 {
		description = fmt.Sprintf("Database %s", input)
	}
	dbProbe := types.NewMap(description, len(input) > 0)
	dbProbe.Add("online", types.NewGenerator("no connection errors", types.BoolType, func() types.Probe {
		return types.NewBool("", data.getOnline())
	}))
	dbProbe.Add("offline", types.NewGenerator("has connection errors", types.BoolType, func() types.Probe {
		return types.NewBool("", !data.getOnline())
	}))

	return dbProbe, ""
}
