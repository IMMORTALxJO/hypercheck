package http

import (
	"net/http"
	"probe/options"
	"probe/probes"

	log "github.com/sirupsen/logrus"
)

const ProbeName = "http"

type HttpResult struct {
	Url  string
	Resp *http.Response
	Err  error
}

var httpResultsCache = map[string]*HttpResult{}

var Checks = map[string]probes.Check{
	"code": probes.Check{
		Name: "code",
		Option: &options.Condition{
			Name:       "code",
			ParserType: "PosInt",
			OpsList:    []string{"<", ">", "=", "<=", ">=", "=="},
		},
		Tests: []string{"code"},
	},
	"online": probes.Check{
		Name: "online",
		Option: &options.Common{
			Name: "online",
		},
		Tests: []string{"online"},
	},
	"offline": probes.Check{
		Name: "offline",
		Option: &options.Common{
			Name: "offline",
		},
		Tests: []string{"offline"},
	},
	"length": probes.Check{
		Name: "length",
		Option: &options.Condition{
			Name:       "length",
			ParserType: "PosInt",
			OpsList:    []string{"<", ">", "=", "<=", ">=", "=="},
		},
		Tests: []string{"length"},
	},
}

var Tests = map[string]func(string, *HttpResult, options.Option) (bool, string){
	"online":  test_online,
	"offline": test_offline,
	"code":    test_code,
	"length":  test_length,
}

func Probe(URL string, opts string) bool {
	checks := probes.OptionsToChecks(opts, Checks)
	checkResult := Collector(URL)
	probeResult := true
	for _, check := range checks {
		for _, testName := range check.Tests {
			test, ok := Tests[testName]
			if !ok {
				log.Fatalf("%s: Probe has no '%s' test", ProbeName, check.Name)
			}
			result, message := test(URL, checkResult, check.Option)
			if result {
				log.Debugf("%s['%s'].%s.success: %s", ProbeName, URL, testName, message)
			} else {
				log.Errorf("%s['%s'].%s.failed: %s", ProbeName, URL, testName, message)
				probeResult = false
			}
		}
	}
	return probeResult
}

func Collector(URL string) *HttpResult {
	if _, ok := httpResultsCache[URL]; !ok {
		resp, err := http.Get(URL)
		log.Debugf("http: Get '%s'", URL)
		httpResultsCache[URL] = &HttpResult{
			URL,
			resp,
			err,
		}
	}
	return httpResultsCache[URL]
}
