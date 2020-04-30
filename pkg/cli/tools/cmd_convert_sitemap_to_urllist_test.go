package tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveBaseUrl(t *testing.T) {
	expected := map[string]string{
		"https://domain.com/test-url":                              "/test-url",
		"https://domain.com/test-url/":                             "/test-url/",
		"https://domain.com/test-url/test":                         "/test-url/test",
		"https://domain.com/test-url/test#test1234":                "/test-url/test#test1234",
		"https://domain.com/test-url/test?x=1234":                  "/test-url/test?x=1234",
		"https://domain.com/test-url/test?x=1234&y=12345":          "/test-url/test?x=1234&y=12345",
		"https://domain.com/test-url/test?x=1234&z=/test/asdf#yay": "/test-url/test?x=1234&z=/test/asdf#yay",
	}

	input := make([]string, len(expected))

	i := 0

	for k := range expected {
		input[i] = k
		i++
	}

	result := removeBaseUrls(input)

	for i, url := range result {
		assert.Equal(t, expected[input[i]], url)
	}
}
