package crawl

import (
	"net/http"
	"net/url"

	"github.com/atomicptr/crab/pkg/crawler"
	"github.com/atomicptr/crab/pkg/meta"
)

func addUserAgentToRequest() crawler.RequestModifierFunc {
	return func(req *http.Request) {
		req.Header.Set("User-Agent", meta.UserAgent)
	}
}

func addPrefixUrlToRequest(prefixUrl string) crawler.RequestModifierFunc {
	return func(req *http.Request) {
		parsedPrefixUrl, err := url.Parse(prefixUrl)
		if err != nil {
			return
		}

		requestUrl, err := url.Parse(req.URL.String())
		if err != nil {
			return
		}

		requestUrl.Scheme = parsedPrefixUrl.Scheme
		requestUrl.Host = parsedPrefixUrl.Host

		req.URL = requestUrl
	}
}

func addCookiesToRequest(flagOptions crawlerFlagOptions) crawler.RequestModifierFunc {
	return func(req *http.Request) {
		for key, value := range flagOptions.CookieMap() {
			req.AddCookie(&http.Cookie{Name: key, Value: value})
		}
	}
}

func addHeadersToRequest(flagOptions crawlerFlagOptions) crawler.RequestModifierFunc {
	return func(req *http.Request) {
		for key, value := range flagOptions.HeaderMap() {
			req.Header.Set(key, value)
		}
	}
}
