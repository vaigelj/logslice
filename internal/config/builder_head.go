package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addHeadFilter appends a HeadFilter to the chain if cfg.HeadLines > 0.
func addHeadFilter(cfg *Config, filters []filter.Filter) ([]filter.Filter, error) {
	if cfg.HeadLines <= 0 {
		return filters, nil
	}
	f, err := filter.NewHeadFilter(cfg.HeadLines)
	if err != nil {
		return nil, fmt.Errorf("head filter: %w", err)
	}
	return append(filters, f), nil
}
