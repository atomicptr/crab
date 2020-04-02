package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsBool(t *testing.T) {
	assert.True(t, ResultTrue.AsBool())
	assert.False(t, ResultFalse.AsBool())
	assert.True(t, ResultDoesNotApply.AsBool())
}

func TestDoesApply(t *testing.T) {
	assert.False(t, ResultDoesNotApply.DoesApply())
	assert.True(t, ResultTrue.DoesApply())
	assert.True(t, ResultFalse.DoesApply())
}

func TestInvert(t *testing.T) {
	assert.Equal(t, ResultTrue, ResultFalse.Invert())
	assert.Equal(t, ResultFalse, ResultTrue.Invert())
	assert.Equal(t, ResultDoesNotApply, ResultDoesNotApply.Invert())
}
