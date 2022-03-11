package fs

import (
	"path/filepath"
	"probe/options"
	"probe/probes"

	log "github.com/sirupsen/logrus"
)

const ProbeName = "fs"

var Checks = map[string]probes.Check{
	"exists": probes.Check{
		Name: "exists",
		Option: &options.Common{
			Name: "exists",
		},
		Tests: []string{"exists"},
	},
	"size": probes.Check{
		Name: "size",
		Option: &options.Condition{
			Name:       "size",
			ParserType: "Bytes",
			OpsList:    []string{"<", ">", "=", "<=", ">=", "=="},
		},
		Tests: []string{"size"},
	},
	"totalsize": probes.Check{
		Name: "totalsize",
		Option: &options.Condition{
			Name:       "totalsize",
			ParserType: "Bytes",
			OpsList:    []string{"<", ">", "=", "<=", ">=", "=="},
		},
		Tests: []string{"totalsize"},
	},
	"count": probes.Check{
		Name: "count",
		Option: &options.Condition{
			Name:       "count",
			ParserType: "PosInt",
			OpsList:    []string{"<", ">", "=", "<=", ">=", "=="},
		},
		Tests: []string{"count"},
	},
	"regular": probes.Check{
		Name: "regular",
		Option: &options.Common{
			Name: "regular",
		},
		Tests: []string{"regular"},
	},
	"dir": probes.Check{
		Name: "dir",
		Option: &options.Common{
			Name: "dir",
		},
		Tests: []string{"dir"},
	},
	"uid": probes.Check{
		Name: "uid",
		Option: &options.Condition{
			Name:       "uid",
			ParserType: "PosInt",
			OpsList:    []string{"<", ">", "=", "<=", ">=", "=="},
		},
		Tests: []string{"uid"},
	},
	"gid": probes.Check{
		Name: "gid",
		Option: &options.Condition{
			Name:       "gid",
			ParserType: "PosInt",
			OpsList:    []string{"<", ">", "=", "<=", ">=", "=="},
		},
		Tests: []string{"gid"},
	},
}

var Tests = map[string]func(string, []FileWrapper, options.Option) (bool, string){
	"exists":    test_exists,
	"size":      test_size,
	"totalsize": test_totalsize,
	"regular":   test_regular,
	"dir":       test_dir,
	"count":     test_count,
	"uid":       test_uid,
	"gid":       test_gid,
}

func Probe(filesGlob string, opts string) bool {
	checks := probes.OptionsToChecks(opts, Checks)
	files := Collector(filesGlob)
	probeResult := true
	for _, check := range checks {
		for _, testName := range check.Tests {
			test, ok := Tests[testName]
			if !ok {
				log.Fatalf("%s: Probe has no '%s' test", ProbeName, check.Name)
			}
			result, message := test(filesGlob, files, check.Option)
			if result {
				log.Infof("%s['%s'].%s.success: %s", ProbeName, filesGlob, testName, message)
			} else {
				log.Errorf("%s['%s'].%s.failed: %s", ProbeName, filesGlob, testName, message)
				probeResult = false
			}
		}
	}
	return probeResult
}

func Collector(filesGlob string) []FileWrapper {
	paths, err := filepath.Glob(filesGlob)
	if err != nil {
		log.Fatalf("Wrong files glob pattern '%s'", filesGlob)
	}
	var files []FileWrapper
	for _, path := range paths {
		files = append(files, getFileWrapper(path))
	}
	return files
}
