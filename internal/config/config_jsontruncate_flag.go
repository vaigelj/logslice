package config

import "flag"

func init() {
	registerJSONTruncateFlags()
}

// registerJSONTruncateFlags registers CLI flags for the JSON truncate filter.
func registerJSONTruncateFlags() {
	flag.StringVar(
		&defaultConfig.JSONTruncateField,
		"json-truncate-field",
		"",
		"JSON field name whose string value should be truncated",
	)
	flag.IntVar(
		&defaultConfig.JSONTruncateMaxLen,
		"json-truncate-max",
		0,
		"maximum byte length for the JSON field value (0 = disabled)",
	)
	flag.StringVar(
		&defaultConfig.JSONTruncateSuffix,
		"json-truncate-suffix",
		"...",
		"suffix appended when a JSON field value is truncated",
	)
}
