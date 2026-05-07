package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addContextLinesFilter appends a ContextLinesFilter to the chain when the
// config specifies a non-nil anchor filter and at least one context line.
func addContextLinesFilter(cfg *Config, chain *filter.Chain) error {
	if cfg.ContextAfter == 0 && cfg.ContextBefore == 0 {
		return nil
	}
	if cfg.Regex == "" {
		return fmt.Errorf("context-lines: --regex is required when using --context-before or --context-after")
	}

	anchor, err := filter.NewRegexFilter(cfg.Regex)
	if err != nil {
		return fmt.Errorf("context-lines anchor: %w", err)
	}

	cf, err := filter.NewContextLinesFilter(anchor, cfg.ContextBefore, cfg.ContextAfter)
	if err != nil {
		return fmt.Errorf("context-lines: %w", err)
	}

	chain.Add(cf)
	return nil
}
