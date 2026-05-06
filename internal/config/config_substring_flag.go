package config

// This file registers the --substring / -s flag for the Config parser.
// It is kept separate to keep config.go focused on core flag definitions.

import "flag"

func init() {
	// registerSubstringFlags is called by Parse via the flagSet setup.
	// The actual registration happens inside Parse using the flagSet passed there.
	// This file documents the intent; the flag is wired in config.go Parse.
	_ = flag.CommandLine // ensure flag package is imported for documentation.
}

// substringFlagName is the canonical CLI flag name for substring filtering.
const substringFlagName = "substring"

// substringFlagShort is the short alias for the substring flag.
const substringFlagShort = "s"

// substringFlagUsage describes the substring flag for help output.
const substringFlagUsage = "filter lines containing this substring (use -i for case-insensitive)"
