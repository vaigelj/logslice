package filter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSONPathValueFilter matches lines where a dot-notation JSON path equals an expected value.
type JSONPathValueFilter struct {
	path          string
	expected      string
	caseInsensitive bool
}

// NewJSONPathValueFilter creates a filter that checks a dot-notation JSON path against an expected value.
func NewJSONPathValueFilter(path, expected string, caseInsensitive bool) (*JSONPathValueFilter, error) {
	if path == "" {
		return nil, fmt.Errorf("%w: path", ErrEmptyParam)
	}
	if expected == "" {
		return nil, fmt.Errorf("%w: expected", ErrEmptyParam)
	}
	return &JSONPathValueFilter{
		path:          path,
		expected:      expected,
		caseInsensitive: caseInsensitive,
	}, nil
}

// Match returns true if the JSON path in the line resolves to the expected value.
func (f *JSONPathValueFilter) Match(line string) bool {
	var root map[string]interface{}
	if err := json.Unmarshal([]byte(line), &root); err != nil {
		return false
	}
	val, ok := resolvePath(root, strings.Split(f.path, "."))
	if !ok {
		return false
	}
	actual := fmt.Sprintf("%v", val)
	if f.caseInsensitive {
		return strings.EqualFold(actual, f.expected)
	}
	return actual == f.expected
}

// Path returns the configured dot-notation JSON path.
func (f *JSONPathValueFilter) Path() string { return f.path }

// Expected returns the expected value string.
func (f *JSONPathValueFilter) Expected() string { return f.expected }

// CaseInsensitive returns whether matching is case-insensitive.
func (f *JSONPathValueFilter) CaseInsensitive() bool { return f.caseInsensitive }

// resolvePath walks a nested map following the key segments.
func resolvePath(node map[string]interface{}, segments []string) (interface{}, bool) {
	if len(segments) == 0 {
		return nil, false
	}
	val, ok := node[segments[0]]
	if !ok {
		return nil, false
	}
	if len(segments) == 1 {
		return val, true
	}
	child, ok := val.(map[string]interface{})
	if !ok {
		return nil, false
	}
	return resolvePath(child, segments[1:])
}
