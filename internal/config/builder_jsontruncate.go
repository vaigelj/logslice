package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addJSONTruncateFilter appends a JSONTruncateFilter to the chain if configured.
func addJSONTruncateFilter(cfg *Config, chain *filter.Chain) error {
	if cfg.JSONTruncateField == "" || cfg.JSONTruncateMaxLen <= 0 {
		return nil
	}
	f, err := filter.NewJSONTruncateFilter(
		cfg.JSONTruncateField,
		cfg.JSONTruncateMaxLen,
		cfg.JSONTruncateSuffix,
	)
	if err != nil {
		return fmt.Errorf("jsontruncate filter: %w", err)
	}
	chain.Add(f)
	return nil
}
