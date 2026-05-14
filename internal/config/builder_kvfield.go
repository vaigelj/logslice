package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addKVFieldFilter appends a KVFieldFilter to the chain if configured.
func addKVFieldFilter(cfg *Config, chain *filter.Chain) error {
	if cfg.KVKey == "" {
		return nil
	}
	f, err := filter.NewKVFieldFilter(
		cfg.KVKey,
		cfg.KVSep,
		cfg.KVValue,
		cfg.KVCaseInsensitive,
	)
	if err != nil {
		return fmt.Errorf("kv-field filter: %w", err)
	}
	chain.Add(f)
	return nil
}
