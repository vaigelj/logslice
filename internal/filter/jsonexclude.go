package filter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSONExcludeFilter rejects lines where a JSON field matches a given value.
type JSONExcludeFilter struct {
	field           string
	expected        string
	caseInsensitive bool
}

// NewJSONExcludeFilter creates a filter that rejects lines where the named
// JSON field equals expected. Returns ErrEmptyField or ErrEmptyValue if
// either argument is blank.
func NewJSONExcludeFilter(field, expected string, caseInsensitive bool) (*JSONExcludeFilter, error) {
	if field == "" {
		return nil, ErrEmptyField
	}
	if expected == "" {
		return nil, ErrEmptyValue
	}
	return &JSONExcludeFilter{
		field:           field,
		expected:        expected,
		caseInsensitive: caseInsensitive,
	}, nil
}

// Match returns false (rejects) when the JSON field equals the expected value.
func (f *JSONExcludeFilter) Match(line string) bool {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return true
	}
	val, ok := obj[f.field]
	if !ok {
		return true
	}
	got := fmt.Sprintf("%v", val)
	exp := f.expected
	if f.caseInsensitive {
		got = strings.ToLower(got)
		exp = strings.ToLower(exp)
	}
	return got != exp
}

// Field returns the JSON field name.
func (f *JSONExcludeFilter) Field() string { return f.field }

// Expected returns the value that triggers exclusion.
func (f *JSONExcludeFilter) Expected() string { return f.expected }

// CaseInsensitive reports whether matching is case-insensitive.
func (f *JSONExcludeFilter) CaseInsensitive() bool { return f.caseInsensitive }
