package filter

import "fmt"

// SkipLinesFilter skips the first N lines and passes through everything after.
type SkipLinesFilter struct {
	skip int
	seen int
}

// NewSkipLinesFilter creates a filter that skips the first n lines.
// n must be greater than zero.
func NewSkipLinesFilter(n int) (*SkipLinesFilter, error) {
	if n <= 0 {
		return nil, fmt.Errorf("%w: skip must be greater than zero, got %d", ErrInvalidArgument, n)
	}
	return &SkipLinesFilter{skip: n}, nil
}

// Match returns false for the first n lines, then true for all subsequent lines.
func (f *SkipLinesFilter) Match(line string) bool {
	f.seen++
	return f.seen > f.skip
}

// Transform returns the line unchanged.
func (f *SkipLinesFilter) Transform(line string) string {
	return line
}

// Skip returns the number of lines to skip.
func (f *SkipLinesFilter) Skip() int {
	return f.skip
}

// Seen returns the total number of lines evaluated so far.
func (f *SkipLinesFilter) Seen() int {
	return f.seen
}
