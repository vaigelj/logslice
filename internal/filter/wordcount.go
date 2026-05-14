package filter

import (
	"fmt"
	"strings"
)

// WordCountFilter matches lines whose word count falls within [min, max].
// A word is any whitespace-delimited token.
type WordCountFilter struct {
	min int
	max int
}

// NewWordCountFilter creates a WordCountFilter.
// min must be >= 0, max must be >= min (0 means no upper bound).
func NewWordCountFilter(min, max int) (*WordCountFilter, error) {
	if min < 0 {
		return nil, fmt.Errorf("%w: min word count must be >= 0", ErrInvalidArgument)
	}
	if max != 0 && max < min {
		return nil, fmt.Errorf("%w: max word count must be >= min", ErrInvalidArgument)
	}
	return &WordCountFilter{min: min, max: max}, nil
}

// Match returns true when the line's word count is within the configured range.
func (f *WordCountFilter) Match(line string) bool {
	count := len(strings.Fields(line))
	if count < f.min {
		return false
	}
	if f.max != 0 && count > f.max {
		return false
	}
	return true
}

// Transform returns the line unchanged.
func (f *WordCountFilter) Transform(line string) string { return line }

// Min returns the configured minimum word count.
func (f *WordCountFilter) Min() int { return f.min }

// Max returns the configured maximum word count (0 = unlimited).
func (f *WordCountFilter) Max() int { return f.max }
