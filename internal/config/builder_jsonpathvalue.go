package config

import (
	"flag"
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

var (
	jsonPathValuePath      string
	jsonPathValueExpected  string
	jsonPathValueCI        bool
)

func init() {
	registerJSONPathValueFlags(flag.CommandLine)
}

func registerJSONPathValueFlags(fs *flag.FlagSet) {
	fs.StringVar(&jsonPathValuePath, "json-path", "", "dot-notation JSON path to match (e.g. request.method)")
	fs.StringVar(&jsonPathValueExpected, "json-path-value", "", "expected value for the JSON path")
	fs.BoolVar(&jsonPathValueCI, "json-path-ci", false, "case-insensitive JSON path value matching")
}

func addJSONPathValueFilter(cfg *Config, chain *filter.Chain) error {
	if cfg.JSONPathValuePath == "" || cfg.JSONPathValueExpected == "" {
		return nil
	}
	f, err := filter.NewJSONPathValueFilter(cfg.JSONPathValuePath, cfg.JSONPathValueExpected, cfg.JSONPathValueCI)
	if err != nil {
		return fmt.Errorf("json-path filter: %w", err)
	}
	chain.Add(f)
	return nil
}
