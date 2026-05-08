package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
)

func TestBuildFilterChain_WithReplace(t *testing.T) {
	cfg := &config.Config{
		ReplacePattern: "foo",
		ReplaceWith:    "bar",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
}

func TestBuildFilterChain_ReplaceEmptyPatternSkipped(t *testing.T) {
	cfg := &config.Config{
		ReplacePattern: "",
		ReplaceWith:    "bar",
	}
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error when pattern is empty: %v", err)
	}
}

func TestBuildFilterChain_ReplaceWithRegex(t *testing.T) {
	cfg := &config.Config{
		ReplacePattern: `\d+`,
		ReplaceWith:    "NUM",
		ReplaceRegex:   true,
	}
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBuildFilterChain_ReplaceInvalidRegex(t *testing.T) {
	cfg := &config.Config{
		ReplacePattern: "[",
		ReplaceWith:    "x",
		ReplaceRegex:   true,
	}
	_, err := config.BuildFilterChain(cfg)
	if err == nil {
		t.Fatal("expected error for invalid regex pattern")
	}
}

func TestBuildFilterChain_ReplaceAndRegexFilter(t *testing.T) {
	cfg := &config.Config{
		Pattern:        "error",
		ReplacePattern: "error",
		ReplaceWith:    "ERR",
	}
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
