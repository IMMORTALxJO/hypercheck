package fs

import (
	"probe/options"

	log "github.com/sirupsen/logrus"
)

func test_regular(files_glob string, files []FileWrapper, opts options.Option) (bool, string) {
	for _, file := range files {
		if !file.IsRegular() {
			return false, "path '" + file.Name() + "' is not a regular file"
		}
		log.Debugf("path '%s' is regular file", file.Name())
	}
	return true, "each path is a regular file"
}

var test_regular_cases_pos = [][]string{
	[]string{"./assets/file_1kb", "regular"},
	[]string{"./assets/file_*", "regular"},
}
var test_regular_cases_neg = [][]string{
	[]string{"./assets/*", "regular"},
	[]string{"./assets/dir_10kb", "regular"},
}
