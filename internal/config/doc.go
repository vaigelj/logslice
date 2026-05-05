// Package config handles command-line flag parsing and validation for the
// logslice utility.
//
// Usage:
//
//	cfg, err := config.Parse(os.Args[1:], os.Stderr)
//	if err != nil {
//	    os.Exit(2)
//	}
//
// The returned Config struct is consumed by cmd/logslice/main.go to wire
// together the reader, filter chain, and output writer.
package config
