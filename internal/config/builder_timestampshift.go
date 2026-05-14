package config

import (
	"fmt"
	"time"

	"github.com/yourorg/logslice/internal/filter"
)

// addTimestampShiftFilter appends a TimestampShiftFilter to the chain when
// both a layout and a non-zero shift duration are configured.
func addTimestampShiftFilter(cfg *Config, chain *filter.Chain) error {
	if cfg.TimestampShiftLayout == "" || cfg.TimestampShiftDuration == 0 {
		return nil
	}
	f, err := filter.NewTimestampShiftFilter(cfg.TimestampShiftLayout, cfg.TimestampShiftDuration)
	if err != nil {
		return fmt.Errorf("timestamp-shift: %w", err)
	}
	chain.Add(f)
	return nil
}

func init() {
	registerTimestampShiftFlags()
}

func registerTimestampShiftFlags() {
	// Flags are registered in config_timestampshift_flag.go via init().
	_ = time.Hour // ensure time import is used
}
