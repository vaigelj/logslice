// Command logslice is a fast log file slicer and filter utility.
// It supports filtering log lines by regular expression and/or time range.
//
// Usage:
//
//	logslice [flags] [file...]
//
// Flags:
//
//	-pattern string    regex pattern to match log lines
//	-start   string    start of time range (inclusive)
//	-end     string    end of time range (inclusive)
//	-layout  string    Go time layout for parsing timestamps (default: "2006-01-02T15:04:05")
//	-buf     int       buffer size in bytes for reader and writer (default: 65536)
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/user/logslice/internal/filter"
	"github.com/user/logslice/internal/output"
	"github.com/user/logslice/internal/processor"
	"github.com/user/logslice/internal/reader"
)

const (
	defaultLayout  = "2006-01-02T15:04:05"
	defaultBufSize = 65536
)

func main() {
	pattern := flag.String("pattern", "", "regex pattern to match log lines")
	start := flag.String("start", "", "start of time range (inclusive)")
	end := flag.String("end", "", "end of time range (inclusive)")
	layout := flag.String("layout", defaultLayout, "Go time layout for parsing timestamps")
	bufSize := flag.Int("buf", defaultBufSize, "buffer size in bytes for reader and writer")
	flag.Parse()

	if err := run(*pattern, *start, *end, *layout, *bufSize, flag.Args()); err != nil {
		log.Fatalf("logslice: %v", err)
	}
}

func run(pattern, start, end, layout string, bufSize int, args []string) error {
	// Build filter chain from provided flags.
	var filters []filter.Filter

	if pattern != "" {
		rf, err := filter.NewRegexFilter(pattern)
		if err != nil {
			return fmt.Errorf("invalid pattern: %w", err)
		}
		filters = append(filters, rf)
	}

	if start != "" || end != "" {
		trf, err := filter.NewTimeRangeFilter(layout, start, end)
		if err != nil {
			return fmt.Errorf("invalid time range: %w", err)
		}
		filters = append(filters, trf)
	}

	chain := filter.NewChain(filters...)

	// Determine input source: files provided as arguments or stdin.
	var input io.Reader
	switch len(args) {
	case 0:
		input = os.Stdin
	case 1:
		f, err := os.Open(args[0])
		if err != nil {
			return fmt.Errorf("open file: %w", err)
		}
		defer f.Close()
		input = f
	default:
		// Concatenate multiple files.
		readers := make([]io.Reader, 0, len(args))
		for _, path := range args {
			f, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("open file %q: %w", path, err)
			}
			defer f.Close() //nolint:gocritic // deferred inside loop intentionally for cleanup
			readers = append(readers, f)
		}
		input = io.MultiReader(readers...)
	}

	lr := reader.NewLineReader(input, reader.WithBufSize(bufSize))
	w := output.NewWriter(os.Stdout, output.WithBufSize(bufSize))

	proc := processor.New(lr, w, chain)
	if err := proc.Run(); err != nil {
		return fmt.Errorf("processing: %w", err)
	}

	stats := &output.Stats{
		Read:    lr.LinesRead(),
		Written: w.Count(),
	}
	stats.Print(os.Stderr)

	return nil
}
