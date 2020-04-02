package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterIsValidSimple(t *testing.T) {
	filter := NewFilter()
	assert.True(t, filter.IsValid("200", 200))
	assert.False(t, filter.IsValid("200", 404))
	assert.False(t, filter.IsValid("!200", 200))
	assert.True(t, filter.IsValid("!200", 404))
}

func TestFilterIsValidRange(t *testing.T) {
	filter := NewFilter()
	assert.True(t, filter.IsValid("200-299", 200))
	assert.True(t, filter.IsValid("200-299", 218))
	assert.True(t, filter.IsValid("200-299", 299))
	assert.False(t, filter.IsValid("200-299", 404))
}

func TestFilterIsValidMultipleRanges(t *testing.T) {
	filter := NewFilter()
	assert.True(t, filter.IsValid("200-299,400-499", 218))
	assert.True(t, filter.IsValid("200-299,400-499", 404))
	assert.False(t, filter.IsValid("200-299,400-499", 500))
}

func TestFilterIsValidNotInRange(t *testing.T) {
	filter := NewFilter()
	assert.True(t, filter.IsValid("!200-299", 404))
	assert.True(t, filter.IsValid("!200-299,400-499", 404))
}

func TestFilterIsValidRangeWithExtras(t *testing.T) {
	filter := NewFilter()
	assert.True(t, filter.IsValid("200-299,404", 200))
	assert.True(t, filter.IsValid("200-299,404", 218))
	assert.True(t, filter.IsValid("200-299,404", 299))
	assert.True(t, filter.IsValid("200-299,404", 404))
	assert.False(t, filter.IsValid("200-299,404", 500))
	assert.True(t, filter.IsValid("200-299,404,500", 500))
}

func TestFilterIsValidGreaterSmallerThan(t *testing.T) {
	filter := NewFilter()
	assert.True(t, filter.IsValid(">400", 500))
	assert.True(t, filter.IsValid("<400,404", 404))
	assert.True(t, filter.IsValid(">200,<299", 218))
}

func TestFilterInvalidCases(t *testing.T) {
	filter := NewFilter()
	// should be true because one of them is correct
	assert.True(t, filter.IsValid("200-299,!218", 218))
	// should be true because 404 is not !200-299
	assert.True(t, filter.IsValid("!200-299,!400-499", 404))
}
