package config

import "flag"

func init() {
	registerKVFieldFlags()
}

func registerKVFieldFlags() {
	flag.StringVar(&globalFlags.KVKey, "kv-key", "",
		"filter lines by key=value pair: key name (requires --kv-value)")
	flag.StringVar(&globalFlags.KVSep, "kv-sep", "=",
		"separator between key and value for --kv-key (default '=')")
	flag.StringVar(&globalFlags.KVValue, "kv-value", "",
		"expected value for --kv-key")
	flag.BoolVar(&globalFlags.KVCaseInsensitive, "kv-ignore-case", false,
		"make --kv-key matching case-insensitive")
}
