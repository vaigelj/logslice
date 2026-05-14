package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addWordCountFilter appends a WordCountFilter to the chain when either
// --word-min or --word-max is non-zero.
func addWordCountFilter(cfg *Config, chain *[]filter.Filter) error {
	if cfg.WordMin == 0 && cfg.WordMax == 0 {
		return nil
	}
	f, err := filter.NewWordCountFilter(cfg.WordMin, cfg.WordMax)
	if err != nil {
		return fmt.Errorf("wordcount filter: %w", err)
	}
	*chain = append(*chain, f)
	return nil
}
