package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addTruncateFilter appends a TruncateFilter to the chain when TruncateAt > 0.
// The filter is informational only at the chain level (Match always returns
// true); callers that want transformed output should use the Transform method
// directly on the filter returned via the chain's filters.
func addTruncateFilter(cfg *Config, chain *filter.Chain) error {
	if cfg.TruncateAt <= 0 {
		return nil
	}
	f, err := filter.NewTruncateFilter(cfg.TruncateAt, cfg.TruncateSuffix)
	if err != nil {
		return fmt.Errorf("truncate filter: %w", err)
	}
	chain.Add(f)
	return nil
}
