package filter

import (
	"fmt"
	"regexp"
)

// MultiRegexFilter matches lines that match ALL provided regex patterns.
type MultiRegexFilter struct {
	patterns []*regexp.Regexp
	raw      []string
}

// NewMultiRegexFilter compiles all patterns and returns a filter that requires
// every pattern to match a line. Returns an error if any pattern is empty or
// invalid.
func NewMultiRegexFilter(patterns []string) (*MultiRegexFilter, error) {
	if len(patterns) == 0 {
		return nil, fmt.Errorf("%w: at least one pattern required", ErrEmptyPattern)
	}
	compiled := make([]*regexp.Regexp, 0, len(patterns))
	for _, p := range patterns {
		if p == "" {
			return nil, fmt.Errorf("%w: pattern must not be empty", ErrEmptyPattern)
		}
		re, err := regexp.Compile(p)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrInvalidPattern, err.Error())
		}
		compiled = append(compiled, re)
	}
	return &MultiRegexFilter{patterns: compiled, raw: patterns}, nil
}

// Match returns true only when every pattern matches the line.
func (f *MultiRegexFilter) Match(line string) bool {
	for _, re := range f.patterns {
		if !re.MatchString(line) {
			return false
		}
	}
	return true
}

// Transform returns the line unchanged.
func (f *MultiRegexFilter) Transform(line string) string { return line }

// Patterns returns the raw pattern strings.
func (f *MultiRegexFilter) Patterns() []string { return f.raw }
