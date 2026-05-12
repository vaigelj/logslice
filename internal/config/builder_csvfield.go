package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addCSVFieldFilter appends a CSVFieldFilter to the chain if the expected value is set.
func addCSVFieldFilter(chain []filter.Filter, cfg *Config) ([]filter.Filter, error) {
	if cfg.CSVFieldValue == "" {
		return chain, nil
	}
	f, err := filter.NewCSVFieldFilter(cfg.CSVFieldIndex, cfg.CSVFieldValue, cfg.CSVFieldIgnoreCase)
	if err != nil {
		return nil, fmt.Errorf("csv-field filter: %w", err)
	}
	return append(chain, f), nil
}
