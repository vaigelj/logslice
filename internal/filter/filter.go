// Package filter defines the Filter interface and common filter primitives
// used throughout logslice to decide which log lines are retained.
package filter

// Filter is the core interface implemented by every filter type.
// Match returns true when the given line should be included in the output.
type Filter interface {
	Match(line string) bool
}

// MatchAll is a no-op Filter that accepts every line.
// It is useful as a default when no filtering is required.
type MatchAll struct{}

// Match always returns true.
func (MatchAll) Match(_ string) bool { return true }

// MatchNone is a Filter that rejects every line.
// Primarily useful in tests.
type MatchNone struct{}

// Match always returns false.
func (MatchNone) Match(_ string) bool { return false }
