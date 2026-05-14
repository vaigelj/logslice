package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addSeverityFilter appends a SeverityFilter to filters if MinSeverity is set.
func addSeverityFilter(cfg *Config, filters *[]filter.Filter) error {
	if cfg.MinSeverity == "" {
		return nil
	}
	f, err := filter.NewSeverityFilter(cfg.MinSeverity)
	if err != nil {
		return fmt.Errorf("severity filter: %w", err)
	}
	*filters = append(*filters, f)
	return nil
}
