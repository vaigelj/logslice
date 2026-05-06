package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addSuffixFilter appends a SuffixFilter to filters if cfg.Suffix is set.
// It returns an error if the filter cannot be constructed.
func addSuffixFilter(cfg *Config, filters *[]filter.Filter) error {
	if cfg.Suffix == "" {
		return nil
	}
	f, err := filter.NewSuffixFilter(cfg.Suffix, cfg.CaseInsensitive)
	if err != nil {
		return fmt.Errorf("suffix filter: %w", err)
	}
	*filters = append(*filters, f)
	return nil
}
