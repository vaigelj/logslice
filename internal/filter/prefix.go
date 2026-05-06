package filter

import "strings"

// PrefixFilter matches lines that begin with a given string prefix.
type PrefixFilter struct {
	prefix     string
	caseFold   bool
	normalized string
}

// NewPrefixFilter creates a PrefixFilter for the given prefix.
// If caseFold is true, matching is case-insensitive.
// Returns ErrEmptyPattern if prefix is empty.
func NewPrefixFilter(prefix string, caseFold bool) (*PrefixFilter, error) {
	if prefix == "" {
		return nil, ErrEmptyPattern
	}
	norm := prefix
	if caseFold {
		norm = strings.ToLower(prefix)
	}
	return &PrefixFilter{
		prefix:     prefix,
		caseFold:   caseFold,
		normalized: norm,
	}, nil
}

// Match returns true if line starts with the configured prefix.
func (f *PrefixFilter) Match(line string) bool {
	if f.caseFold {
		return strings.HasPrefix(strings.ToLower(line), f.normalized)
	}
	return strings.HasPrefix(line, f.prefix)
}

// Prefix returns the raw prefix string.
func (f *PrefixFilter) Prefix() string { return f.prefix }

// CaseFold returns whether case-insensitive matching is enabled.
func (f *PrefixFilter) CaseFold() bool { return f.caseFold }
