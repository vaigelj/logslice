package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
	"github.com/yourorg/logslice/internal/filter"
)

func TestBuildFilterChain_WithCountLimit(t *testing.T) {
	cfg := &config.Config{MaxMatchCount: 2}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
	matches := 0
	for _, line := range []string{"line1", "line2", "line3"} {
		if chain.Match(line) {
			matches++
		}
	}
	if matches != 2 {
		t.Errorf("expected 2 matches, got %d", matches)
	}
}

func TestBuildFilterChain_CountLimitZeroSkipped(t *testing.T) {
	cfg := &config.Config{MaxMatchCount: 0}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// With no filters, chain should match everything (or be nil — both acceptable)
	if chain != nil {
		for _, line := range []string{"a", "b", "c"} {
			if !chain.Match(line) {
				t.Errorf("expected line %q to match when no limit set", line)
			}
		}
	}
}

func TestBuildFilterChain_CountLimitAndRegex(t *testing.T) {
	cfg := &config.Config{
		Pattern:       "warn",
		MaxMatchCount: 1,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	results := []struct {
		line string
		want bool
	}{
		{"warn: disk full", true},
		{"warn: low memory", false},
		{"info: all good", false},
	}
	for _, tc := range results {
		got := chain.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestBuildFilterChain_NoCountLimit(t *testing.T) {
	cfg := &config.Config{}
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_ = filter.Filter(nil) // compile-time check
}
