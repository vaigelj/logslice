package filter

import "errors"

// Sentinel errors shared across filter types.
var (
	// ErrNilFilter is returned when a nil Filter is provided where a concrete
	// implementation is required (e.g. NewInvertFilter).
	ErrNilFilter = errors.New("filter: inner filter must not be nil")

	// ErrEmptyLayout is returned when a time-range filter is constructed with
	// an empty time layout string.
	ErrEmptyLayout = errors.New("filter: time layout must not be empty")

	// ErrEndBeforeStart is returned when the end time precedes the start time
	// in a time-range filter.
	ErrEndBeforeStart = errors.New("filter: end time must not be before start time")
)
