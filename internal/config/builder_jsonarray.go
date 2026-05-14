package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addJSONArrayFilter appends a JSONArrayFilter to filters if cfg.JSONArrayField
// and cfg.JSONArrayValue are both non-empty.
func addJSONArrayFilter(cfg *Config, filters *[]filter.Filter) error {
	if cfg.JSONArrayField == "" || cfg.JSONArrayValue == "" {
		return nil
	}
	f, err := filter.NewJSONArrayFilter(cfg.JSONArrayField, cfg.JSONArrayValue, cfg.JSONArrayIgnoreCase)
	if err != nil {
		return fmt.Errorf("json-array filter: %w", err)
	}
	*filters = append(*filters, f)
	return nil
}

func init() {
	registerJSONArrayFlags()
}

func registerJSONArrayFlags() {
	flags := defaultFlagSet()
	flags.StringVar(&globalConfig.JSONArrayField, "json-array-field", "", "JSON array field name to search within")
	flags.StringVar(&globalConfig.JSONArrayValue, "json-array-value", "", "value to look for inside the JSON array field")
	flags.BoolVar(&globalConfig.JSONArrayIgnoreCase, "json-array-ignore-case", false, "case-insensitive matching for json-array-value")
}
