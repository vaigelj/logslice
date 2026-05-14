package config

import "flag"

func init() {
	registerSeverityFlag(flag.CommandLine)
}

// registerSeverityFlag registers the --min-severity flag on the given FlagSet.
func registerSeverityFlag(fs *flag.FlagSet) {
	fs.String(
		"min-severity",
		"",
		"Minimum log severity level to include (trace|debug|info|warn|error|fatal)",
	)
}
