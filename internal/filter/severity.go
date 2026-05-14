package filter

import (
	"fmt"
	"strings"
)

// severityLevels maps level name to numeric priority (higher = more severe).
var severityLevels = map[string]int{
	"trace": 0,
	"debug": 1,
	"info":  2,
	"warn":  3,
	"error": 4,
	"fatal": 5,
}

// SeverityFilter matches log lines whose severity level meets a minimum threshold.
// It scans each line for a known level keyword (case-insensitive).
type SeverityFilter struct {
	minLevel string
	minRank  int
}

// NewSeverityFilter creates a SeverityFilter that passes lines at or above minLevel.
func NewSeverityFilter(minLevel string) (*SeverityFilter, error) {
	if minLevel == "" {
		return nil, fmt.Errorf("%w: severity level must not be empty", ErrInvalidArgument)
	}
	rank, ok := severityLevels[strings.ToLower(minLevel)]
	if !ok {
		return nil, fmt.Errorf("%w: unknown severity level %q", ErrInvalidArgument, minLevel)
	}
	return &SeverityFilter{minLevel: strings.ToLower(minLevel), minRank: rank}, nil
}

// Match returns true if the line contains a severity keyword >= the minimum level.
func (f *SeverityFilter) Match(line string) bool {
	lower := strings.ToLower(line)
	best := -1
	for lvl, rank := range severityLevels {
		if strings.Contains(lower, lvl) && rank > best {
			best = rank
		}
	}
	if best < 0 {
		return false
	}
	return best >= f.minRank
}

// MinLevel returns the configured minimum severity level string.
func (f *SeverityFilter) MinLevel() string { return f.minLevel }
