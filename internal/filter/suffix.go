package filter

import "strings"

// SuffixFilter matches lines that end with a given suffix.
type SuffixFilter struct {
	suffix          string
	caseInsensitive bool
}

// NewSuffixFilter creates a SuffixFilter for the given suffix.
// Returns ErrEmptyPattern if suffix is empty.
func NewSuffixFilter(suffix string, caseInsensitive bool) (*SuffixFilter, error) {
	if suffix == "" {
		return nil, ErrEmptyPattern
	}
	if caseInsensitive {
		suffix = strings.ToLower(suffix)
	}
	return &SuffixFilter{
		suffix:          suffix,
		caseInsensitive: caseInsensitive,
	}, nil
}

// Match returns true if line ends with the configured suffix.
func (f *SuffixFilter) Match(line string) bool {
	if f.caseInsensitive {
		return strings.HasSuffix(strings.ToLower(line), f.suffix)
	}
	return strings.HasSuffix(line, f.suffix)
}

// Suffix returns the configured suffix string.
func (f *SuffixFilter) Suffix() string {
	return f.suffix
}

// CaseInsensitive returns whether the filter is case-insensitive.
func (f *SuffixFilter) CaseInsensitive() bool {
	return f.caseInsensitive
}
