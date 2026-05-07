package filter

import (
	"strings"
)

// ExcludeAnyFilter rejects lines that contain any of the given terms.
type ExcludeAnyFilter struct {
	terms         []string
	caseInsensitive bool
}

// NewExcludeAnyFilter creates a filter that rejects lines containing any of
// the provided terms. Returns ErrNoTerms if terms is empty.
func NewExcludeAnyFilter(terms []string, caseInsensitive bool) (*ExcludeAnyFilter, error) {
	if len(terms) == 0 {
		return nil, ErrNoTerms
	}
	normalized := make([]string, len(terms))
	for i, t := range terms {
		if caseInsensitive {
			normalized[i] = strings.ToLower(t)
		} else {
			normalized[i] = t
		}
	}
	return &ExcludeAnyFilter{terms: normalized, caseInsensitive: caseInsensitive}, nil
}

// Match returns false if the line contains any of the excluded terms.
func (f *ExcludeAnyFilter) Match(line string) bool {
	candidate := line
	if f.caseInsensitive {
		candidate = strings.ToLower(line)
	}
	for _, t := range f.terms {
		if strings.Contains(candidate, t) {
			return false
		}
	}
	return true
}

// Terms returns the normalised exclusion terms.
func (f *ExcludeAnyFilter) Terms() []string { return f.terms }

// CaseInsensitive reports whether matching is case-insensitive.
func (f *ExcludeAnyFilter) CaseInsensitive() bool { return f.caseInsensitive }
