package config

import (
	"fmt"

	"github.com/user/logslice/internal/filter"
)

// BuildFilterChain constructs a filter.Chain from the provided Config.
// Filters are only added when the corresponding flag was supplied.
// Returns an error if any filter cannot be compiled (e.g. bad regex).
func BuildFilterChain(cfg *Config) (*filter.Chain, error) {
	var filters []filter.Filter

	if cfg.RegexPattern != "" {
		rf, err := filter.NewRegexFilter(cfg.RegexPattern)
		if err != nil {
			return nil, fmt.Errorf("regex filter: %w", err)
		}
		filters = append(filters, rf)
	}

	if cfg.TimeStart != "" || cfg.TimeEnd != "" {
		tf, err := filter.NewTimeRangeFilter(cfg.TimeLayout, cfg.TimeStart, cfg.TimeEnd)
		if err != nil {
			return nil, fmt.Errorf("time-range filter: %w", err)
		}
		filters = append(filters, tf)
	}

	return filter.NewChain(filters...), nil
}
