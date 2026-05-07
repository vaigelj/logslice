package filter

import (
	"fmt"
	"sync/atomic"
)

// CountLimitFilter passes lines until a maximum match count is reached.
type CountLimitFilter struct {
	maxCount int64
	matched  atomic.Int64
}

// NewCountLimitFilter creates a filter that passes at most maxCount lines.
// Returns an error if maxCount is less than 1.
func NewCountLimitFilter(maxCount int64) (*CountLimitFilter, error) {
	if maxCount < 1 {
		return nil, fmt.Errorf("%w: maxCount must be >= 1, got %d", ErrInvalidArgument, maxCount)
	}
	return &CountLimitFilter{maxCount: maxCount}, nil
}

// Match returns true if the number of matched lines so far is less than maxCount.
func (f *CountLimitFilter) Match(line string) bool {
	current := f.matched.Add(1)
	return current <= f.maxCount
}

// MaxCount returns the configured maximum match count.
func (f *CountLimitFilter) MaxCount() int64 {
	return f.maxCount
}

// Matched returns the current number of lines that have been evaluated.
func (f *CountLimitFilter) Matched() int64 {
	return f.matched.Load()
}
