package filter

import (
	"fmt"
	"strings"
)

// KVFieldFilter matches lines that contain a key=value pair where the value
// satisfies the expected string (optionally case-insensitive).
type KVFieldFilter struct {
	key           string
	expected      string
	sep           string
	caseInsensitive bool
}

// NewKVFieldFilter creates a filter that matches lines containing key<sep>value.
// sep defaults to "=" if empty.
func NewKVFieldFilter(key, sep, expected string, caseInsensitive bool) (*KVFieldFilter, error) {
	if key == "" {
		return nil, fmt.Errorf("%w: key must not be empty", ErrInvalidFilter)
	}
	if expected == "" {
		return nil, fmt.Errorf("%w: expected value must not be empty", ErrInvalidFilter)
	}
	if sep == "" {
		sep = "="
	}
	return &KVFieldFilter{
		key:           key,
		expected:      expected,
		sep:           sep,
		caseInsensitive: caseInsensitive,
	}, nil
}

// Match returns true if the line contains key<sep><expected>.
func (f *KVFieldFilter) Match(line string) bool {
	target := f.key + f.sep
	search := line
	exp := f.expected
	if f.caseInsensitive {
		search = strings.ToLower(line)
		target = strings.ToLower(target)
		exp = strings.ToLower(exp)
	}
	idx := strings.Index(search, target)
	if idx < 0 {
		return false
	}
	rest := search[idx+len(target):]
	// value ends at next whitespace or end of string
	end := strings.IndexAny(rest, " \t,;")
	var val string
	if end < 0 {
		val = rest
	} else {
		val = rest[:end]
	}
	return val == exp
}

// Transform returns the line unchanged.
func (f *KVFieldFilter) Transform(line string) string { return line }

// Key returns the configured key.
func (f *KVFieldFilter) Key() string { return f.key }

// Sep returns the configured separator.
func (f *KVFieldFilter) Sep() string { return f.sep }

// Expected returns the configured expected value.
func (f *KVFieldFilter) Expected() string { return f.expected }

// CaseInsensitive returns whether matching is case-insensitive.
func (f *KVFieldFilter) CaseInsensitive() bool { return f.caseInsensitive }
