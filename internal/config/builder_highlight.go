package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// ANSI colour presets available via --highlight-color flag.
var ansiColors = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\034[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"bold":    "\033[1m",
}

// addHighlightFilter appends a HighlightFilter to the chain when the config
// specifies a --highlight pattern. The colour defaults to yellow when the
// supplied name is not recognised.
func addHighlightFilter(cfg *Config, filters []filter.Filter) ([]filter.Filter, error) {
	if cfg.Highlight == "" {
		return filters, nil
	}

	colorCode, ok := ansiColors[cfg.HighlightColor]
	if !ok {
		// Fall back to yellow so output is still useful.
		colorCode = ansiColors["yellow"]
	}

	h, err := filter.NewHighlightFilter(cfg.Highlight, colorCode)
	if err != nil {
		return nil, fmt.Errorf("highlight filter: %w", err)
	}
	return append(filters, h), nil
}
