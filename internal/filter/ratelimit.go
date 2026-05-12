package filter

import (
	"fmt"
	"time"
)

// RateLimitFilter passes at most N lines per duration window.
// Once the window expires a new window begins and the counter resets.
type RateLimitFilter struct {
	maxPerWindow int
	window       time.Duration
	count        int
	windowStart  time.Time
}

// NewRateLimitFilter creates a RateLimitFilter that allows at most
// maxPerWindow lines within each rolling window of the given duration.
func NewRateLimitFilter(maxPerWindow int, window time.Duration) (*RateLimitFilter, error) {
	if maxPerWindow <= 0 {
		return nil, fmt.Errorf("%w: maxPerWindow must be > 0, got %d", ErrInvalidArgument, maxPerWindow)
	}
	if window <= 0 {
		return nil, fmt.Errorf("%w: window must be > 0, got %s", ErrInvalidArgument, window)
	}
	return &RateLimitFilter{
		maxPerWindow: maxPerWindow,
		window:       window,
		windowStart:  time.Now(),
	}, nil
}

// Match returns true if the line is within the current rate limit window.
func (f *RateLimitFilter) Match(line string) bool {
	now := time.Now()
	if now.Sub(f.windowStart) >= f.window {
		f.windowStart = now
		f.count = 0
	}
	if f.count < f.maxPerWindow {
		f.count++
		return true
	}
	return false
}

// MaxPerWindow returns the configured maximum lines per window.
func (f *RateLimitFilter) MaxPerWindow() int { return f.maxPerWindow }

// Window returns the configured duration window.
func (f *RateLimitFilter) Window() time.Duration { return f.window }

// Count returns the number of lines matched in the current window.
func (f *RateLimitFilter) Count() int { return f.count }
