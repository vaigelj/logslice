package filter

import "fmt"

// MinLengthFilter rejects lines shorter than a minimum byte length.
type MinLengthFilter struct {
	minLen int
}

// NewMinLengthFilter creates a MinLengthFilter that passes lines with at least
// minLen bytes (not counting the trailing newline). Returns an error if minLen
// is negative.
func NewMinLengthFilter(minLen int) (*MinLengthFilter, error) {
	if minLen < 0 {
		return nil, fmt.Errorf("%w: minLen must be >= 0, got %d", ErrInvalidArgument, minLen)
	}
	return &MinLengthFilter{minLen: minLen}, nil
}

// Match returns true when the line length is >= the configured minimum.
func (f *MinLengthFilter) Match(line string) bool {
	return len(line) >= f.minLen
}

// MinLen returns the configured minimum length.
func (f *MinLengthFilter) MinLen() int {
	return f.minLen
}
