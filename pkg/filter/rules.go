package filter

// RuleResult is the result of a rule execution
type RuleResult uint8

const (
	// ResultFalse means the Rule did fail
	ResultFalse RuleResult = iota
	// ResultTrue means the Rule succeeded
	ResultTrue
	// ResultDoesNotApply means the Rule was not applicable for the given query
	ResultDoesNotApply
)

// Rule is a function that returns a result depending on the given query and value
type Rule func(query string, value int64) RuleResult

// AsBool converts result into boolean, keep in mind that a rule that did not apply
// will always return true
func (res RuleResult) AsBool() bool {
	if !res.DoesApply() {
		return true
	}
	return res == ResultTrue
}

// DoesApply checks if the rule did apply to the rule
func (res RuleResult) DoesApply() bool {
	return res != ResultDoesNotApply
}

// Invert the boolean result, ResultDoesNotApply will stay the same though
func (res RuleResult) Invert() RuleResult {
	if res == ResultDoesNotApply {
		return ResultDoesNotApply
	}

	if res == ResultFalse {
		return ResultTrue
	}

	return ResultFalse
}
