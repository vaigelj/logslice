package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addNumberRangeFilter appends a NumberRangeFilter to the chain when the
// required flags are provided. NumRangeDelimiter, NumRangeField, NumRangeMin,
// and NumRangeMax must all be set (delimiter non-empty, max >= min).
func addNumberRangeFilter(cfg *Config, chain *filter.Chain) error {
	if cfg.NumRangeDelimiter == "" {
		return nil
	}
	if cfg.NumRangeMax < cfg.NumRangeMin {
		return fmt.Errorf("--num-range-max (%g) must be >= --num-range-min (%g)",
			cfg.NumRangeMax, cfg.NumRangeMin)
	}
	f, err := filter.NewNumberRangeFilter(
		cfg.NumRangeDelimiter,
		cfg.NumRangeField,
		cfg.NumRangeMin,
		cfg.NumRangeMax,
	)
	if err != nil {
		return fmt.Errorf("number range filter: %w", err)
	}
	chain.Add(f)
	return nil
}
