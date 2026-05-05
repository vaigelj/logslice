package filter

// Filter is the common interface for all log line filters.
type Filter interface {
	Match(line string) bool
}

// Chain combines multiple filters with AND semantics:
// a line must match all filters to be included.
type Chain struct {
	filters []Filter
}

// NewChain creates a Chain from the provided filters.
// Nil filters are silently ignored.
func NewChain(filters ...Filter) *Chain {
	active := make([]Filter, 0, len(filters))
	for _, f := range filters {
		if f != nil {
			active = append(active, f)
		}
	}
	return &Chain{filters: active}
}

// Match returns true only if every filter in the chain matches the line.
// An empty chain matches all lines.
func (c *Chain) Match(line string) bool {
	for _, f := range c.filters {
		if !f.Match(line) {
			return false
		}
	}
	return true
}

// Len returns the number of active filters in the chain.
func (c *Chain) Len() int {
	return len(c.filters)
}

// Add appends additional filters to the chain.
func (c *Chain) Add(filters ...Filter) {
	for _, f := range filters {
		if f != nil {
			c.filters = append(c.filters, f)
		}
	}
}
