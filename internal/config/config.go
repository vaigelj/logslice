// Package config provides CLI argument parsing and validation for logslice.
package config

import (
	"errors"
	"flag"
	"io"
	"time"
)

// Config holds all runtime configuration parsed from command-line flags.
type Config struct {
	InputFile   string
	OutputFile  string
	RegexPattern string
	TimeLayout  string
	TimeStart   string
	TimeEnd     string
	BufSize     int
	Quiet       bool
}

// Parse reads flags from args (excluding the program name) and returns a Config.
// w is used for flag usage output.
func Parse(args []string, w io.Writer) (*Config, error) {
	fs := flag.NewFlagSet("logslice", flag.ContinueOnError)
	fs.SetOutput(w)

	cfg := &Config{}

	fs.StringVar(&cfg.InputFile, "input", "", "input log file (default: stdin)")
	fs.StringVar(&cfg.OutputFile, "output", "", "output file (default: stdout)")
	fs.StringVar(&cfg.RegexPattern, "regex", "", "regex pattern to filter lines")
	fs.StringVar(&cfg.TimeLayout, "time-layout", time.RFC3339, "Go time layout for parsing timestamps")
	fs.StringVar(&cfg.TimeStart, "time-start", "", "start of time range (inclusive)")
	fs.StringVar(&cfg.TimeEnd, "time-end", "", "end of time range (inclusive)")
	fs.IntVar(&cfg.BufSize, "buf-size", 0, "buffer size in bytes (0 = default 64 KiB)")
	fs.BoolVar(&cfg.Quiet, "quiet", false, "suppress stats output")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	if (c.TimeStart != "" || c.TimeEnd != "") && c.TimeLayout == "" {
		return errors.New("--time-layout is required when using --time-start or --time-end")
	}
	if c.BufSize < 0 {
		return errors.New("--buf-size must be non-negative")
	}
	return nil
}
