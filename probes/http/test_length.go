package http

import (
	"probe/options"

	log "github.com/sirupsen/logrus"
)

func test_length(URL string, httpResult *HttpResult, opt options.Option) (bool, string) {
	if !options.CompareIntWithOp(uint64(httpResult.Resp.ContentLength), opt.GetValueInt(), opt.GetOperation()) {
		return false, " content length is wrong"
	}

	log.Debugf("url '%s' content length is wrong", URL)
	return true, "url content length is ok"
}

var test_length_cases_pos = [][]string{
	[]string{"https://postman-echo.com/status/200", "length==14"},
}

var test_length_cases_neg = [][]string{
	[]string{"https://postman-echo.com/status/200", "length!=14"},
}
