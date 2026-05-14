package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addJSONPresentFilter appends a JSONPresentFilter to the chain when the
// --json-present or --json-absent flags are provided.
func addJSONPresentFilter(cfg *Config, chain *filter.Chain) error {
	field := cfg.JSONPresentField
	inverted := false

	if field == "" {
		field = cfg.JSONAbsentField
		inverted = true
	}

	if field == "" {
		return nil
	}

	f, err := filter.NewJSONPresentFilter(field, inverted)
	if err != nil {
		return fmt.Errorf("json-present filter: %w", err)
	}
	chain.Add(f)
	return nil
}

func init() {
	registerJSONPresentFlags()
}

func registerJSONPresentFlags() {
	registerFlag("json-present", "", "match lines where the JSON field is present (dot notation supported)")
	registerFlag("json-absent", "", "match lines where the JSON field is absent (dot notation supported)")
}
