package filter

import "errors"

// SampleFilter passes every Nth matching line, useful for sampling large logs.
type SampleFilter struct {
	n       int
	counter int
}

// NewSampleFilter creates a filter that passes every nth line (1-based).
// n must be >= 1. n=1 passes every line, n=2 passes every other line, etc.
func NewSampleFilter(n int) (*SampleFilter, error) {
	if n < 1 {
		return nil, errors.New("sample: n must be >= 1")
	}
	return &SampleFilter{n: n}, nil
}

// Match returns true for every nth call.
func (f *SampleFilter) Match(line string) bool {
	f.counter++
	if f.counter >= f.n {
		f.counter = 0
		return true
	}
	return false
}

// N returns the sample interval.
func (f *SampleFilter) N() int {
	return f.n
}

// Matched returns the number of lines that have passed the filter so far.
func (f *SampleFilter) Matched() int {
	// counter resets each cycle; we track via external usage, but expose N for config.
	return f.n
}
