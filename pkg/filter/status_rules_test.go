package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRuleIsValue(t *testing.T) {
	assert.Equal(t, ResultTrue, RuleIsValue("200", 200))
	assert.Equal(t, ResultFalse, RuleIsValue("200", 404))
	assert.Equal(t, ResultDoesNotApply, RuleIsValue("!200", 200))
	assert.Equal(t, ResultDoesNotApply, RuleIsValue("!200", 404))
}

func TestRuleIsNotValue(t *testing.T) {
	assert.Equal(t, ResultFalse, RuleIsNotValue("!200", 200))
	assert.Equal(t, ResultTrue, RuleIsNotValue("!200", 404))
	assert.Equal(t, ResultDoesNotApply, RuleIsNotValue("200", 200))
	assert.Equal(t, ResultDoesNotApply, RuleIsNotValue("200", 404))
}

func TestRuleIsInRange(t *testing.T) {
	assert.Equal(t, ResultTrue, RuleIsInRange("200-299", 200))
	assert.Equal(t, ResultTrue, RuleIsInRange("200-299", 218))
	assert.Equal(t, ResultTrue, RuleIsInRange("200-299", 299))
	assert.Equal(t, ResultFalse, RuleIsInRange("200-299", 404))
	assert.Equal(t, ResultDoesNotApply, RuleIsInRange("200", 200))
	assert.Equal(t, ResultDoesNotApply, RuleIsInRange("200-", 200))
	assert.Equal(t, ResultDoesNotApply, RuleIsInRange("-200", 200))
}

func TestRuleIsNotInRange(t *testing.T) {
	assert.Equal(t, ResultTrue, RuleIsNotInRange("!200-299", 404))
	assert.Equal(t, ResultTrue, RuleIsNotInRange("!200-299", 199))
	assert.Equal(t, ResultTrue, RuleIsNotInRange("!200-299", 300))
	assert.Equal(t, ResultFalse, RuleIsNotInRange("!200-299", 200))
	assert.Equal(t, ResultFalse, RuleIsNotInRange("!200-299", 218))
	assert.Equal(t, ResultFalse, RuleIsNotInRange("!200-299", 299))
	assert.Equal(t, ResultDoesNotApply, RuleIsNotInRange("200-299", 200))
	assert.Equal(t, ResultDoesNotApply, RuleIsNotInRange("200-299", 218))
	assert.Equal(t, ResultDoesNotApply, RuleIsNotInRange("200-299", 299))
	assert.Equal(t, ResultDoesNotApply, RuleIsNotInRange("200-299", 404))
	assert.Equal(t, ResultDoesNotApply, RuleIsNotInRange("!200", 200))
	assert.Equal(t, ResultDoesNotApply, RuleIsNotInRange("!200-", 200))
	assert.Equal(t, ResultDoesNotApply, RuleIsNotInRange("!-200", 200))
}

func TestRuleIsGreaterThan(t *testing.T) {
	assert.Equal(t, ResultTrue, RuleIsGreaterThan(">200", 201))
	assert.Equal(t, ResultFalse, RuleIsGreaterThan(">200", 200))
	assert.Equal(t, ResultDoesNotApply, RuleIsGreaterThan("200", 200))
	assert.Equal(t, ResultDoesNotApply, RuleIsGreaterThan("<200", 200))
	assert.Equal(t, ResultDoesNotApply, RuleIsGreaterThan("200-299", 200))
}

func TestRuleIsGreaterThanInvalidQuery(t *testing.T) {
	assert.Equal(t, ResultDoesNotApply, RuleIsGreaterThan(">test", 200))
}

func TestRuleIsSmallerThanInvalidQuery(t *testing.T) {
	assert.Equal(t, ResultDoesNotApply, RuleIsSmallerThan("<test", 200))
}

func TestRuleIsSmallerThan(t *testing.T) {
	assert.Equal(t, ResultTrue, RuleIsSmallerThan("<400", 399))
	assert.Equal(t, ResultFalse, RuleIsSmallerThan("<400", 400))
	assert.Equal(t, ResultDoesNotApply, RuleIsSmallerThan("200", 200))
	assert.Equal(t, ResultDoesNotApply, RuleIsSmallerThan(">200", 200))
	assert.Equal(t, ResultDoesNotApply, RuleIsSmallerThan("200-299", 200))
}
