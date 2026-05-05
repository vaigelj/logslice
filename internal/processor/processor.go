// Package processor wires together reading, filtering, and writing
// into a single pipeline that processes log files line by line.
package processor

import (
	"context"
	"fmt"
	"io"

	"github.com/yourorg/logslice/internal/filter"
	"github.com/yourorg/logslice/internal/output"
	"github.com/yourorg/logslice/internal/reader"
)

// Config holds all parameters needed to run a processing job.
type Config struct {
	Input      io.Reader
	Output     io.Writer
	Filter     filter.Filter
	BufSize    int
	PrintStats bool
}

// Processor executes the log-slicing pipeline.
type Processor struct {
	cfg Config
}

// New creates a Processor from the given Config.
// An error is returned if required fields are missing.
func New(cfg Config) (*Processor, error) {
	if cfg.Input == nil {
		return nil, fmt.Errorf("processor: input reader must not be nil")
	}
	if cfg.Output == nil {
		return nil, fmt.Errorf("processor: output writer must not be nil")
	}
	return &Processor{cfg: cfg}, nil
}

// Run reads every line from the input, applies the filter chain, and writes
// matching lines to the output. It respects ctx cancellation.
func (p *Processor) Run(ctx context.Context) (*output.Stats, error) {
	lr := reader.NewLineReader(p.cfg.Input, p.cfg.BufSize)
	w := output.NewWriter(p.cfg.Output, p.cfg.BufSize)

	var totalRead int64

	for lr.Lines(ctx) {
		line := lr.Text()
		totalRead++

		if p.cfg.Filter != nil && !p.cfg.Filter.Match(line) {
			continue
		}

		if err := w.WriteLine(line); err != nil {
			return nil, fmt.Errorf("processor: write error: %w", err)
		}
	}

	if err := lr.Err(); err != nil {
		return nil, fmt.Errorf("processor: read error: %w", err)
	}

	if err := w.Flush(); err != nil {
		return nil, fmt.Errorf("processor: flush error: %w", err)
	}

	stats := &output.Stats{
		LinesRead:    totalRead,
		LinesWritten: w.Count(),
	}

	if p.cfg.PrintStats {
		stats.Print(p.cfg.Output)
	}

	return stats, nil
}
