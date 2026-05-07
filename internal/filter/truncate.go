package filter

import "fmt"

// TruncateFilter truncates lines that exceed a maximum byte length,
// appending an optional suffix to indicate truncation.
type TruncateFilter struct {
	maxLen int
	suffix string
}

// NewTruncateFilter creates a TruncateFilter that truncates lines longer than
// maxLen bytes. The suffix (e.g. "...") is appended after truncation.
// maxLen must be positive and must be >= len(suffix).
func NewTruncateFilter(maxLen int, suffix string) (*TruncateFilter, error) {
	if maxLen <= 0 {
		return nil, fmt.Errorf("%w: maxLen must be positive, got %d", ErrInvalidArgument, maxLen)
	}
	if len(suffix) > maxLen {
		return nil, fmt.Errorf("%w: suffix length %d exceeds maxLen %d", ErrInvalidArgument, len(suffix), maxLen)
	}
	return &TruncateFilter{maxLen: maxLen, suffix: suffix}, nil
}

// Match always returns true but mutates the line content if it exceeds maxLen.
// Because Filter.Match takes a string (immutable), truncation is exposed via
// the Transform method instead; Match simply reports whether the line is long
// enough to need truncation (useful when used standalone).
func (f *TruncateFilter) Match(line string) bool {
	return true
}

// Transform returns the (possibly truncated) version of line.
func (f *TruncateFilter) Transform(line string) string {
	if len(line) <= f.maxLen {
		return line
	}
	cutAt := f.maxLen - len(f.suffix)
	return line[:cutAt] + f.suffix
}

// MaxLen returns the configured maximum line length.
func (f *TruncateFilter) MaxLen() int { return f.maxLen }

// Suffix returns the truncation suffix string.
func (f *TruncateFilter) Suffix() string { return f.suffix }
