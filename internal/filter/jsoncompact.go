package filter

import (
	"encoding/json"
	"strings"
)

// JSONCompactFilter re-encodes each line as compact (minified) JSON.
// Lines that are not valid JSON are passed through unchanged.
// Match always returns true; the transformation is applied via Transform.
type JSONCompactFilter struct {
	passthrough bool
}

// NewJSONCompactFilter creates a JSONCompactFilter.
// If passthrough is true, non-JSON lines are passed through unchanged;
// if false, non-JSON lines are dropped (Match returns false).
func NewJSONCompactFilter(passthrough bool) (*JSONCompactFilter, error) {
	return &JSONCompactFilter{passthrough: passthrough}, nil
}

// Passthrough returns whether non-JSON lines are passed through.
func (f *JSONCompactFilter) Passthrough() bool { return f.passthrough }

// Match returns true for valid JSON lines, or always true when passthrough is enabled.
func (f *JSONCompactFilter) Match(line string) bool {
	line = strings.TrimSpace(line)
	if !json.Valid([]byte(line)) {
		return f.passthrough
	}
	return true
}

// Transform re-encodes the line as compact JSON.
// If the line is not valid JSON, it is returned unchanged.
func (f *JSONCompactFilter) Transform(line string) string {
	trimmed := strings.TrimSpace(line)
	if !json.Valid([]byte(trimmed)) {
		return line
	}
	var raw json.RawMessage
	if err := json.Unmarshal([]byte(trimmed), &raw); err != nil {
		return line
	}
	out, err := json.Marshal(raw)
	if err != nil {
		return line
	}
	return string(out)
}
