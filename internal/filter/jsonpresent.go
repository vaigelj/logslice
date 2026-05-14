package filter

import (
	"encoding/json"
	"strings"
)

// JSONPresentFilter matches lines where a given JSON field exists (is present),
// regardless of its value. Nested fields are supported via dot notation.
type JSONPresentFilter struct {
	field           string
	invertPresence  bool
}

// NewJSONPresentFilter creates a filter that matches lines containing the
// specified JSON field. If invertPresence is true, lines where the field is
// absent are matched instead.
func NewJSONPresentFilter(field string, invertPresence bool) (*JSONPresentFilter, error) {
	if strings.TrimSpace(field) == "" {
		return nil, ErrEmptyField
	}
	return &JSONPresentFilter{field: field, invertPresence: invertPresence}, nil
}

// Match returns true when the JSON field presence matches the expected
// configuration.
func (f *JSONPresentFilter) Match(line string) bool {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return false
	}
	present := lookupJSONField(obj, f.field)
	if f.invertPresence {
		return !present
	}
	return present
}

// Field returns the target JSON field path.
func (f *JSONPresentFilter) Field() string { return f.field }

// Inverted returns whether the filter matches absent fields.
func (f *JSONPresentFilter) Inverted() bool { return f.invertPresence }

// lookupJSONField traverses a nested JSON object using dot-separated keys.
// Returns true if the final key exists.
func lookupJSONField(obj map[string]interface{}, path string) bool {
	parts := strings.SplitN(path, ".", 2)
	val, ok := obj[parts[0]]
	if !ok {
		return false
	}
	if len(parts) == 1 {
		return true
	}
	nested, ok := val.(map[string]interface{})
	if !ok {
		return false
	}
	return lookupJSONField(nested, parts[1])
}
