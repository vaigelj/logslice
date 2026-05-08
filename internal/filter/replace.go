package filter

import (
	"regexp"
	"strings"
)

// ReplaceFilter is a Transformer that replaces occurrences of a pattern
// (literal string or regex) within each line with a given replacement.
type ReplaceFilter struct {
	pattern     string
	replacement string
	caseInsensitive bool
	re          *regexp.Regexp
}

// NewReplaceFilter creates a ReplaceFilter that replaces all occurrences of
// pattern with replacement. If useRegex is true, pattern is compiled as a
// regular expression; otherwise a literal string replacement is performed.
func NewReplaceFilter(pattern, replacement string, caseInsensitive, useRegex bool) (*ReplaceFilter, error) {
	if pattern == "" {
		return nil, ErrEmptyPattern
	}

	var re *regexp.Regexp
	if useRegex {
		flags := ""
		if caseInsensitive {
			flags = "(?i)"
		}
		var err error
		re, err = regexp.Compile(flags + pattern)
		if err != nil {
			return nil, err
		}
	}

	return &ReplaceFilter{
		pattern:         pattern,
		replacement:     replacement,
		caseInsensitive: caseInsensitive,
		re:              re,
	}, nil
}

// Match always returns true — ReplaceFilter is a transformer, not a gatekeeper.
func (f *ReplaceFilter) Match(line string) bool { return true }

// Transform applies the replacement to the line and returns the result.
func (f *ReplaceFilter) Transform(line string) string {
	if f.re != nil {
		return f.re.ReplaceAllString(line, f.replacement)
	}
	if f.caseInsensitive {
		// case-insensitive literal replace via regex
		re := regexp.MustCompile("(?i)" + regexp.QuoteMeta(f.pattern))
		return re.ReplaceAllString(line, f.replacement)
	}
	return strings.ReplaceAll(line, f.pattern, f.replacement)
}

// Pattern returns the configured search pattern.
func (f *ReplaceFilter) Pattern() string { return f.pattern }

// Replacement returns the configured replacement string.
func (f *ReplaceFilter) Replacement() string { return f.replacement }
