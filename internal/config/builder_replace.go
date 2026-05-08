package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addReplaceFilter appends a ReplaceFilter to the chain when the --replace
// flag is set. The --replace-with flag provides the replacement string
// (defaults to empty string, effectively deleting matches).
func addReplaceFilter(cfg *Config, chain *filter.Chain) error {
	if cfg.ReplacePattern == "" {
		return nil
	}

	f, err := filter.NewReplaceFilter(
		cfg.ReplacePattern,
		cfg.ReplaceWith,
		cfg.CaseInsensitive,
		cfg.ReplaceRegex,
	)
	if err != nil {
		return fmt.Errorf("replace filter: %w", err)
	}

	chain.Add(f)
	return nil
}
