package fs

import (
	"probe/options"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func test_totalsize(files_glob string, files []FileWrapper, opt options.Option) (bool, string) {
	checkOp := opt.GetOperation()
	checkBytes := opt.GetValueBytes()
	sumSize := uint64(0)
	for _, file := range files {
		sumSize += file.Size()
		log.Debugf("total size counter => %d", sumSize)
	}
	log.Debugf("check total size %d %s %d", sumSize, checkOp, checkBytes)
	if !options.CompareIntWithOp(sumSize, checkBytes, checkOp) {
		return false, " total files size is wrong"
	}
	return true, "total size " + strconv.FormatUint(sumSize, 10) + checkOp + opt.GetValue()
}

var test_totalsize_cases_pos = [][]string{
	[]string{"./assets/file_1kb", "totalsize>0Kb"},
	[]string{"./assets/file_1kb", "totalsize>=1Kb"},
	[]string{"./assets/file_1kb", "totalsize==1Kb"},
	[]string{"./assets/file_1kb", "totalsize<=1Kb"},
	[]string{"./assets/file_1kb", "totalsize<2Kb"},
	[]string{"./assets/file_*", "totalsize==10Kb"},
}
var test_totalsize_cases_neg = [][]string{
	[]string{"./assets/file_1kb", "totalsize>2Kb"},
	[]string{"./assets/file_1kb", "totalsize>=2Kb"},
	[]string{"./assets/file_1kb", "totalsize==2Kb"},
	[]string{"./assets/file_1kb", "totalsize<=0Kb"},
	[]string{"./assets/file_1kb", "totalsize<0Kb"},
	[]string{"./assets/file_*", "totalsize==20Kb"},
}
