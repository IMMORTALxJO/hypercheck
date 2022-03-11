package fs

import (
	"probe/options"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func test_uid(files_glob string, files []FileWrapper, opt options.Option) (bool, string) {
	checkValue := opt.GetValueInt()
	checkOp := opt.GetOperation()
	for _, file := range files {
		if !options.CompareIntWithOp(file.getUID(), checkValue, checkOp) {
			return false, "path '" + file.Name() + "' has not valid UID (" + strconv.Itoa(int(file.getUID())) + ")"
		}
		log.Debugf("path '%s' has right UID", file.Name())
	}
	return true, "each path has a right UID"
}

func test_gid(files_glob string, files []FileWrapper, opt options.Option) (bool, string) {
	checkValue := opt.GetValueInt()
	checkOp := opt.GetOperation()
	for _, file := range files {
		if !options.CompareIntWithOp(file.getGID(), checkValue, checkOp) {
			return false, "path '" + file.Name() + "' has not valid GID (" + strconv.Itoa(int(file.getGID())) + ")"
		}
		log.Debugf("path '%s' has right GID", file.Name())
	}
	return true, "each path has a right GID"
}

func test_username(files_glob string, files []FileWrapper, opt options.Option) (bool, string) {
	checkValue := opt.GetValue()
	checkOp := opt.GetOperation()
	for _, file := range files {
		username := file.getUsername()
		if !options.CompareStrWithOp(username, checkValue, checkOp) {
			return false, "path '" + file.Name() + "' has not valid User owner (" + username + ")"
		}
		log.Debugf("path '%s' has right User owner", file.Name())
	}
	return true, "each path has a right User owner"
}

func test_groupname(files_glob string, files []FileWrapper, opt options.Option) (bool, string) {
	checkValue := opt.GetValue()
	checkOp := opt.GetOperation()
	for _, file := range files {
		groupname := file.getGroupname()
		if !options.CompareStrWithOp(groupname, checkValue, checkOp) {
			return false, "path '" + file.Name() + "' has not valid Group owner (" + groupname + ")"
		}
		log.Debugf("path '%s' has right Group owner", file.Name())
	}
	return true, "each path has a right Group owner"
}

var test_owner_cases_pos = [][]string{
	[]string{"./assets/file_1kb", "uid==0"},
	[]string{"./assets/file_*", "uid!=1"},
	[]string{"./assets/file_*", "gid==0"},
	[]string{"./assets/*/*", "gid==0"},
	[]string{"./assets/*/*", "uid==0"},
	[]string{"./assets/*/*", "gid<10"},
}
var test_owner_cases_neg = [][]string{
	[]string{"./assets/file_1kb", "uid==1"},
	[]string{"./assets/file_*", "uid==1"},
	[]string{"./assets/file_*", "gid==1"},
	[]string{"./assets/*/*", "gid!=0"},
}
