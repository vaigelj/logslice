package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
)

func TestBuildFilterChain_WithSample(t *testing.T) {
	cfg := &config.Config{
		SampleN: 3,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
	// 9 lines, every 3rd should match => 3 matches
	matches := 0
	for i := 0; i < 9; i++ {
		if chain.Match("log line") {
			matches++
		}
	}
	if matches != 3 {
		t.Errorf("expected 3 matches, got %d", matches)
	}
}

func TestBuildFilterChain_SampleOneSkipped(t *testing.T) {
	// SampleN=1 should be skipped (no-op)
	cfg := &config.Config{
		SampleN: 1,
	}
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBuildFilterChain_SampleZeroSkipped(t *testing.T) {
	cfg := &config.Config{
		SampleN: 0,
	}
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBuildFilterChain_SampleAndRegex(t *testing.T) {
	cfg := &config.Config{
		Pattern: "error",
		SampleN: 2,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := []string{
		"error alpha",
		"error beta",
		"info gamma",
		"error delta",
	}
	matches := 0
	for _, l := range lines {
		if chain.Match(l) {
			matches++
		}
	}
	// regex passes: alpha, beta, delta (3 lines)
	// sample n=2: beta, (delta skipped by counter reset? no — counter: alpha=1, beta=2->pass, delta=1)
	// => 1 match (beta)
	if matches != 1 {
		t.Errorf("expected 1 match, got %d", matches)
	}
}
