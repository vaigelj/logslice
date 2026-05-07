package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
	"github.com/yourorg/logslice/internal/filter"
)

func TestBuildFilterChain_WithTruncate(t *testing.T) {
	cfg := &config.Config{
		TruncateAt:     20,
		TruncateSuffix: "...",
	}
	chain := filter.NewChain()
	if err := config.ExportAddTruncateFilter(cfg, chain); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain.Len() != 1 {
		t.Errorf("expected 1 filter, got %d", chain.Len())
	}
}

func TestBuildFilterChain_TruncateZeroSkipped(t *testing.T) {
	cfg := &config.Config{TruncateAt: 0}
	chain := filter.NewChain()
	if err := config.ExportAddTruncateFilter(cfg, chain); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain.Len() != 0 {
		t.Errorf("expected 0 filters, got %d", chain.Len())
	}
}

func TestBuildFilterChain_TruncateNegativeSkipped(t *testing.T) {
	cfg := &config.Config{TruncateAt: -1}
	chain := filter.NewChain()
	if err := config.ExportAddTruncateFilter(cfg, chain); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain.Len() != 0 {
		t.Errorf("expected 0 filters, got %d", chain.Len())
	}
}

func TestBuildFilterChain_TruncateSuffixTooLong(t *testing.T) {
	cfg := &config.Config{
		TruncateAt:     2,
		TruncateSuffix: "...",
	}
	chain := filter.NewChain()
	err := config.ExportAddTruncateFilter(cfg, chain)
	if err == nil {
		t.Fatal("expected error when suffix exceeds maxLen")
	}
}

func TestBuildFilterChain_TruncateAndRegex(t *testing.T) {
	cfg := &config.Config{
		Pattern:        "ERROR",
		TruncateAt:     50,
		TruncateSuffix: "[...]",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// regex filter + truncate filter
	if chain.Len() < 2 {
		t.Errorf("expected at least 2 filters, got %d", chain.Len())
	}
}
