package filter

import "strings"

// SubstringFilter matches lines that contain a given substring.
type SubstringFilter struct {
	substring string
	caseInsensitive bool
	normalized string
}

// NewSubstringFilter creates a SubstringFilter for the given substring.
// Returns ErrEmptyPattern if substring is empty.
func NewSubstringFilter(substring string, caseInsensitive bool) (*SubstringFilter, error) {
	if substring == "" {
		return nil, ErrEmptyPattern
	}
	normalized := substring
	if caseInsensitive {
		normalized = strings.ToLower(substring)
	}
	return &SubstringFilter{
		substring:       substring,
		caseInsensitive: caseInsensitive,
		normalized:      normalized,
	}, nil
}

// Match returns true if the line contains the configured substring.
func (f *SubstringFilter) Match(line string) bool {
	if f.caseInsensitive {
		return strings.Contains(strings.ToLower(line), f.normalized)
	}
	return strings.Contains(line, f.substring)
}

// Substring returns the configured substring.
func (f *SubstringFilter) Substring() string { return f.substring }

// CaseInsensitive returns whether matching is case-insensitive.
func (f *SubstringFilter) CaseInsensitive() bool { return f.caseInsensitive }
