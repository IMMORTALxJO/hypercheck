package fs

import (
	"path/filepath"
	"probe/probe"

	log "github.com/sirupsen/logrus"
)

func GenerateProbe(input string) (probe.Probe, string) {
	paths, err := filepath.Glob(input)
	if err != nil {
		return probe.NewMap(), "Wrong files glob pattern"
	}
	fsSizeProbe := probe.NewList()
	fsDirProbe := probe.NewList()
	fsRegularProbe := probe.NewList()
	fsUidProbe := probe.NewList()
	fsUserProbe := probe.NewList()
	fsGidProbe := probe.NewList()
	fsGroupProbe := probe.NewList()

	for _, path := range paths {
		file := getFileWrapper(path)
		fsSizeProbe.Add(probe.NewGenerator(func() probe.Probe {
			return probe.NewNumber(file.getSize(), "bytes")
		}))
		fsDirProbe.Add(probe.NewGenerator(func() probe.Probe {
			return probe.NewBool(file.IsDir())
		}))
		fsRegularProbe.Add(probe.NewGenerator(func() probe.Probe {
			return probe.NewBool(file.IsRegular())
		}))
		fsUidProbe.Add(probe.NewGenerator(func() probe.Probe {
			return probe.NewNumber(file.getUID(), "int")
		}))
		fsGidProbe.Add(probe.NewGenerator(func() probe.Probe {
			return probe.NewNumber(file.getGID(), "int")
		}))
		fsGroupProbe.Add(probe.NewGenerator(func() probe.Probe {
			return probe.NewString(file.getGroupname())
		}))
		fsUserProbe.Add(probe.NewGenerator(func() probe.Probe {
			return probe.NewString(file.getUsername())
		}))
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
	log.Debugf("fs probe initialized for '%s'", input)
	return fsProbe, ""
}
