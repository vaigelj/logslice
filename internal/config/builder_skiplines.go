package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addSkipLinesFilter appends a SkipLinesFilter to the chain if cfg.SkipLines > 0.
func addSkipLinesFilter(cfg *Config, chain []filter.Filter) ([]filter.Filter, error) {
	if cfg.SkipLines <= 0 {
		return chain, nil
	}
	f, err := filter.NewSkipLinesFilter(cfg.SkipLines)
	if err != nil {
		return nil, fmt.Errorf("skip-lines filter: %w", err)
	}
	return append(chain, f), nil
}
