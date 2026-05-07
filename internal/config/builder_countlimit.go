package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addCountLimitFilter appends a CountLimitFilter to the chain if cfg.MaxMatchCount > 0.
func addCountLimitFilter(cfg *Config, filters *[]filter.Filter) error {
	if cfg.MaxMatchCount <= 0 {
		return nil
	}
	f, err := filter.NewCountLimitFilter(cfg.MaxMatchCount)
	if err != nil {
		return fmt.Errorf("count limit filter: %w", err)
	}
	*filters = append(*filters, f)
	return nil
}
