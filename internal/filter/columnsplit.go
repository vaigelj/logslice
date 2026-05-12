package filter

import (
	"fmt"
	"strings"
)

// ColumnSplitFilter splits each line by a delimiter and emits only the
// selected column (zero-based index). Lines that do not contain enough
// columns are dropped (Match returns false).
type ColumnSplitFilter struct {
	delimiter string
	column    int
}

// NewColumnSplitFilter creates a ColumnSplitFilter.
// delimiter must be non-empty and column must be >= 0.
func NewColumnSplitFilter(delimiter string, column int) (*ColumnSplitFilter, error) {
	if delimiter == "" {
		return nil, fmt.Errorf("%w: delimiter must not be empty", ErrInvalidFilter)
	}
	if column < 0 {
		return nil, fmt.Errorf("%w: column index must be >= 0", ErrInvalidFilter)
	}
	return &ColumnSplitFilter{delimiter: delimiter, column: column}, nil
}

// Match returns true when the line has enough columns. The line is also
// transformed so that downstream filters and the writer see only the
// selected column value.
func (f *ColumnSplitFilter) Match(line *string) bool {
	parts := strings.Split(*line, f.delimiter)
	if f.column >= len(parts) {
		return false
	}
	*line = parts[f.column]
	return true
}

// Delimiter returns the configured delimiter.
func (f *ColumnSplitFilter) Delimiter() string { return f.delimiter }

// Column returns the configured column index.
func (f *ColumnSplitFilter) Column() int { return f.column }
