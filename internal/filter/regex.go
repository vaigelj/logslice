package filter

import (
	"fmt"
	"regexp"
)

// RegexFilter holds a compiled regular expression used to match log lines.
type RegexFilter struct {
	pattern *regexp.Regexp
	invert  bool
}

// NewRegexFilter compiles the given pattern and returns a RegexFilter.
// If invert is true, lines that do NOT match the pattern are kept.
func NewRegexFilter(pattern string, invert bool) (*RegexFilter, error) {
	if pattern == "" {
		return nil, fmt.Errorf("regex pattern must not be empty")
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern %q: %w", pattern, err)
	}
	return &RegexFilter{pattern: re, invert: invert}, nil
}

// Match reports whether the given line should be kept according to the filter.
func (f *RegexFilter) Match(line string) bool {
	matched := f.pattern.MatchString(line)
	if f.invert {
		return !matched
	}
	return matched
}

// Pattern returns the string representation of the compiled regex.
func (f *RegexFilter) Pattern() string {
	return f.pattern.String()
}

// Inverted reports whether the filter operates in invert mode.
func (f *RegexFilter) Inverted() bool {
	return f.invert
}
