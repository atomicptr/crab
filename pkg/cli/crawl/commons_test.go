package crawl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidUrl(t *testing.T) {
	table := []struct {
		Url   string
		Valid bool
	}{
		{"https://github.com", true},
		{"", false},
		{"http://domain.com", true},
		{"https://domain.com/with/path/yes", true},
	}

	for _, row := range table {
		assert.Equal(t, row.Valid, isValidUrl(row.Url), row.Url)
	}
}
