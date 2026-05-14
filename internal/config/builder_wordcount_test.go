package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
)

func TestBuildFilterChain_WithWordCount(t *testing.T) {
	cfg := &config.Config{WordMin: 2, WordMax: 5}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
}

func TestBuildFilterChain_WordCountZeroSkipped(t *testing.T) {
	cfg := &config.Config{WordMin: 0, WordMax: 0}
	// Should build successfully with no word-count filter added
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBuildFilterChain_WordCountInvalidRange(t *testing.T) {
	cfg := &config.Config{WordMin: 10, WordMax: 2}
	_, err := config.BuildFilterChain(cfg)
	if err == nil {
		t.Fatal("expected error for invalid word count range")
	}
}

func TestBuildFilterChain_WordCountAndRegex(t *testing.T) {
	cfg := &config.Config{
		Pattern: "error",
		WordMin: 3,
		WordMax: 0,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
}

func TestBuildFilterChain_WordMinOnly(t *testing.T) {
	cfg := &config.Config{WordMin: 1, WordMax: 0}
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
