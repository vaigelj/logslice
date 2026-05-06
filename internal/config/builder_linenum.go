package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addLineNumFilter appends a LineNumFilter to the chain when either
// --line-start or --line-end flags are provided.
func addLineNumFilter(cfg *Config, chain *filter.Chain) error {
	if cfg.LineStart == 0 && cfg.LineEnd == 0 {
		return nil
	}
	f, err := filter.NewLineNumFilter(cfg.LineStart, cfg.LineEnd)
	if err != nil {
		return fmt.Errorf("line-number filter: %w", err)
	}
	chain.Add(f)
	return nil
}
