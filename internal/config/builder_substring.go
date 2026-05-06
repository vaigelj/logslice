package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addSubstringFilter appends a SubstringFilter to filters if cfg.Substring is set.
func addSubstringFilter(cfg *Config, filters *[]filter.Filter) error {
	if cfg.Substring == "" {
		return nil
	}
	f, err := filter.NewSubstringFilter(cfg.Substring, cfg.CaseInsensitive)
	if err != nil {
		return fmt.Errorf("substring filter: %w", err)
	}
	*filters = append(*filters, f)
	return nil
}
