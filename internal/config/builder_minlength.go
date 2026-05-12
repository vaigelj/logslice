package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addMinLengthFilter appends a MinLengthFilter to the chain if cfg.MinLength > 0.
// Returns an error if cfg is nil or if the filter cannot be constructed.
func addMinLengthFilter(cfg *Config, chain *filter.Chain) error {
	if cfg == nil {
		return fmt.Errorf("addMinLengthFilter: config must not be nil")
	}
	if chain == nil {
		return fmt.Errorf("addMinLengthFilter: filter chain must not be nil")
	}
	if cfg.MinLength <= 0 {
		return nil
	}
	f, err := filter.NewMinLengthFilter(cfg.MinLength)
	if err != nil {
		return fmt.Errorf("addMinLengthFilter: %w", err)
	}
	chain.Add(f)
	return nil
}
