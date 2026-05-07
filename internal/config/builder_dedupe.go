package config

import (
	"github.com/yourorg/logslice/internal/filter"
)

// addDedupeFilter appends a DedupeFilter to filters when deduplication is
// enabled in the config.
func addDedupeFilter(cfg *Config, filters *[]filter.Filter) {
	if !cfg.Dedupe {
		return
	}
	*filters = append(*filters, filter.NewDedupeFilter())
}
