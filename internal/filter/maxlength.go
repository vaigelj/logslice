package filter

import "fmt"

// MaxLengthFilter rejects lines whose byte length exceeds a maximum.
type MaxLengthFilter struct {
	max int
}

// NewMaxLengthFilter creates a MaxLengthFilter that passes lines with length
// less than or equal to max. max must be non-negative.
func NewMaxLengthFilter(max int) (*MaxLengthFilter, error) {
	if max < 0 {
		return nil, fmt.Errorf("%w: max length must be non-negative, got %d", ErrInvalidArgument, max)
	}
	return &MaxLengthFilter{max: max}, nil
}

// Match returns true when the line length is within the allowed maximum.
func (f *MaxLengthFilter) Match(line string) bool {
	return len(line) <= f.max
}

// Max returns the configured maximum line length.
func (f *MaxLengthFilter) Max() int {
	return f.max
}
