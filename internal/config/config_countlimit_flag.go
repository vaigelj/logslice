package config

import "flag"

func init() {
	registerCountLimitFlag()
}

// registerCountLimitFlag adds the -max-matches flag to the default flag set.
func registerCountLimitFlag() {
	if flag.Lookup("max-matches") != nil {
		return
	}
	flag.Int64Var(
		new(int64),
		"max-matches",
		0,
		"stop after emitting this many matching lines (0 = unlimited)",
	)
}
