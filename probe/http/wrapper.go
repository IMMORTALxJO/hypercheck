package http

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type httpWrapper struct {
	Url           string
	Resp          *http.Response
	Err           error
	content       string
	contentCached bool
}

func (w *httpWrapper) GetCode() uint64 {
	if !w.GetOnline() {
		return uint64(0)
	}
	return uint64(w.Resp.StatusCode)
}

func (w *httpWrapper) GetOnline() bool {
	return w.Err == nil
}

func (w *httpWrapper) GetContent() string {
	if !w.GetOnline() {
		return ""
	}
	if w.contentCached == false {
		log.Debugf("no cache, w.contentCached=%v", w.contentCached)
		defer w.Resp.Body.Close()
		body, _ := ioutil.ReadAll(w.Resp.Body)
		w.content = string(body)
		w.contentCached = true
	}
	log.Debugf("body: '%s'", w.content)
	return w.content
}

func (w *httpWrapper) GetHeaders() []string {
	var headers []string
	if !w.GetOnline() {
		return headers
	}
	for h, v := range w.Resp.Header {
		headers = append(headers, h+": "+v[0])
	}
	log.Debugf("headers: '%s'", headers)
	return headers
}

var httpCache = map[string]*httpWrapper{}

func getHttpWrapper(url string) *httpWrapper {
	_, ok := httpCache[url]
	if !ok {
		resp, err := http.Get(url)
		httpCache[url] = &httpWrapper{
			Url:  url,
			Resp: resp,
			Err:  err,
		}
	} else {
		log.Debugf("got httpWrapper for %s from cache", url)
	}
	return httpCache[url]
}
