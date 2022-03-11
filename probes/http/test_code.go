package http

import (
	"probe/options"

	log "github.com/sirupsen/logrus"
)

func test_code(URL string, httpResult *HttpResult, opt options.Option) (bool, string) {
	if !options.CompareIntWithOp(uint64(httpResult.Resp.StatusCode), opt.GetValueInt(), opt.GetOperation()) {
		return false, " status code is wrong"
	}

	log.Debugf("url '%s' status code is wrong", URL)
	return true, "url status code is ok"
}

var test_code_cases_pos = [][]string{
	[]string{"https://postman-echo.com/status/200", "code==200"},
	[]string{"https://postman-echo.com/status/200", "code<500"},
	[]string{"https://postman-echo.com/status/200", "code!=502"},
	[]string{"https://postman-echo.com/basic-auth", "code==401"},
	[]string{"https://postman:password@postman-echo.com/basic-auth", "code==200"},
}

var test_code_cases_neg = [][]string{
	[]string{"https://postman-echo.com/status/502", "code==200"},
	[]string{"https://postman-echo.com/status/502", "code<=200"},
}
