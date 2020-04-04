package filter

import (
	"strings"
)

// Filter is a collection of rules
type Filter struct {
	Rules []Rule
}

// NewFilter returns a new filter with a set number of rules
func NewFilter() *Filter {
	filter := Filter{
		Rules: []Rule{
			RuleIsValue,
			RuleIsNotValue,
			RuleIsInRange,
			RuleIsNotInRange,
			RuleIsGreaterThan,
			RuleIsSmallerThan,
		},
	}
	return &filter
}

// IsValid checks if the query is valid for the given value
func (f *Filter) IsValid(query string, value int64) bool {
	// empty means no filter was supplied which is always good!
	if query == "" {
		return true
	}

	queries := strings.Split(query, ",")

	result := false

	for _, q := range queries {
		q = strings.TrimSpace(q)

		for _, rule := range f.Rules {
			res := rule(q, value)
			if res.DoesApply() {
				result = result || res.AsBool()
			}
		}
	}

	return result
}
