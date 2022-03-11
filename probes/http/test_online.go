package http

import (
	"probe/options"

	log "github.com/sirupsen/logrus"
)

func test_online(URL string, httpResult *HttpResult, opt options.Option) (bool, string) {
	if httpResult.Err != nil {
		return false, "url '" + URL + "' is not online"
	}
	log.Debugf("url '%s' is online", URL)
	return true, "url is online"
}

func test_offline(URL string, httpResult *HttpResult, opt options.Option) (bool, string) {
	if httpResult.Err == nil {
		return false, "url '" + URL + "' is not offline"
	}
	log.Debugf("url '%s' is offline", URL)
	return true, "url is offline"
}

var test_online_cases_pos = [][]string{
	[]string{"https://postman-echo.com/status/200", "online"},
	[]string{"https://offline-not-found.com", "offline"},
}

var test_online_cases_neg = [][]string{
	[]string{"https://offline-not-found.com", "online"},
	[]string{"https://postman-echo.com/status/200", "offline"},
}
