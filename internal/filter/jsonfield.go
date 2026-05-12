package filter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSONFieldFilter matches lines where a JSON field equals an expected value.
type JSONFieldFilter struct {
	field         string
	expected      string
	caseInsensitive bool
}

// NewJSONFieldFilter creates a filter that matches lines where the given
// top-level JSON field equals the expected value.
// Returns an error if field or expected is empty.
func NewJSONFieldFilter(field, expected string, caseInsensitive bool) (*JSONFieldFilter, error) {
	if field == "" {
		return nil, fmt.Errorf("%w: json field name must not be empty", ErrInvalidFilter)
	}
	if expected == "" {
		return nil, fmt.Errorf("%w: json expected value must not be empty", ErrInvalidFilter)
	}
	return &JSONFieldFilter{
		field:           field,
		expected:        expected,
		caseInsensitive: caseInsensitive,
	}, nil
}

// Match returns true when the line is valid JSON and the specified field
// matches the expected value.
func (f *JSONFieldFilter) Match(line string) bool {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return false
	}
	val, ok := obj[f.field]
	if !ok {
		return false
	}
	actual := fmt.Sprintf("%v", val)
	if f.caseInsensitive {
		return strings.EqualFold(actual, f.expected)
	}
	return actual == f.expected
}

// Transform returns the line unchanged.
func (f *JSONFieldFilter) Transform(line string) string { return line }

// Field returns the JSON field name being matched.
func (f *JSONFieldFilter) Field() string { return f.field }

// Expected returns the expected value.
func (f *JSONFieldFilter) Expected() string { return f.expected }

// CaseInsensitive reports whether matching is case-insensitive.
func (f *JSONFieldFilter) CaseInsensitive() bool { return f.caseInsensitive }
