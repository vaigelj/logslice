package processor_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/filter"
	"github.com/yourorg/logslice/internal/processor"
)

func TestNew_NilInput(t *testing.T) {
	_, err := processor.New(processor.Config{Output: &bytes.Buffer{}})
	if err == nil {
		t.Fatal("expected error for nil input, got nil")
	}
}

func TestNew_NilOutput(t *testing.T) {
	_, err := processor.New(processor.Config{Input: strings.NewReader("")})
	if err == nil {
		t.Fatal("expected error for nil output, got nil")
	}
}

func TestRun_NoFilter(t *testing.T) {
	input := "line one\nline two\nline three\n"
	out := &bytes.Buffer{}

	p, err := processor.New(processor.Config{
		Input:  strings.NewReader(input),
		Output: out,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	stats, err := p.Run(context.Background())
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}

	if stats.LinesRead != 3 {
		t.Errorf("LinesRead: want 3, got %d", stats.LinesRead)
	}
	if stats.LinesWritten != 3 {
		t.Errorf("LinesWritten: want 3, got %d", stats.LinesWritten)
	}
}

func TestRun_WithRegexFilter(t *testing.T) {
	input := "ERROR something failed\nINFO all good\nERROR another failure\n"
	out := &bytes.Buffer{}

	f, err := filter.NewRegexFilter("ERROR")
	if err != nil {
		t.Fatalf("filter error: %v", err)
	}

	p, err := processor.New(processor.Config{
		Input:  strings.NewReader(input),
		Output: out,
		Filter: f,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	stats, err := p.Run(context.Background())
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}

	if stats.LinesRead != 3 {
		t.Errorf("LinesRead: want 3, got %d", stats.LinesRead)
	}
	if stats.LinesWritten != 2 {
		t.Errorf("LinesWritten: want 2, got %d", stats.LinesWritten)
	}
	if !strings.Contains(out.String(), "ERROR") {
		t.Errorf("output should contain ERROR lines, got: %q", out.String())
	}
}

func TestRun_EmptyInput(t *testing.T) {
	out := &bytes.Buffer{}
	p, _ := processor.New(processor.Config{
		Input:  strings.NewReader(""),
		Output: out,
	})

	stats, err := p.Run(context.Background())
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}
	if stats.LinesRead != 0 || stats.LinesWritten != 0 {
		t.Errorf("expected zero stats, got read=%d written=%d", stats.LinesRead, stats.LinesWritten)
	}
}
