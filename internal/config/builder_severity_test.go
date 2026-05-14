package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
	"github.com/yourorg/logslice/internal/filter"
)

func TestBuildFilterChain_WithSeverity(t *testing.T) {
	cfg := &config.Config{MinSeverity: "error"}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
	if !chain.Match("ERROR disk full") {
		t.Error("expected match for ERROR line")
	}
	if chain.Match("INFO server started") {
		t.Error("expected no match for INFO line")
	}
}

func TestBuildFilterChain_SeverityEmptySkipped(t *testing.T) {
	cfg := &config.Config{MinSeverity: ""}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// No severity filter means all lines pass (chain is empty or passes everything)
	_ = chain
}

func TestBuildFilterChain_SeverityInvalidLevel(t *testing.T) {
	cfg := &config.Config{MinSeverity: "critical"}
	_, err := config.BuildFilterChain(cfg)
	if err == nil {
		t.Fatal("expected error for unknown severity level")
	}
}

func TestBuildFilterChain_SeverityAndRegex(t *testing.T) {
	cfg := &config.Config{
		MinSeverity: "warn",
		Pattern:     "timeout",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !chain.Match("WARN connection timeout reached") {
		t.Error("expected match for warn+timeout line")
	}
	if chain.Match("WARN disk space low") {
		t.Error("expected no match: warn but no 'timeout'")
	}
	if chain.Match("INFO connection timeout") {
		t.Error("expected no match: has timeout but below warn")
	}
}

func TestBuildFilterChain_SeverityDebugPassesAll(t *testing.T) {
	cfg := &config.Config{MinSeverity: "trace"}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	levels := []string{"trace", "debug", "info", "warn", "error", "fatal"}
	for _, lvl := range levels {
		line := lvl + " some message"
		if !chain.Match(line) {
			t.Errorf("expected match for level %s with min=trace", lvl)
		}
	}
}

var _ filter.Filter = (*filter.SeverityFilter)(nil)
