package config

import "flag"

func init() {
	// MinLength flag is registered here so it participates in flag.Parse.
	// The actual field is populated by Parse() via the flagSet defined in config.go.
	_ = flag.Int // ensure flag package is used; real registration happens in Parse via flagSet
}

// registerMinLengthFlag registers -min-length on the provided FlagSet and
// stores the parsed value into cfg.
func registerMinLengthFlag(fs *flag.FlagSet, cfg *Config) {
	fs.IntVar(&cfg.MinLength, "min-length", 0, "keep only lines with at least this many characters (0 = disabled)")
}
