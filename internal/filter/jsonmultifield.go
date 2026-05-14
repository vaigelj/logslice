package filter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSONMultiFieldFilter matches lines where all specified JSON field=value pairs are satisfied.
type JSONMultiFieldFilter struct {
	pairs           []jsonFieldPair
	caseInsensitive bool
}

type jsonFieldPair struct {
	field    string
	expected string
}

// NewJSONMultiFieldFilter creates a filter that requires all given field=value pairs to match.
// pairs is a slice of "field=value" strings.
func NewJSONMultiFieldFilter(pairs []string, caseInsensitive bool) (*JSONMultiFieldFilter, error) {
	if len(pairs) == 0 {
		return nil, fmt.Errorf("%w: at least one field=value pair required", ErrInvalidArgument)
	}
	parsed := make([]jsonFieldPair, 0, len(pairs))
	for _, p := range pairs {
		idx := strings.Index(p, "=")
		if idx <= 0 {
			return nil, fmt.Errorf("%w: invalid pair %q, expected field=value", ErrInvalidArgument, p)
		}
		field := p[:idx]
		value := p[idx+1:]
		if value == "" {
			return nil, fmt.Errorf("%w: empty value for field %q", ErrInvalidArgument, field)
		}
		parsed = append(parsed, jsonFieldPair{field: field, expected: value})
	}
	return &JSONMultiFieldFilter{pairs: parsed, caseInsensitive: caseInsensitive}, nil
}

// Match returns true if all field=value pairs are satisfied in the JSON line.
func (f *JSONMultiFieldFilter) Match(line string) bool {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return false
	}
	for _, p := range f.pairs {
		val, ok := obj[p.field]
		if !ok {
			return false
		}
		actual := fmt.Sprintf("%v", val)
		expected := p.expected
		if f.caseInsensitive {
			actual = strings.ToLower(actual)
			expected = strings.ToLower(expected)
		}
		if actual != expected {
			return false
		}
	}
	return true
}

// Transform returns the line unchanged.
func (f *JSONMultiFieldFilter) Transform(line string) string { return line }

// Pairs returns the parsed field=value pairs.
func (f *JSONMultiFieldFilter) Pairs() []jsonFieldPair { return f.pairs }

// CaseInsensitive returns whether matching is case-insensitive.
func (f *JSONMultiFieldFilter) CaseInsensitive() bool { return f.caseInsensitive }
