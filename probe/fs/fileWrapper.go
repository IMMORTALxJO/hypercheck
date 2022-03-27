package fs

import (
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"

	log "github.com/sirupsen/logrus"
)

var filesCache = map[string]*fileWrapper{}

type fileWrapper struct {
	Path         string
	Info         os.FileInfo
	stat         *syscall.Stat_t
	statComputed bool
	size         uint64
	sizeCounted  bool
}

func (w *fileWrapper) IsRegular() bool {
	mode := w.Info.Mode()
	return mode.IsRegular()
}

func (w *fileWrapper) IsDir() bool {
	return w.Info.IsDir()
}

func (w *fileWrapper) getSize() uint64 {
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

func (w *fileWrapper) Stat() *syscall.Stat_t {
	if w.statComputed {
		return w.stat
	}
	w.stat = w.Info.Sys().(*syscall.Stat_t)
	return w.stat
}

func (w *fileWrapper) Name() string {
	return w.Path
}

func (w *fileWrapper) getUID() uint64 {
	return uint64(w.Stat().Uid)
}
func (w *fileWrapper) getGID() uint64 {
	return uint64(w.Stat().Gid)
}
func (w *fileWrapper) getUsername() string {
	usr, _ := user.LookupId(strconv.FormatUint(w.getUID(), 10))
	return usr.Name
}
func (w *fileWrapper) getGroupname() string {
	grp, _ := user.LookupGroupId(strconv.FormatUint(w.getGID(), 10))
	return grp.Name
}

func getFileWrapper(path string) *fileWrapper {
	_, ok := filesCache[path]
	if !ok {
		fileInfo, _ := os.Stat(path)
		filesCache[path] = &fileWrapper{
			Path: path,
			Info: fileInfo,
		}
	} else {
		log.Debugf("got fileWrapper for %s from cache", path)
	}
	return filesCache[path]
}

func getDirSize(d *fileWrapper) uint64 {
	paths, _ := filepath.Glob(d.Path + "/*")
	sumSize := uint64(0)
	for _, path := range paths {
		file := getFileWrapper(path)
		sumSize += file.getSize()
		log.Debugf("getDirSize(%s) => %d", d.Path, sumSize)
	}
	log.Debugf("getDirSize(%s) = %d", d.Path, sumSize)
	return sumSize
}
