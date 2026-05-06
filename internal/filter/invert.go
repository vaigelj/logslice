package filter

// InvertFilter wraps another Filter and negates its Match result.
// It is useful when you want to exclude lines that match a given pattern
// rather than include them.
type InvertFilter struct {
	inner Filter
}

// NewInvertFilter returns an InvertFilter that negates the result of inner.
// It returns an error if inner is nil.
func NewInvertFilter(inner Filter) (*InvertFilter, error) {
	if inner == nil {
		return nil, ErrNilFilter
	}
	return &InvertFilter{inner: inner}, nil
}

// Match returns true when the wrapped filter returns false, and vice-versa.
func (f *InvertFilter) Match(line []byte) bool {
	return !f.inner.Match(line)
}

// Inner returns the wrapped Filter.
func (f *InvertFilter) Inner() Filter {
	return f.inner
}
