package filter

import (
	"regexp"
	"strconv"
	"strings"
)

// RuleIsValue that checks if a value is present
func RuleIsValue(query string, value int64) RuleResult {
	queryValue, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		return ResultDoesNotApply
	}

	if queryValue == value {
		return ResultTrue
	}

	return ResultFalse
}

// RuleIsNotValue that checks if a given value is not present
func RuleIsNotValue(query string, value int64) RuleResult {
	if !strings.HasPrefix(query, "!") {
		return ResultDoesNotApply
	}
	return RuleIsValue(query[1:], value).Invert()
}

// RuleIsInRange checks if the value is within a given range
func RuleIsInRange(query string, value int64) RuleResult {
	re := regexp.MustCompile("([0-9]+)-([0-9]+)")

	parts := re.FindStringSubmatch(query)

	// parts should be [query rangeStart rangeEnd], if it isn't it does not apply to us
	if len(parts) != 3 {
		return ResultDoesNotApply
	}

	rangeStart, _ := strconv.ParseInt(parts[1], 10, 64)
	rangeEnd, _ := strconv.ParseInt(parts[2], 10, 64)

	if rangeStart <= value && value <= rangeEnd {
		return ResultTrue
	}

	return ResultFalse
}

// RuleIsNotInRange checks if the value is not within range
func RuleIsNotInRange(query string, value int64) RuleResult {
	if !strings.HasPrefix(query, "!") {
		return ResultDoesNotApply
	}
	return RuleIsInRange(query[1:], value).Invert()
}

// RuleIsGreaterThan checks if the value is greater than the given query value
func RuleIsGreaterThan(query string, value int64) RuleResult {
	if !strings.HasPrefix(query, ">") {
		return ResultDoesNotApply
	}

	queryValue, err := strconv.ParseInt(query[1:], 10, 64)
	if err != nil {
		return ResultDoesNotApply
	}

	if value > queryValue {
		return ResultTrue
	}

	return ResultFalse
}

// RuleIsSmallerThan checks if the value is smaller than the given query value
func RuleIsSmallerThan(query string, value int64) RuleResult {
	if !strings.HasPrefix(query, "<") {
		return ResultDoesNotApply
	}

	queryValue, err := strconv.ParseInt(query[1:], 10, 64)
	if err != nil {
		return ResultDoesNotApply
	}

	if value < queryValue {
		return ResultTrue
	}

	return ResultFalse
}
