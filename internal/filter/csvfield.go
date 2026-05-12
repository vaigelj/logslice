package filter

import (
	"encoding/csv"
	"fmt"
	"strings"
)

// CSVFieldFilter matches lines where the CSV field at a given index equals the expected value.
type CSVFieldFilter struct {
	index     int
	expected  string
	ignoreCase bool
}

// NewCSVFieldFilter creates a filter that checks a specific CSV field.
// index is zero-based. Returns an error if index is negative or expected is empty.
func NewCSVFieldFilter(index int, expected string, ignoreCase bool) (*CSVFieldFilter, error) {
	if index < 0 {
		return nil, fmt.Errorf("%w: csv field index must be >= 0, got %d", ErrInvalidArgument, index)
	}
	if expected == "" {
		return nil, fmt.Errorf("%w: csv expected value must not be empty", ErrInvalidArgument)
	}
	return &CSVFieldFilter{index: index, expected: expected, ignoreCase: ignoreCase}, nil
}

// Match returns true if the CSV field at the configured index matches the expected value.
func (f *CSVFieldFilter) Match(line string) bool {
	r := csv.NewReader(strings.NewReader(line))
	fields, err := r.Read()
	if err != nil || f.index >= len(fields) {
		return false
	}
	val := fields[f.index]
	if f.ignoreCase {
		return strings.EqualFold(val, f.expected)
	}
	return val == f.expected
}

// Transform returns the line unchanged.
func (f *CSVFieldFilter) Transform(line string) string { return line }

// Index returns the zero-based field index.
func (f *CSVFieldFilter) Index() int { return f.index }

// Expected returns the expected field value.
func (f *CSVFieldFilter) Expected() string { return f.expected }

// IgnoreCase returns whether matching is case-insensitive.
func (f *CSVFieldFilter) IgnoreCase() bool { return f.ignoreCase }
