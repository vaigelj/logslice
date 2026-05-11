package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addTemplateFilter appends a TemplateFilter to the chain if both
// --template-pattern and --template-text flags are provided.
func addTemplateFilter(cfg *Config, chain *filter.Chain) error {
	if cfg.TemplatePattern == "" || cfg.TemplateText == "" {
		return nil
	}
	f, err := filter.NewTemplateFilter(cfg.TemplatePattern, cfg.TemplateText)
	if err != nil {
		return fmt.Errorf("template filter: %w", err)
	}
	chain.Add(f)
	return nil
}
