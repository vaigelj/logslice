package filter

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// JSONRegexFilter matches lines where a JSON field's value matches a regular expression.
type JSONRegexFilter struct {
	field   string
	pattern string
	re      *regexp.Regexp
}

// NewJSONRegexFilter creates a filter that matches lines where the given JSON
// field's string value matches the provided regular expression pattern.
func NewJSONRegexFilter(field, pattern string) (*JSONRegexFilter, error) {
	if field == "" {
		return nil, fmt.Errorf("%w: field", ErrEmptyParam)
	}
	if pattern == "" {
		return nil, fmt.Errorf("%w: pattern", ErrEmptyParam)
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %w", err)
	}
	return &JSONRegexFilter{field: field, pattern: pattern, re: re}, nil
}

// Match returns true if the line is valid JSON and the specified field's value
// matches the compiled regular expression.
func (f *JSONRegexFilter) Match(line string) bool {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return false
	}
	val, ok := obj[f.field]
	if !ok {
		return false
	}
	str, ok := val.(string)
	if !ok {
		str = fmt.Sprintf("%v", val)
	}
	return f.re.MatchString(str)
}

// Field returns the JSON field name being matched.
func (f *JSONRegexFilter) Field() string { return f.field }

// Pattern returns the regular expression pattern.
func (f *JSONRegexFilter) Pattern() string { return f.pattern }
