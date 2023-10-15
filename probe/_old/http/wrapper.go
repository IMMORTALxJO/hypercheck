package http

import (
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type httpWrapper struct {
	Url           string
	content       string
	contentCached bool
	resp          *http.Response
	err           error
}

func (w *httpWrapper) getResp() *http.Response {
	return w.resp
}

func (w *httpWrapper) getError() error {
	if w.resp == nil {
		w.load()
	}
	return w.err
}

func (w *httpWrapper) load() {
	w.resp, w.err = http.Get(w.Url)
	log.Debugf("http.resp loaded %s", w.Url)
}

func (w *httpWrapper) getCode() uint64 {
	if !w.getOnline() {
		return uint64(0)
	}
	return uint64(w.getResp().StatusCode)
}

func (w *httpWrapper) getOnline() bool {
	err := w.getError()
	if err != nil {
		log.Error(err)
	}
	return err == nil
}

func (w *httpWrapper) getContent() string {
	if !w.getOnline() {
		return ""
	}
	if !w.contentCached {
		log.Debugf("no cache, w.contentCached=%v", w.contentCached)
		defer w.getResp().Body.Close()
		body, _ := io.ReadAll(w.getResp().Body)
		w.content = string(body)
		w.contentCached = true
	}
	log.Debugf("body: '%s'", w.content)
	return w.content
}

func (w *httpWrapper) getHeaders() []string {
	var headers []string
	if !w.getOnline() {
		return headers
	}
	for h, v := range w.getResp().Header {
		headers = append(headers, h+": "+v[0])
		log.Debugf("header: '%s: %s'", h, v[0])
	}
	log.Debugf("headers: '%s'", headers)
	return headers
}

var httpCache = map[string]*httpWrapper{}

func getHttpWrapper(url string) *httpWrapper {
	_, ok := httpCache[url]
	if !ok {
		httpCache[url] = &httpWrapper{
			Url: url,
		}
	} else {
		log.Debugf("got httpWrapper for %s from cache", url)
	}
	return httpCache[url]
}
