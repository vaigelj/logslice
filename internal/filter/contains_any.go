package filter

import (
	"strings"
)

// ContainsAnyFilter matches lines that contain at least one of the given substrings.
type ContainsAnyFilter struct {
	terms           []string
	caseInsensitive bool
}

// NewContainsAnyFilter creates a filter that matches if a line contains any of the provided terms.
// Returns ErrEmptyValue if no terms are provided.
func NewContainsAnyFilter(terms []string, caseInsensitive bool) (*ContainsAnyFilter, error) {
	if len(terms) == 0 {
		return nil, ErrEmptyValue
	}
	normalized := make([]string, len(terms))
	for i, t := range terms {
		if caseInsensitive {
			normalized[i] = strings.ToLower(t)
		} else {
			normalized[i] = t
		}
	}
	return &ContainsAnyFilter{terms: normalized, caseInsensitive: caseInsensitive}, nil
}

// Match returns true if the line contains at least one of the filter terms.
func (f *ContainsAnyFilter) Match(line string) bool {
	candidate := line
	if f.caseInsensitive {
		candidate = strings.ToLower(line)
	}
	for _, t := range f.terms {
		if strings.Contains(candidate, t) {
			return true
		}
	}
	return false
}

// Terms returns the normalized list of terms used for matching.
func (f *ContainsAnyFilter) Terms() []string { return f.terms }

// CaseInsensitive reports whether the filter operates case-insensitively.
func (f *ContainsAnyFilter) CaseInsensitive() bool { return f.caseInsensitive }
