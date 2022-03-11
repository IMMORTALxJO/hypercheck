package fs

import (
	"probe/options"

	log "github.com/sirupsen/logrus"
)

func test_dir(files_glob string, files []FileWrapper, opt options.Option) (bool, string) {
	for _, file := range files {
		if !file.IsDir() {
			return false, "path '" + file.Name() + "' is not a directory"
		}
		log.Debugf("path '%s' is directory", file.Name())
	}
	return true, "each path is a directory"
}

var test_dir_cases_pos = [][]string{
	[]string{"./assets/dir_10kb", "dir"},
}
var test_dir_cases_neg = [][]string{
	[]string{"./assets/file_1kb", "dir"},
	[]string{"./assets/*", "dir"},
}
