package config

import (
	"io"
	"testing"
	"time"
)

func TestParse_Defaults(t *testing.T) {
	cfg, err := Parse([]string{}, io.Discard)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.InputFile != "" {
		t.Errorf("expected empty InputFile, got %q", cfg.InputFile)
	}
	if cfg.TimeLayout != time.RFC3339 {
		t.Errorf("expected RFC3339 layout, got %q", cfg.TimeLayout)
	}
	if cfg.BufSize != 0 {
		t.Errorf("expected BufSize 0, got %d", cfg.BufSize)
	}
	if cfg.Quiet {
		t.Error("expected Quiet=false by default")
	}
}

func TestParse_AllFlags(t *testing.T) {
	args := []string{
		"-input", "app.log",
		"-output", "out.log",
		"-regex", `ERROR`,
		"-time-layout", "2006-01-02",
		"-time-start", "2024-01-01",
		"-time-end", "2024-12-31",
		"-buf-size", "8192",
		"-quiet",
	}
	cfg, err := Parse(args, io.Discard)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.InputFile != "app.log" {
		t.Errorf("InputFile: got %q", cfg.InputFile)
	}
	if cfg.RegexPattern != "ERROR" {
		t.Errorf("RegexPattern: got %q", cfg.RegexPattern)
	}
	if cfg.BufSize != 8192 {
		t.Errorf("BufSize: got %d", cfg.BufSize)
	}
	if !cfg.Quiet {
		t.Error("expected Quiet=true")
	}
}

func TestParse_InvalidFlag(t *testing.T) {
	_, err := Parse([]string{"-unknown-flag"}, io.Discard)
	if err == nil {
		t.Fatal("expected error for unknown flag")
	}
}

func TestParse_NegativeBufSize(t *testing.T) {
	_, err := Parse([]string{"-buf-size", "-1"}, io.Discard)
	if err == nil {
		t.Fatal("expected error for negative buf-size")
	}
}

func TestParse_TimeRangeWithoutLayout(t *testing.T) {
	// layout defaults to RFC3339, so providing a range should be fine
	_, err := Parse([]string{"-time-start", "2024-01-01T00:00:00Z"}, io.Discard)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
