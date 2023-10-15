package fs

import (
	"fmt"
	"hypercheck/probe/types"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

const Name = "FS"

func GenerateProbe(input string) (types.Probe, string) {
	description := "Filesystem files check"
	if len(input) > 0 {
		description = fmt.Sprintf("FileSystem %s", input)
	}

	fsProbe := types.NewMap(description, len(input) > 0)
	fsProbe.Add("size", types.NewGenerator("files size", "List[Number]", func() types.Probe {
		paths, _ := filepath.Glob(input)
		listProbe := types.NewList("")
		for _, path := range paths {
			file := getFileWrapper(path)
			listProbe.Add(types.NewNumber("", file.getSize(), "bytes"))
		}
		return listProbe
	}))
	fsProbe.Add("dir", types.NewGenerator("is directory", "List[Bool]", func() types.Probe {
		paths, _ := filepath.Glob(input)
		listProbe := types.NewList("")
		for _, path := range paths {
			file := getFileWrapper(path)
			listProbe.Add(types.NewBool("", file.isDir()))
		}
		return listProbe
	}))
	fsProbe.Add("regular", types.NewGenerator("is regular file", "List[Bool]", func() types.Probe {
		paths, _ := filepath.Glob(input)
		listProbe := types.NewList("")
		for _, path := range paths {
			file := getFileWrapper(path)
			listProbe.Add(types.NewBool("", file.isRegular()))
		}
		return listProbe
	}))
	fsProbe.Add("uid", types.NewGenerator("UID", "List[Number]", func() types.Probe {
		paths, _ := filepath.Glob(input)
		listProbe := types.NewList("")
		for _, path := range paths {
			file := getFileWrapper(path)
			listProbe.Add(types.NewNumber("", file.getUID(), "int"))
		}
		return listProbe
	}))
	fsProbe.Add("gid", types.NewGenerator("GID", "List[Number]", func() types.Probe {
		paths, _ := filepath.Glob(input)
		listProbe := types.NewList("")
		for _, path := range paths {
			file := getFileWrapper(path)
			listProbe.Add(types.NewNumber("", file.getGID(), "int"))
		}
		return listProbe
	}))
	fsProbe.Add("user", types.NewGenerator("Username", "List[String]", func() types.Probe {
		paths, _ := filepath.Glob(input)
		listProbe := types.NewList("")
		for _, path := range paths {
			file := getFileWrapper(path)
			listProbe.Add(types.NewString("", file.getUsername()))
		}
		return listProbe
	}))
	fsProbe.Add("group", types.NewGenerator("Groupname", "List[String]", func() types.Probe {
		paths, _ := filepath.Glob(input)
		listProbe := types.NewList("")
		for _, path := range paths {
			file := getFileWrapper(path)
			listProbe.Add(types.NewString("", file.getGroupname()))
		}
		return listProbe
	}))
	fsProbe.Add("count", types.NewGenerator("files count", types.NumberType, func() types.Probe {
		paths, _ := filepath.Glob(input)
		return types.NewNumber("", uint64(len(paths)), "int")
	}))
	fsProbe.Add("exists", types.NewGenerator("at least one file found", types.BoolType, func() types.Probe {
		paths, _ := filepath.Glob(input)
		return types.NewBool("", len(paths) > 0)
	}))
	log.Debugf("fs probe initialized for '%s'", input)
	return fsProbe, ""
}
