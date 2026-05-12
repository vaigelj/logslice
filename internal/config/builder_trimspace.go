package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addTrimSpaceFilter appends a TrimSpaceFilter to filters when trim flags are set.
func addTrimSpaceFilter(cfg *Config, filters *[]filter.Filter) error {
	if !cfg.TrimLeft && !cfg.TrimRight {
		return nil
	}
	f, err := filter.NewTrimSpaceFilter(cfg.TrimLeft, cfg.TrimRight)
	if err != nil {
		return fmt.Errorf("trimspace filter: %w", err)
	}
	*filters = append(*filters, f)
	return nil
}
