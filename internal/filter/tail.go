package filter

import "fmt"

// TailFilter keeps only the last N matching lines by buffering them in a
// ring-buffer. Match always returns true so every line passes through the
// chain; the final slice is obtained via Lines() after all input is consumed.
type TailFilter struct {
	max    int
	buf    []string
	head   int // next write position (ring)
	count  int // total lines seen
}

// NewTailFilter creates a TailFilter that retains the last max lines.
// max must be >= 1.
func NewTailFilter(max int) (*TailFilter, error) {
	if max < 1 {
		return nil, fmt.Errorf("%w: tail max must be >= 1, got %d", ErrInvalidArgument, max)
	}
	return &TailFilter{
		max: max,
		buf: make([]string, max),
	}, nil
}

// Match records the line in the ring-buffer and always returns true so the
// line continues through the filter chain. The caller is responsible for
// collecting output via Lines() once all input has been processed.
func (f *TailFilter) Match(line string) bool {
	f.buf[f.head] = line
	f.head = (f.head + 1) % f.max
	f.count++
	return true
}

// Lines returns the last N lines in order. If fewer than max lines were seen
// all of them are returned.
func (f *TailFilter) Lines() []string {
	if f.count == 0 {
		return nil
	}
	if f.count <= f.max {
		out := make([]string, f.count)
		copy(out, f.buf[:f.count])
		return out
	}
	// ring is full; oldest entry is at f.head
	out := make([]string, f.max)
	copy(out, f.buf[f.head:])
	copy(out[f.max-f.head:], f.buf[:f.head])
	return out
}

// Max returns the configured maximum number of tail lines.
func (f *TailFilter) Max() int { return f.max }

// Seen returns the total number of lines that have been recorded.
func (f *TailFilter) Seen() int { return f.count }
