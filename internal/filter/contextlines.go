package filter

import (
	"errors"
)

// ContextLinesFilter keeps lines that fall within N lines before or after a
// matching anchor filter. It is stateful and must be used with a single pass.
type ContextLinesFilter struct {
	anchor  Filter
	before  int
	after   int
	buf     []string
	countdown int
}

// NewContextLinesFilter creates a filter that includes up to `before` lines
// preceding and `after` lines following any line matched by `anchor`.
func NewContextLinesFilter(anchor Filter, before, after int) (*ContextLinesFilter, error) {
	if anchor == nil {
		return nil, errors.New("contextlines: anchor filter must not be nil")
	}
	if before < 0 {
		return nil, errors.New("contextlines: before must be >= 0")
	}
	if after < 0 {
		return nil, errors.New("contextlines: after must be >= 0")
	}
	return &ContextLinesFilter{
		anchor:    anchor,
		before:    before,
		after:     after,
		buf:       make([]string, 0, before),
		countdown: 0,
	}, nil
}

// Match returns true if the line should be included based on context rules.
func (f *ContextLinesFilter) Match(line string) bool {
	// Roll the pre-context buffer.
	if f.before > 0 {
		f.buf = append(f.buf, line)
		if len(f.buf) > f.before+1 {
			f.buf = f.buf[1:]
		}
	}

	if f.anchor.Match(line) {
		f.countdown = f.after + 1
		return true
	}

	if f.countdown > 0 {
		f.countdown--
		return true
	}

	return false
}

// Before returns the configured pre-context line count.
func (f *ContextLinesFilter) Before() int { return f.before }

// After returns the configured post-context line count.
func (f *ContextLinesFilter) After() int { return f.after }

// Anchor returns the underlying anchor filter.
func (f *ContextLinesFilter) Anchor() Filter { return f.anchor }
