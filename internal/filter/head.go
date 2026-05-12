package filter

import "fmt"

// HeadFilter passes only the first N lines of input, regardless of other filters.
type HeadFilter struct {
	maxLines int
	seen     int
}

// NewHeadFilter creates a HeadFilter that passes only the first maxLines lines.
// maxLines must be >= 1.
func NewHeadFilter(maxLines int) (*HeadFilter, error) {
	if maxLines < 1 {
		return nil, fmt.Errorf("%w: head maxLines must be >= 1, got %d", ErrInvalidArgument, maxLines)
	}
	return &HeadFilter{maxLines: maxLines}, nil
}

// Match returns true while fewer than maxLines lines have been seen.
func (f *HeadFilter) Match(line string) bool {
	if f.seen >= f.maxLines {
		return false
	}
	f.seen++
	return true
}

// MaxLines returns the configured maximum number of lines.
func (f *HeadFilter) MaxLines() int { return f.maxLines }

// Seen returns the number of lines matched so far.
func (f *HeadFilter) Seen() int { return f.seen }
