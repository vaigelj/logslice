package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// BuildFilterChain constructs a filter.Chain from the parsed Config.
// Filters are added in order: prefix (if set), regex (if set),
// time-range (if set), and each is optionally inverted.
func BuildFilterChain(cfg *Config) (*filter.Chain, error) {
	var filters []filter.Filter

	if cfg.Prefix != "" {
		pf, err := filter.NewPrefixFilter(cfg.Prefix, cfg.PrefixCaseFold)
		if err != nil {
			return nil, fmt.Errorf("prefix filter: %w", err)
		}
		var f filter.Filter = pf
		if cfg.InvertPrefix {
			inv, err := filter.NewInvertFilter(f)
			if err != nil {
				return nil, fmt.Errorf("invert prefix filter: %w", err)
			}
			f = inv
		}
		filters = append(filters, f)
	}

	if cfg.Pattern != "" {
		rf, err := filter.NewRegexFilter(cfg.Pattern)
		if err != nil {
			return nil, fmt.Errorf("regex filter: %w", err)
		}
		var f filter.Filter = rf
		if cfg.Invert {
			inv, err := filter.NewInvertFilter(f)
			if err != nil {
				return nil, fmt.Errorf("invert regex filter: %w", err)
			}
			f = inv
		}
		filters = append(filters, f)
	}

	if cfg.TimeLayout != "" {
		tf, err := filter.NewTimeRangeFilter(cfg.TimeLayout, cfg.TimeStart, cfg.TimeEnd)
		if err != nil {
			return nil, fmt.Errorf("time-range filter: %w", err)
		}
		filters = append(filters, tf)
	}

	return filter.NewChain(filters...), nil
}
