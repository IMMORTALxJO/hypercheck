package fs

import (
	"probe/options"
	"strconv"
)

func test_exists(files_glob string, files []FileWrapper, opts options.Option) (bool, string) {
	if len(files) == 0 {
		return false, "no files found"
	}
	return true, strconv.Itoa(len(files)) + " files found"
}

var test_exists_cases_pos = [][]string{
	[]string{"./assets/file_1kb", "exists"},
	[]string{"./assets/file_*", "exists"},
	[]string{"./assets/dir_10kb", "exists"},
}
var test_exists_cases_neg = [][]string{
	[]string{"./assets/file_not_found", "exists"},
	[]string{"./assets/notfound*", "exists"},
}
