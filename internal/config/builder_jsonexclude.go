package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

func init() {
	registerJSONExcludeFlags()
}

func registerJSONExcludeFlags() {
	flags.StringVar(&cfg.JSONExcludeField, "json-exclude-field", "", "JSON field name to match for exclusion")
	flags.StringVar(&cfg.JSONExcludeValue, "json-exclude-value", "", "JSON field value that causes line exclusion")
	flags.BoolVar(&cfg.JSONExcludeCaseInsensitive, "json-exclude-ci", false, "case-insensitive JSON exclude matching")
}

func addJSONExcludeFilter(chain []filter.Filter, cfg *Config) ([]filter.Filter, error) {
	if cfg.JSONExcludeField == "" || cfg.JSONExcludeValue == "" {
		return chain, nil
	}
	f, err := filter.NewJSONExcludeFilter(cfg.JSONExcludeField, cfg.JSONExcludeValue, cfg.JSONExcludeCaseInsensitive)
	if err != nil {
		return nil, fmt.Errorf("json-exclude: %w", err)
	}
	return append(chain, f), nil
}
