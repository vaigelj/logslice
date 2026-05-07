package filter

import (
	"fmt"
	"regexp"
	"strings"
)

// HighlightFilter is a Transformer that wraps matched substrings with ANSI
// colour escape codes so they stand out in terminal output.
type HighlightFilter struct {
	pattern *regexp.Regexp
	color   string
}

const ansiReset = "\033[0m"

// NewHighlightFilter returns a HighlightFilter that wraps every occurrence of
// pattern in the given ANSI colour code (e.g. "\033[31m" for red).
// An empty pattern or color returns ErrEmptyPattern / ErrEmptyValue.
func NewHighlightFilter(pattern, color string) (*HighlightFilter, error) {
	if pattern == "" {
		return nil, ErrEmptyPattern
	}
	if color == "" {
		return nil, fmt.Errorf("%w: color", ErrEmptyValue)
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidPattern, err)
	}
	return &HighlightFilter{pattern: re, color: color}, nil
}

// Match always returns true — the highlight filter never drops lines.
func (h *HighlightFilter) Match(_ string) bool { return true }

// Transform replaces every match of the pattern with a coloured version.
func (h *HighlightFilter) Transform(line string) string {
	return h.pattern.ReplaceAllStringFunc(line, func(m string) string {
		return h.color + m + ansiReset
	})
}

// Pattern returns the compiled pattern string.
func (h *HighlightFilter) Pattern() string { return h.pattern.String() }

// Color returns the ANSI escape code used for highlighting.
func (h *HighlightFilter) Color() string { return h.color }

// String returns a human-readable description.
func (h *HighlightFilter) String() string {
	return fmt.Sprintf("highlight(%s)", strings.TrimSpace(h.pattern.String()))
}
