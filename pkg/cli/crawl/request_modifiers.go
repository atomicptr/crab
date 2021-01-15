package crawl

import (
	"net/http"
	"net/url"

	"github.com/atomicptr/crab/pkg/crawler"
	"github.com/atomicptr/crab/pkg/meta"
)

func addUserAgentToRequest() crawler.RequestModifierFunc {
	return func(req *http.Request) {
		req.Header.Set("User-Agent", meta.UserAgent())
	}
}

func addPrefixUrlToRequest(prefixUrl string) crawler.RequestModifierFunc {
	return func(req *http.Request) {
		parsedPrefixUrl, err := url.Parse(prefixUrl)
		if err != nil || parsedPrefixUrl.String() == "" {
			// prefix url couldn't be parsed just abort
			return
		}

		requestUrl, _ := url.Parse(req.URL.String())

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

func addHttpBasicAuthToRequest(flagOptions crawlerFlagOptions) crawler.RequestModifierFunc {
	return func(req *http.Request) {
		req.SetBasicAuth(flagOptions.AuthUsername, flagOptions.AuthPassword)
	}
}
