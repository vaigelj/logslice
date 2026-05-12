package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addIPFilter appends an IPFilter to filters if cfg.CIDR is non-empty.
func addIPFilter(cfg *Config, filters *[]filter.Filter) error {
	if cfg.CIDR == "" {
		return nil
	}
	f, err := filter.NewIPFilter(cfg.CIDR)
	if err != nil {
		return fmt.Errorf("ip filter: %w", err)
	}
	*filters = append(*filters, f)
	return nil
}
