package filter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSONArrayFilter matches lines where a JSON array field contains a given value.
type JSONArrayFilter struct {
	field     string
	value     string
	ignoreCase bool
}

// NewJSONArrayFilter creates a filter that checks if a JSON array field contains value.
// field must be a top-level JSON key whose value is a JSON array of strings.
func NewJSONArrayFilter(field, value string, ignoreCase bool) (*JSONArrayFilter, error) {
	if field == "" {
		return nil, fmt.Errorf("%w: field", ErrEmptyParam)
	}
	if value == "" {
		return nil, fmt.Errorf("%w: value", ErrEmptyParam)
	}
	return &JSONArrayFilter{field: field, value: value, ignoreCase: ignoreCase}, nil
}

// Match returns true if the line is valid JSON and the named array field contains value.
func (f *JSONArrayFilter) Match(line string) bool {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return false
	}
	raw, ok := obj[f.field]
	if !ok {
		return false
	}
	var items []string
	if err := json.Unmarshal(raw, &items); err != nil {
		return false
	}
	needle := f.value
	if f.ignoreCase {
		needle = strings.ToLower(needle)
	}
	for _, item := range items {
		candidate := item
		if f.ignoreCase {
			candidate = strings.ToLower(candidate)
		}
		if candidate == needle {
			return true
		}
	}
	return false
}

// Field returns the JSON field name.
func (f *JSONArrayFilter) Field() string { return f.field }

// Value returns the target value.
func (f *JSONArrayFilter) Value() string { return f.value }

// IgnoreCase returns whether matching is case-insensitive.
func (f *JSONArrayFilter) IgnoreCase() bool { return f.ignoreCase }
