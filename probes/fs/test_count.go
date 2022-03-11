package fs

import (
	"probe/options"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func test_count(files_glob string, files []FileWrapper, opt options.Option) (bool, string) {
	checkOp := opt.GetOperation()
	checkValue := opt.GetValueInt()
	filesCount := uint64(len(files))
	compStr := strconv.FormatUint(filesCount, 10) + checkOp + opt.GetValue()
	errMessage := "files count ! " + compStr
	if !options.CompareIntWithOp(filesCount, checkValue, checkOp) {
		return false, errMessage
	}
	log.Debug("files count is ok")
	return true, "files count " + compStr
}

var test_count_cases_pos = [][]string{
	[]string{"./assets/file_1kb", "count==1"},
	[]string{"./assets/file_*", "count==4"},
}
var test_count_cases_neg = [][]string{
	[]string{"./assets/file_1kb", "count==2"},
	[]string{"./assets/file_*", "count>4"},
}
