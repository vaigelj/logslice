package filter

import "fmt"

// LineNumFilter matches lines whose 1-based line number falls within
// an inclusive [Min, Max] range. Either bound may be 0 to indicate
// "no limit" in that direction.
type LineNumFilter struct {
	min int
	max int
	current int
}

// NewLineNumFilter creates a LineNumFilter. min and max are 1-based
// line numbers. Pass 0 for min to start from the first line; pass 0
// for max to read until EOF.
func NewLineNumFilter(min, max int) (*LineNumFilter, error) {
	if min < 0 {
		return nil, fmt.Errorf("%w: min line number must be >= 0, got %d", ErrInvalidArgument, min)
	}
	if max < 0 {
		return nil, fmt.Errorf("%w: max line number must be >= 0, got %d", ErrInvalidArgument, max)
	}
	if min > 0 && max > 0 && min > max {
		return nil, fmt.Errorf("%w: min (%d) must not exceed max (%d)", ErrInvalidArgument, min, max)
	}
	return &LineNumFilter{min: min, max: max}, nil
}

// Match increments the internal counter and returns true when the
// current line number is within [min, max].
func (f *LineNumFilter) Match(_ string) bool {
	f.current++
	if f.min > 0 && f.current < f.min {
		return false
	}
	if f.max > 0 && f.current > f.max {
		return false
	}
	return true
}

// Min returns the configured lower bound.
func (f *LineNumFilter) Min() int { return f.min }

// Max returns the configured upper bound.
func (f *LineNumFilter) Max() int { return f.max }

// Current returns the number of lines seen so far.
func (f *LineNumFilter) Current() int { return f.current }
