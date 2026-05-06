package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
)

func TestBuildFilterChain_WithSubstring(t *testing.T) {
	cfg := &config.Config{
		Substring: "ERROR",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !chain.Match("2024-01-01 ERROR disk full") {
		t.Error("expected chain to match line containing ERROR")
	}
	if chain.Match("2024-01-01 INFO all good") {
		t.Error("expected chain to reject line without ERROR")
	}
}

func TestBuildFilterChain_SubstringCaseInsensitive(t *testing.T) {
	cfg := &config.Config{
		Substring:       "error",
		CaseInsensitive: true,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !chain.Match("2024-01-01 ERROR disk full") {
		t.Error("expected chain to match ERROR (case-insensitive)")
	}
	if !chain.Match("2024-01-01 error disk full") {
		t.Error("expected chain to match error (case-insensitive)")
	}
}

func TestBuildFilterChain_SubstringAndPrefix(t *testing.T) {
	cfg := &config.Config{
		Prefix:    "2024",
		Substring: "ERROR",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !chain.Match("2024-01-01 ERROR disk full") {
		t.Error("expected chain to match line with correct prefix and substring")
	}
	if chain.Match("2023-01-01 ERROR disk full") {
		t.Error("expected chain to reject line with wrong prefix")
	}
	if chain.Match("2024-01-01 INFO all good") {
		t.Error("expected chain to reject line without substring")
	}
}

func TestBuildFilterChain_NoSubstring(t *testing.T) {
	cfg := &config.Config{}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !chain.Match("any line should pass") {
		t.Error("expected empty chain to match any line")
	}
}
