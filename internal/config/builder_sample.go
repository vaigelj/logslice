package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addSampleFilter appends a SampleFilter to filters if cfg.SampleN > 1.
// SampleN=1 is a no-op (every line passes) so we skip it.
func addSampleFilter(cfg *Config, filters *[]filter.Filter) error {
	if cfg.SampleN <= 1 {
		return nil
	}
	f, err := filter.NewSampleFilter(cfg.SampleN)
	if err != nil {
		return fmt.Errorf("sample filter: %w", err)
	}
	*filters = append(*filters, f)
	return nil
}
