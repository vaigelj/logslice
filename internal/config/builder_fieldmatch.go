package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addFieldMatchFilter appends a FieldMatchFilter to filters if the config
// specifies a non-empty FieldValue. Both FieldDelimiter and FieldIndex must
// also be set meaningfully (delimiter non-empty, index >= 0).
func addFieldMatchFilter(cfg *Config, filters []filter.Filter) ([]filter.Filter, error) {
	if cfg.FieldValue == "" {
		return filters, nil
	}
	delimiter := cfg.FieldDelimiter
	if delimiter == "" {
		delimiter = " "
	}
	f, err := filter.NewFieldMatchFilter(delimiter, cfg.FieldIndex, cfg.FieldValue, cfg.CaseInsensitive)
	if err != nil {
		return nil, fmt.Errorf("field-match filter: %w", err)
	}
	return append(filters, f), nil
}
