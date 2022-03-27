package fs

import (
	"log"
	"path/filepath"
	"probe/probe"
)

func GenerateProbe(input string) probe.Probe {
	fsSizeProbe := probe.NewList()
	fsDirProbe := probe.NewList()
	fsRegularProbe := probe.NewList()
	fsUidProbe := probe.NewList()
	fsUserProbe := probe.NewList()
	fsGidProbe := probe.NewList()
	fsGroupProbe := probe.NewList()

	paths, err := filepath.Glob(input)
	if err != nil {
		log.Fatalf("Wrong files glob pattern '%s'", input)
	}
	for _, path := range paths {
		file := getFileWrapper(path)
		size := file.getSize()
		fsSizeProbe.Add(probe.NewNumber(size, "bytes"))
		fsDirProbe.Add(probe.NewBool(file.IsDir()))
		fsRegularProbe.Add(probe.NewBool(file.IsRegular()))
		fsUidProbe.Add(probe.NewNumber(file.getUID(), "int"))
		fsGidProbe.Add(probe.NewNumber(file.getGID(), "int"))
		fsGroupProbe.Add(probe.NewString(file.getGroupname()))
		fsUserProbe.Add(probe.NewString(file.getUsername()))
	}
	fsProbe := probe.NewMap()
	fsProbe.Add("size", fsSizeProbe)
	fsProbe.Add("dir", fsDirProbe)
	fsProbe.Add("regular", fsRegularProbe)
	fsProbe.Add("uid", fsUidProbe)
	fsProbe.Add("gid", fsGidProbe)
	fsProbe.Add("user", fsUserProbe)
	fsProbe.Add("group", fsGroupProbe)

	fsProbe.Add("count", probe.NewNumber(uint64(len(paths)), "int"))
	fsProbe.Add("exists", probe.NewBool(len(paths) > 0))
	return fsProbe
}
