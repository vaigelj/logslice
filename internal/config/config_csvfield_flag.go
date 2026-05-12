package config

import "flag"

func init() {
	registerCSVFieldFlags()
}

func registerCSVFieldFlags() {
	flag.IntVar(
		&defaultConfig.CSVFieldIndex,
		"csv-field-index",
		0,
		"zero-based CSV column index to match against (used with -csv-field-value)",
	)
	flag.StringVar(
		&defaultConfig.CSVFieldValue,
		"csv-field-value",
		"",
		"expected value for the CSV field at -csv-field-index",
	)
	flag.BoolVar(
		&defaultConfig.CSVFieldIgnoreCase,
		"csv-field-ignore-case",
		false,
		"perform case-insensitive comparison for CSV field match",
	)
}
