package filter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSONTruncateFilter truncates a string field in a JSON log line to a maximum length.
type JSONTruncateFilter struct {
	field   string
	maxLen  int
	suffix  string
}

// NewJSONTruncateFilter creates a filter that truncates the given JSON string field.
func NewJSONTruncateFilter(field string, maxLen int, suffix string) (*JSONTruncateFilter, error) {
	if field == "" {
		return nil, fmt.Errorf("%w: field must not be empty", ErrInvalidArgument)
	}
	if maxLen <= 0 {
		return nil, fmt.Errorf("%w: maxLen must be positive", ErrInvalidArgument)
	}
	if len(suffix) >= maxLen {
		return nil, fmt.Errorf("%w: suffix length must be less than maxLen", ErrInvalidArgument)
	}
	return &JSONTruncateFilter{field: field, maxLen: maxLen, suffix: suffix}, nil
}

// Match always returns true; this filter only transforms.
func (f *JSONTruncateFilter) Match(line string) bool { return true }

// Transform truncates the specified JSON field value if it exceeds maxLen.
func (f *JSONTruncateFilter) Transform(line string) string {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return line
	}
	val, ok := obj[f.field]
	if !ok {
		return line
	}
	str, ok := val.(string)
	if !ok {
		return line
	}
	if len(str) > f.maxLen {
		truncated := str[:f.maxLen-len(f.suffix)] + f.suffix
		obj[f.field] = truncated
	}
	out, err := json.Marshal(obj)
	if err != nil {
		return line
	}
	return strings.TrimSpace(string(out))
}

// Field returns the target JSON field name.
func (f *JSONTruncateFilter) Field() string { return f.field }

// MaxLen returns the configured maximum length.
func (f *JSONTruncateFilter) MaxLen() int { return f.maxLen }

// Suffix returns the truncation suffix.
func (f *JSONTruncateFilter) Suffix() string { return f.suffix }
