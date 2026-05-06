package filter

import (
	"fmt"
	"strings"
)

// FieldMatchFilter matches log lines where a specific delimited field equals a value.
type FieldMatchFilter struct {
	delimiter string
	fieldIndex int
	value      string
	caseInsensitive bool
}

// NewFieldMatchFilter creates a filter that checks if the field at fieldIndex
// (0-based) in a delimiter-split line equals value.
func NewFieldMatchFilter(delimiter string, fieldIndex int, value string, caseInsensitive bool) (*FieldMatchFilter, error) {
	if delimiter == "" {
		return nil, fmt.Errorf("%w: delimiter must not be empty", ErrInvalidArgument)
	}
	if fieldIndex < 0 {
		return nil, fmt.Errorf("%w: fieldIndex must be non-negative", ErrInvalidArgument)
	}
	if value == "" {
		return nil, fmt.Errorf("%w: value must not be empty", ErrInvalidArgument)
	}
	return &FieldMatchFilter{
		delimiter:       delimiter,
		fieldIndex:      fieldIndex,
		value:           value,
		caseInsensitive: caseInsensitive,
	}, nil
}

// Match returns true if the field at the configured index equals the target value.
func (f *FieldMatchFilter) Match(line string) bool {
	parts := strings.Split(line, f.delimiter)
	if f.fieldIndex >= len(parts) {
		return false
	}
	field := parts[f.fieldIndex]
	if f.caseInsensitive {
		return strings.EqualFold(field, f.value)
	}
	return field == f.value
}

// Delimiter returns the configured delimiter.
func (f *FieldMatchFilter) Delimiter() string { return f.delimiter }

// FieldIndex returns the configured field index.
func (f *FieldMatchFilter) FieldIndex() int { return f.fieldIndex }

// Value returns the target match value.
func (f *FieldMatchFilter) Value() string { return f.value }

// CaseInsensitive returns whether matching is case-insensitive.
func (f *FieldMatchFilter) CaseInsensitive() bool { return f.caseInsensitive }
