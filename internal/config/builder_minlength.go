package config

import "github.com/yourorg/logslice/internal/filter"

// addMinLengthFilter appends a MinLengthFilter to the chain if cfg.MinLength > 0.
func addMinLengthFilter(cfg *Config, chain *filter.Chain) error {
	if cfg.MinLength <= 0 {
		return nil
	}
	f, err := filter.NewMinLengthFilter(cfg.MinLength)
	if err != nil {
		return err
	}
	chain.Add(f)
	return nil
}
