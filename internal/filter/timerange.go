package filter

import (
	"fmt"
	"time"
)

// TimeRangeFilter filters log lines based on a parsed timestamp within a time range.
type TimeRangeFilter struct {
	start  time.Time
	end    time.Time
	layout string
}

// NewTimeRangeFilter creates a new TimeRangeFilter.
// start and end are RFC3339 strings; layout is the timestamp format in log lines.
func NewTimeRangeFilter(start, end, layout string) (*TimeRangeFilter, error) {
	if layout == "" {
		return nil, fmt.Errorf("timerange: layout must not be empty")
	}

	var s, e time.Time
	var err error

	if start != "" {
		s, err = time.Parse(time.RFC3339, start)
		if err != nil {
			return nil, fmt.Errorf("timerange: invalid start time %q: %w", start, err)
		}
	}

	if end != "" {
		e, err = time.Parse(time.RFC3339, end)
		if err != nil {
			return nil, fmt.Errorf("timerange: invalid end time %q: %w", end, err)
		}
	}

	if !s.IsZero() && !e.IsZero() && e.Before(s) {
		return nil, fmt.Errorf("timerange: end time must not be before start time")
	}

	return &TimeRangeFilter{start: s, end: e, layout: layout}, nil
}

// Match returns true if the line contains a timestamp that falls within [start, end].
// Lines whose timestamps cannot be parsed are excluded.
func (f *TimeRangeFilter) Match(line string) bool {
	t, err := extractTime(line, f.layout)
	if err != nil {
		return false
	}
	if !f.start.IsZero() && t.Before(f.start) {
		return false
	}
	if !f.end.IsZero() && t.After(f.end) {
		return false
	}
	return true
}

// Start returns the filter's start boundary.
func (f *TimeRangeFilter) Start() time.Time { return f.start }

// End returns the filter's end boundary.
func (f *TimeRangeFilter) End() time.Time { return f.end }

// Layout returns the timestamp layout used for parsing.
func (f *TimeRangeFilter) Layout() string { return f.layout }

// extractTime attempts to parse a timestamp from the beginning of a log line.
func extractTime(line, layout string) (time.Time, error) {
	if len(line) < len(layout) {
		return time.Time{}, fmt.Errorf("line too short to contain timestamp")
	}
	return time.Parse(layout, line[:len(layout)])
}
