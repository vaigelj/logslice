package filter

import (
	"fmt"
	"strconv"
	"strings"
)

// NumberRangeFilter matches lines that contain a numeric field within [Min, Max].
type NumberRangeFilter struct {
	delimiter string
	fieldIdx  int
	min       float64
	max       float64
}

// NewNumberRangeFilter creates a filter that parses a numeric field from each line
// (split by delimiter at fieldIdx) and matches if the value is within [min, max].
func NewNumberRangeFilter(delimiter string, fieldIdx int, min, max float64) (*NumberRangeFilter, error) {
	if delimiter == "" {
		return nil, fmt.Errorf("%w: delimiter must not be empty", ErrInvalidArgument)
	}
	if fieldIdx < 0 {
		return nil, fmt.Errorf("%w: field index must be non-negative", ErrInvalidArgument)
	}
	if min > max {
		return nil, fmt.Errorf("%w: min (%g) exceeds max (%g)", ErrInvalidArgument, min, max)
	}
	return &NumberRangeFilter{
		delimiter: delimiter,
		fieldIdx:  fieldIdx,
		min:       min,
		max:       max,
	}, nil
}

// Match returns true if the numeric field at fieldIdx is within [min, max].
func (f *NumberRangeFilter) Match(line string) bool {
	parts := strings.Split(line, f.delimiter)
	if f.fieldIdx >= len(parts) {
		return false
	}
	v, err := strconv.ParseFloat(strings.TrimSpace(parts[f.fieldIdx]), 64)
	if err != nil {
		return false
	}
	return v >= f.min && v <= f.max
}

// Transform returns the line unchanged.
func (f *NumberRangeFilter) Transform(line string) string { return line }

// Delimiter returns the configured field delimiter.
func (f *NumberRangeFilter) Delimiter() string { return f.delimiter }

// FieldIndex returns the zero-based field index.
func (f *NumberRangeFilter) FieldIndex() int { return f.fieldIdx }

// Min returns the lower bound of the numeric range.
func (f *NumberRangeFilter) Min() float64 { return f.min }

// Max returns the upper bound of the numeric range.
func (f *NumberRangeFilter) Max() float64 { return f.max }
