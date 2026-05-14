package filter

import "errors"

// BookmarkFilter passes lines between a start pattern match and an end pattern
// match (inclusive). Once the end pattern is matched the filter resets and
// waits for the next start match.
type BookmarkFilter struct {
	start   *RegexFilter
	end     *RegexFilter
	active  bool
	matched int
}

// NewBookmarkFilter creates a BookmarkFilter that activates on startPattern
// and deactivates (inclusively) on endPattern.
func NewBookmarkFilter(startPattern, endPattern string) (*BookmarkFilter, error) {
	if startPattern == "" {
		return nil, errors.New("bookmark: start pattern must not be empty")
	}
	if endPattern == "" {
		return nil, errors.New("bookmark: end pattern must not be empty")
	}
	start, err := NewRegexFilter(startPattern)
	if err != nil {
		return nil, err
	}
	end, err := NewRegexFilter(endPattern)
	if err != nil {
		return nil, err
	}
	return &BookmarkFilter{start: start, end: end}, nil
}

// Match returns true when the line falls within an active bookmark region.
func (f *BookmarkFilter) Match(line string) bool {
	if !f.active && f.start.Match(line) {
		f.active = true
	}
	if !f.active {
		return false
	}
	f.matched++
	if f.end.Match(line) {
		f.active = false
	}
	return true
}

// Matched returns the number of lines passed through so far.
func (f *BookmarkFilter) Matched() int { return f.matched }

// StartPattern returns the compiled start regex pattern string.
func (f *BookmarkFilter) StartPattern() string { return f.start.Pattern() }

// EndPattern returns the compiled end regex pattern string.
func (f *BookmarkFilter) EndPattern() string { return f.end.Pattern() }
