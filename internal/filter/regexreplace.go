package filter

import (
	"fmt"
	"regexp"
)

// RegexReplaceFilter transforms lines by replacing all regex matches with a
// template string. Named capture groups may be referenced as $name.
// Match always returns true so every line passes through (possibly transformed).
type RegexReplaceFilter struct {
	pattern     string
	replacement string
	re          *regexp.Regexp
}

// NewRegexReplaceFilter constructs a RegexReplaceFilter.
// pattern must be a valid regular expression and non-empty.
// replacement may be empty (effectively deletes matches).
func NewRegexReplaceFilter(pattern, replacement string) (*RegexReplaceFilter, error) {
	if pattern == "" {
		return nil, ErrEmptyPattern
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidPattern, err)
	}
	return &RegexReplaceFilter{
		pattern:     pattern,
		replacement: replacement,
		re:          re,
	}, nil
}

// Match always returns true; transformation is applied via Transform.
func (f *RegexReplaceFilter) Match(_ string) bool { return true }

// Transform replaces all occurrences of the pattern in line with the
// replacement string, expanding capture group references.
func (f *RegexReplaceFilter) Transform(line string) string {
	return f.re.ReplaceAllString(line, f.replacement)
}

// Pattern returns the compiled pattern string.
func (f *RegexReplaceFilter) Pattern() string { return f.pattern }

// Replacement returns the replacement template string.
func (f *RegexReplaceFilter) Replacement() string { return f.replacement }
