package filter

import "sync"

// DedupeFilter rejects lines that have already been seen.
type DedupeFilter struct {
	mu   sync.Mutex
	seen map[string]struct{}
}

// NewDedupeFilter returns a DedupeFilter that discards duplicate lines.
func NewDedupeFilter() *DedupeFilter {
	return &DedupeFilter{
		seen: make(map[string]struct{}),
	}
}

// Match returns true if this is the first time the line has been seen.
func (f *DedupeFilter) Match(line string) bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, exists := f.seen[line]; exists {
		return false
	}
	f.seen[line] = struct{}{}
	return true
}

// Seen returns the number of unique lines recorded so far.
func (f *DedupeFilter) Seen() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.seen)
}
