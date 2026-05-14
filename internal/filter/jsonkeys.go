package filter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSONKeysFilter matches lines where a JSON object contains all of the required keys.
type JSONKeysFilter struct {
	keys            []string
	caseInsensitive bool
}

// NewJSONKeysFilter returns a filter that accepts JSON lines containing all specified keys.
// keys must be non-empty. caseInsensitive applies to key comparison.
func NewJSONKeysFilter(keys []string, caseInsensitive bool) (*JSONKeysFilter, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("%w: keys list must not be empty", ErrInvalidArgument)
	}
	normalized := make([]string, len(keys))
	for i, k := range keys {
		if k == "" {
			return nil, fmt.Errorf("%w: key at index %d must not be empty", ErrInvalidArgument, i)
		}
		if caseInsensitive {
			normalized[i] = strings.ToLower(k)
		} else {
			normalized[i] = k
		}
	}
	return &JSONKeysFilter{keys: normalized, caseInsensitive: caseInsensitive}, nil
}

// Match returns true if the line is valid JSON and contains all required keys.
func (f *JSONKeysFilter) Match(line string) bool {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return false
	}
	for _, required := range f.keys {
		found := false
		for k := range obj {
			candidate := k
			if f.caseInsensitive {
				candidate = strings.ToLower(k)
			}
			if candidate == required {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Keys returns the normalized required keys.
func (f *JSONKeysFilter) Keys() []string { return f.keys }

// CaseInsensitive returns whether key matching is case-insensitive.
func (f *JSONKeysFilter) CaseInsensitive() bool { return f.caseInsensitive }
