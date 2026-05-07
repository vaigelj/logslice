package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
)

func TestBuildFilterChain_ContextAfterOnly(t *testing.T) {
	cfg := &config.Config{
		Regex:        "ERROR",
		ContextAfter: 2,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
}

func TestBuildFilterChain_ContextBeforeOnly(t *testing.T) {
	cfg := &config.Config{
		Regex:         "WARN",
		ContextBefore: 1,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
}

func TestBuildFilterChain_ContextRequiresRegex(t *testing.T) {
	cfg := &config.Config{
		ContextAfter: 1,
	}
	_, err := config.BuildFilterChain(cfg)
	if err == nil {
		t.Fatal("expected error when context flags used without --regex")
	}
}

func TestBuildFilterChain_NoContext(t *testing.T) {
	cfg := &config.Config{}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
}

func TestBuildFilterChain_ContextBothSides(t *testing.T) {
	cfg := &config.Config{
		Regex:         "CRITICAL",
		ContextBefore: 2,
		ContextAfter:  3,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
}
