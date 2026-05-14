package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
)

func TestBuildFilterChain_WithKVField(t *testing.T) {
	cfg := &config.Config{
		KVKey:   "level",
		KVSep:   "=",
		KVValue: "info",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain.Len() != 1 {
		t.Errorf("expected 1 filter, got %d", chain.Len())
	}
}

func TestBuildFilterChain_KVFieldEmptyKeySkipped(t *testing.T) {
	cfg := &config.Config{
		KVKey:   "",
		KVValue: "info",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain.Len() != 0 {
		t.Errorf("expected 0 filters, got %d", chain.Len())
	}
}

func TestBuildFilterChain_KVFieldCaseInsensitive(t *testing.T) {
	cfg := &config.Config{
		KVKey:             "level",
		KVSep:             "=",
		KVValue:           "INFO",
		KVCaseInsensitive: true,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain.Len() != 1 {
		t.Errorf("expected 1 filter, got %d", chain.Len())
	}
}

func TestBuildFilterChain_KVFieldEmptyValueError(t *testing.T) {
	cfg := &config.Config{
		KVKey:   "level",
		KVSep:   "=",
		KVValue: "",
	}
	_, err := config.BuildFilterChain(cfg)
	if err == nil {
		t.Fatal("expected error for empty kv value when key is set")
	}
}

func TestBuildFilterChain_KVFieldAndRegex(t *testing.T) {
	cfg := &config.Config{
		Pattern: "^2024",
		KVKey:   "status",
		KVSep:   "=",
		KVValue: "ok",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain.Len() != 2 {
		t.Errorf("expected 2 filters, got %d", chain.Len())
	}
}
