package filter

import (
	"fmt"
	"strings"
	"time"
)

// TimestampShiftFilter rewrites a timestamp found at the start of each line
// by adding (or subtracting) a fixed duration. Lines that do not contain a
// parseable timestamp are passed through unchanged.
type TimestampShiftFilter struct {
	layout string
	shift  time.Duration
}

// NewTimestampShiftFilter returns a filter that parses timestamps using
// layout and shifts them by shift. layout must be non-empty.
func NewTimestampShiftFilter(layout string, shift time.Duration) (*TimestampShiftFilter, error) {
	if strings.TrimSpace(layout) == "" {
		return nil, ErrEmptyLayout
	}
	return &TimestampShiftFilter{layout: layout, shift: shift}, nil
}

// Match always returns true; the filter is used for its Transform side-effect.
func (f *TimestampShiftFilter) Match(_ string) bool { return true }

// Transform rewrites the leading timestamp in line if one is found.
func (f *TimestampShiftFilter) Transform(line string) string {
	if len(line) < len(f.layout) {
		return line
	}
	candidate := line[:len(f.layout)]
	t, err := time.Parse(f.layout, candidate)
	if err != nil {
		return line
	}
	shifted := t.Add(f.shift)
	return fmt.Sprintf("%s%s", shifted.Format(f.layout), line[len(f.layout):])
}

// Layout returns the timestamp layout string.
func (f *TimestampShiftFilter) Layout() string { return f.layout }

// Shift returns the configured duration shift.
func (f *TimestampShiftFilter) Shift() time.Duration { return f.shift }
