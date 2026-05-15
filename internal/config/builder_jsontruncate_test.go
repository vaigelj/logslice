package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
)

func TestBuildFilterChain_WithJSONTruncate(t *testing.T) {
	cfg := &config.Config{
		JSONTruncateField:  "message",
		JSONTruncateMaxLen: 20,
		JSONTruncateSuffix: "...",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
}

func TestBuildFilterChain_JSONTruncateEmptyFieldSkipped(t *testing.T) {
	cfg := &config.Config{
		JSONTruncateField:  "",
		JSONTruncateMaxLen: 20,
	}
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBuildFilterChain_JSONTruncateZeroMaxLenSkipped(t *testing.T) {
	cfg := &config.Config{
		JSONTruncateField:  "msg",
		JSONTruncateMaxLen: 0,
	}
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBuildFilterChain_JSONTruncateSuffixTooLong(t *testing.T) {
	cfg := &config.Config{
		JSONTruncateField:  "msg",
		JSONTruncateMaxLen: 3,
		JSONTruncateSuffix: "...",
	}
	_, err := config.BuildFilterChain(cfg)
	if err == nil {
		t.Fatal("expected error for suffix length >= maxLen")
	}
}

func TestBuildFilterChain_JSONTruncateAndRegex(t *testing.T) {
	cfg := &config.Config{
		Pattern:            "ERROR",
		JSONTruncateField:  "msg",
		JSONTruncateMaxLen: 50,
		JSONTruncateSuffix: "...",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
}
