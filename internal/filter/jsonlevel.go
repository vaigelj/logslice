package filter

import (
	"encoding/json"
	"strings"
)

// JSONLevelFilter matches log lines where a JSON field contains a value
// from a specified set of allowed levels (e.g. "error", "warn").
type JSONLevelFilter struct {
	field  string
	levels map[string]struct{}
	caseInsensitive bool
}

// NewJSONLevelFilter creates a filter that matches lines where the given JSON
// field equals one of the provided levels. Returns an error if field or levels
// are empty.
func NewJSONLevelFilter(field string, levels []string, caseInsensitive bool) (*JSONLevelFilter, error) {
	if field == "" {
		return nil, ErrEmptyField
	}
	if len(levels) == 0 {
		return nil, ErrEmptyValue
	}
	m := make(map[string]struct{}, len(levels))
	for _, l := range levels {
		key := l
		if caseInsensitive {
			key = strings.ToLower(l)
		}
		m[key] = struct{}{}
	}
	return &JSONLevelFilter{field: field, levels: m, caseInsensitive: caseInsensitive}, nil
}

// Match returns true if the JSON field value is one of the allowed levels.
func (f *JSONLevelFilter) Match(line string) bool {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return false
	}
	v, ok := obj[f.field]
	if !ok {
		return false
	}
	s, ok := v.(string)
	if !ok {
		return false
	}
	if f.caseInsensitive {
		s = strings.ToLower(s)
	}
	_, found := f.levels[s]
	return found
}

// Field returns the JSON field name being checked.
func (f *JSONLevelFilter) Field() string { return f.field }

// CaseInsensitive returns whether matching is case-insensitive.
func (f *JSONLevelFilter) CaseInsensitive() bool { return f.caseInsensitive }
