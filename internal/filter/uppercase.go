package filter

import "strings"

// UppercaseFilter transforms matching lines to uppercase.
// It always returns true for Match (transform-only filter).
type UppercaseFilter struct {
	mode uppercaseMode
}

type uppercaseMode int

const (
	// UppercaseFull converts the entire line to uppercase.
	UppercaseFull uppercaseMode = iota
	// UppercaseLower converts the entire line to lowercase.
	UppercaseLower
	// UppercaseTitle converts the line to title case.
	UppercaseTitle
)

// NewUppercaseFilter creates a new UppercaseFilter.
// mode must be one of UppercaseFull, UppercaseLower, or UppercaseTitle.
func NewUppercaseFilter(mode uppercaseMode) (*UppercaseFilter, error) {
	if mode != UppercaseFull && mode != UppercaseLower && mode != UppercaseTitle {
		return nil, ErrInvalidArgument
	}
	return &UppercaseFilter{mode: mode}, nil
}

// Match always returns true; this is a transform-only filter.
func (f *UppercaseFilter) Match(line string) bool {
	return true
}

// Transform applies the case transformation to the line.
func (f *UppercaseFilter) Transform(line string) string {
	switch f.mode {
	case UppercaseFull:
		return strings.ToUpper(line)
	case UppercaseLower:
		return strings.ToLower(line)
	case UppercaseTitle:
		return strings.ToTitle(line)
	default:
		return line
	}
}

// Mode returns the configured transformation mode.
func (f *UppercaseFilter) Mode() uppercaseMode {
	return f.mode
}
