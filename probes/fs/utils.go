package fs

import (
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"

	log "github.com/sirupsen/logrus"
)

var filesCache = map[string]FileWrapper{}

type FileWrapper struct {
	Path         string
	Info         os.FileInfo
	stat         *syscall.Stat_t
	statComputed bool
	size         uint64
	sizeCounted  bool
}

func (w FileWrapper) IsRegular() bool {
	mode := w.Info.Mode()
	return mode.IsRegular()
}

func (w FileWrapper) IsDir() bool {
	return w.Info.IsDir()
}

func (w *FileWrapper) Size() uint64 {
	if w.sizeCounted {
		return w.size
	}
	if w.IsDir() {
		w.size = getDirSize(w)
	} else {
		w.size = uint64(w.Info.Size())
	}
	w.sizeCounted = true
	return w.size
}

func (w *FileWrapper) Stat() *syscall.Stat_t {
	if w.statComputed {
		return w.stat
	}
	w.stat = w.Info.Sys().(*syscall.Stat_t)
	return w.stat
}

func (w FileWrapper) Name() string {
	return w.Path
}

func (w FileWrapper) getUID() uint64 {
	return uint64(w.Stat().Uid)
}
func (w FileWrapper) getGID() uint64 {
	return uint64(w.Stat().Gid)
}
func (w FileWrapper) getUsername() string {
	usr, _ := user.LookupId(strconv.FormatUint(w.getUID(), 10))
	return usr.Name
}
func (w FileWrapper) getGroupname() string {
	grp, _ := user.LookupGroupId(strconv.FormatUint(w.getGID(), 10))
	return grp.Name
}

func getFileWrapper(path string) FileWrapper {
	_, ok := filesCache[path]
	if !ok {
		fileInfo, _ := os.Stat(path)
		filesCache[path] = FileWrapper{
			Path: path,
			Info: fileInfo,
		}
	} else {
		log.Debugf("got FileWrapper for %s from cache", path)
	}
	return filesCache[path]
}

func getDirSize(d *FileWrapper) uint64 {
	paths, _ := filepath.Glob(d.Path + "/*")
	sumSize := uint64(0)
	for _, path := range paths {
		file := getFileWrapper(path)
		sumSize += file.Size()
		log.Debugf("getDirSize(%s) => %d", d.Path, sumSize)
	}
	log.Debugf("getDirSize(%s) = %d", d.Path, sumSize)
	return sumSize
}
