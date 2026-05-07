package config

import (
	"strings"

	"github.com/yourorg/logslice/internal/filter"
)

// addContainsAnyFilter appends a ContainsAnyFilter to filters if ContainsAny terms are set.
// Terms are expected as a comma-separated string in cfg.ContainsAny.
func addContainsAnyFilter(cfg *Config, filters []filter.Filter) ([]filter.Filter, error) {
	if cfg.ContainsAny == "" {
		return filters, nil
	}
	terms := splitTerms(cfg.ContainsAny)
	if len(terms) == 0 {
		return filters, nil
	}
	f, err := filter.NewContainsAnyFilter(terms, cfg.CaseInsensitive)
	if err != nil {
		return nil, err
	}
	return append(filters, f), nil
}

// splitTerms splits a comma-separated string into trimmed, non-empty tokens.
func splitTerms(raw string) []string {
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
