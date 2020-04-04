package crawl

import (
	"github.com/atomicptr/crab/pkg/crawler"
	"github.com/atomicptr/crab/pkg/meta"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddUserAgentToRequest(t *testing.T) {
	modifier := crawler.RequestModifier{}
	modifier.With(addUserAgentToRequest())

	req := httptest.NewRequest("GET", "/", strings.NewReader(""))
	modifier.Do(req)

	assert.NotEmpty(t, req.Header.Get("User-Agent"))
	assert.Equal(t, meta.UserAgent(), req.Header.Get("User-Agent"))
}

func TestAddPrefixUrlToRequest(t *testing.T) {
	prefixUrl := "https://atomicptr.de"

	modifier := crawler.RequestModifier{}
	modifier.With(addPrefixUrlToRequest(prefixUrl))

	urls := []string{
		"https://example.com/page-1",
		"https://example.com/page-2",
		"https://example.com/page-3",
		"https://example.com/page-4",
		"https://example.com/page-5",
		"https://example.co.uk/page-1",
		"https://example.super.long.suburl.domain.com/page-1",
	}

	for _, url := range urls {
		req := httptest.NewRequest("GET", url, strings.NewReader(""))
		modifier.Do(req)

		assert.True(t, strings.HasPrefix(req.URL.String(), prefixUrl))
	}
}

func TestAddPrefixUrlToRequestWithInvalidPrefixUrl(t *testing.T) {
	prefixUrl := "https://this is not a valid url"

	modifier := crawler.RequestModifier{}
	modifier.With(addPrefixUrlToRequest(prefixUrl))

	req := httptest.NewRequest("GET", "https://example.com", strings.NewReader(""))
	modifier.Do(req)

	assert.Equal(t, "https://example.com", req.URL.String())
}

func TestAddCookiesToRequest(t *testing.T) {
	modifier := crawler.RequestModifier{}
	modifier.With(addCookiesToRequest(crawlerFlagOptions{
		CookieStrings: []string{
			"a=b",
			"b=c",
			"test=asdf",
		},
	}))

	req := httptest.NewRequest("GET", "/", strings.NewReader(""))
	modifier.Do(req)

	cookie, err := req.Cookie("a")
	assert.Nil(t, err)
	assert.Equal(t, "b", cookie.Value)

	cookie, err = req.Cookie("b")
	assert.Nil(t, err)
	assert.Equal(t, "c", cookie.Value)

	cookie, err = req.Cookie("test")
	assert.Nil(t, err)
	assert.Equal(t, "asdf", cookie.Value)
}

func TestAddHeadersToRequest(t *testing.T) {
	modifier := crawler.RequestModifier{}
	modifier.With(addHeadersToRequest(crawlerFlagOptions{
		HeaderStrings: []string{
			"a=b",
			"b=c",
			"test=asdf",
		},
	}))

	req := httptest.NewRequest("GET", "/", strings.NewReader(""))
	modifier.Do(req)

	assert.NotEmpty(t, req.Header.Get("a"))
	assert.NotEmpty(t, req.Header.Get("b"))
	assert.NotEmpty(t, req.Header.Get("test"))
	assert.Equal(t, "b", req.Header.Get("a"))
	assert.Equal(t, "c", req.Header.Get("b"))
	assert.Equal(t, "asdf", req.Header.Get("test"))
}
