package filter

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// JSONRangeFilter matches lines where a numeric JSON field falls within [Min, Max].
type JSONRangeFilter struct {
	field string
	min   float64
	max   float64
}

// NewJSONRangeFilter creates a filter that matches lines where the numeric value
// of the given JSON field is within the inclusive range [min, max].
func NewJSONRangeFilter(field string, min, max float64) (*JSONRangeFilter, error) {
	if field == "" {
		return nil, fmt.Errorf("%w: field must not be empty", ErrInvalidArgument)
	}
	if min > max {
		return nil, fmt.Errorf("%w: min (%s) exceeds max (%s)",
			ErrInvalidArgument,
			strconv.FormatFloat(min, 'f', -1, 64),
			strconv.FormatFloat(max, 'f', -1, 64),
		)
	}
	return &JSONRangeFilter{field: field, min: min, max: max}, nil
}

// Match returns true if the JSON field value in line is a number within [min, max].
func (f *JSONRangeFilter) Match(line string) bool {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return false
	}
	v, ok := obj[f.field]
	if !ok {
		return false
	}
	var num float64
	switch val := v.(type) {
	case float64:
		num = val
	case string:
		parsed, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return false
		}
		num = parsed
	default:
		return false
	}
	return num >= f.min && num <= f.max
}

// Field returns the JSON field name being evaluated.
func (f *JSONRangeFilter) Field() string { return f.field }

// Min returns the lower bound of the range.
func (f *JSONRangeFilter) Min() float64 { return f.min }

// Max returns the upper bound of the range.
func (f *JSONRangeFilter) Max() float64 { return f.max }
