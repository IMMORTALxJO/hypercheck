package fs

import (
	"probe/options"

	log "github.com/sirupsen/logrus"
)

func test_size(files_glob string, files []FileWrapper, opt options.Option) (bool, string) {
	checkOp := opt.GetOperation()
	checkBytes := opt.GetValueBytes()
	for _, file := range files {
		fileSize := file.Size()
		log.Debugf("check size of '%s' %d %s %d", file.Name(), fileSize, checkOp, checkBytes)
		if !options.CompareIntWithOp(fileSize, checkBytes, checkOp) {
			return false, " total files size is wrong"
		}
		log.Debugf("'%s' size is ok", file.Name())
	}
	return true, "each file size " + checkOp + opt.GetValue()
}

var test_size_cases_pos = [][]string{
	[]string{"./assets/file_1kb", "size>0Kb"},
	[]string{"./assets/file_1kb", "size>=1Kb"},
	[]string{"./assets/file_1kb", "size==1Kb"},
	[]string{"./assets/file_1kb", "size<=1Kb"},
	[]string{"./assets/file_1kb", "size<2Kb"},
	[]string{"./assets/file_*", "size<=4Kb"},
	[]string{"./assets/dir_10kb", "size==10Kb"},
	[]string{"./assets/dir_5kb", "size==5Kb"},
	[]string{"./assets/dir_5kb/dir_3kb", "size==3Kb"},
	[]string{"./assets/deep_1kb", "size==1Kb"},
}
var test_size_cases_neg = [][]string{
	[]string{"./assets/file_1kb", "size>2Kb"},
	[]string{"./assets/file_1kb", "size>=2Kb"},
	[]string{"./assets/file_1kb", "size==2Kb"},
	[]string{"./assets/file_1kb", "size<=0Kb"},
	[]string{"./assets/file_1kb", "size<0Kb"},
	[]string{"./assets/file_*", "size<=2Kb"},
}
