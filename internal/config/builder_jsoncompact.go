package config

import (
	"flag"

	"github.com/yourorg/logslice/internal/filter"
)

var (
	jsonCompact    bool
	jsonCompactPT  bool
)

func init() {
	flag.BoolVar(&jsonCompact, "json-compact", false,
		"re-encode each line as compact (minified) JSON")
	flag.BoolVar(&jsonCompactPT, "json-compact-passthrough", true,
		"pass non-JSON lines through unchanged when --json-compact is set")
}

// addJSONCompactFilter appends a JSONCompactFilter to the chain when --json-compact is set.
func addJSONCompactFilter(cfg *Config, filters []filter.Filter) ([]filter.Filter, error) {
	if !cfg.JSONCompact {
		return filters, nil
	}
	f, err := filter.NewJSONCompactFilter(cfg.JSONCompactPassthrough)
	if err != nil {
		return nil, err
	}
	return append(filters, f), nil
}
