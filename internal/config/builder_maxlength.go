package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addMaxLengthFilter appends a MaxLengthFilter to the chain when cfg.MaxLength
// is greater than zero.
func addMaxLengthFilter(cfg *Config, filters []filter.Filter) ([]filter.Filter, error) {
	if cfg.MaxLength <= 0 {
		return filters, nil
	}
	f, err := filter.NewMaxLengthFilter(cfg.MaxLength)
	if err != nil {
		return nil, fmt.Errorf("max-length filter: %w", err)
	}
	return append(filters, f), nil
}
