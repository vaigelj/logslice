package filter

// NotEmptyFilter rejects lines that are empty or contain only whitespace.
type NotEmptyFilter struct {
	trimWhitespace bool
}

// NewNotEmptyFilter creates a filter that rejects empty lines.
// If trimWhitespace is true, lines consisting solely of whitespace are also
// rejected.
func NewNotEmptyFilter(trimWhitespace bool) *NotEmptyFilter {
	return &NotEmptyFilter{trimWhitespace: trimWhitespace}
}

// Match returns true when the line is non-empty.
func (f *NotEmptyFilter) Match(line string) bool {
	if f.trimWhitespace {
		for _, r := range line {
			if r != ' ' && r != '\t' && r != '\r' && r != '\n' {
				return true
			}
		}
		return false
	}
	return len(line) > 0
}

// TrimWhitespace returns whether the filter treats whitespace-only lines as
// empty.
func (f *NotEmptyFilter) TrimWhitespace() bool {
	return f.trimWhitespace
}
