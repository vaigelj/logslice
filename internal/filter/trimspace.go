package filter

import "strings"

// TrimSpaceFilter is a transformer that trims leading and/or trailing
// whitespace from each line before passing it downstream.
type TrimSpaceFilter struct {
	trimLeft  bool
	trimRight bool
}

// NewTrimSpaceFilter creates a TrimSpaceFilter.
// At least one of trimLeft or trimRight must be true; otherwise
// ErrInvalidArgument is returned.
func NewTrimSpaceFilter(trimLeft, trimRight bool) (*TrimSpaceFilter, error) {
	if !trimLeft && !trimRight {
		return nil, ErrInvalidArgument
	}
	return &TrimSpaceFilter{trimLeft: trimLeft, trimRight: trimRight}, nil
}

// Match always returns true — this filter only transforms lines.
func (f *TrimSpaceFilter) Match(_ string) bool { return true }

// Transform trims whitespace according to the filter's configuration.
func (f *TrimSpaceFilter) Transform(line string) string {
	switch {
	case f.trimLeft && f.trimRight:
		return strings.TrimSpace(line)
	case f.trimLeft:
		return strings.TrimLeftFunc(line, func(r rune) bool {
			return r == ' ' || r == '\t'
		})
	default:
		return strings.TrimRightFunc(line, func(r rune) bool {
			return r == ' ' || r == '\t'
		})
	}
}

// TrimLeft reports whether leading whitespace is trimmed.
func (f *TrimSpaceFilter) TrimLeft() bool { return f.trimLeft }

// TrimRight reports whether trailing whitespace is trimmed.
func (f *TrimSpaceFilter) TrimRight() bool { return f.trimRight }
