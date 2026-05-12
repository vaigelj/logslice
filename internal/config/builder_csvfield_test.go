package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
)

func TestBuildFilterChain_WithCSVField(t *testing.T) {
	cfg := &config.Config{
		CSVFieldIndex: 1,
		CSVFieldValue: "ERROR",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
	if !chain.Match("ts,ERROR,msg") {
		t.Error("expected match")
	}
	if chain.Match("ts,INFO,msg") {
		t.Error("expected no match")
	}
}

func TestBuildFilterChain_CSVFieldCaseInsensitive(t *testing.T) {
	cfg := &config.Config{
		CSVFieldIndex:      0,
		CSVFieldValue:      "warn",
		CSVFieldIgnoreCase: true,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !chain.Match("WARN,some message") {
		t.Error("expected case-insensitive match")
	}
}

func TestBuildFilterChain_NoCSVField(t *testing.T) {
	cfg := &config.Config{}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// chain matches everything when no filters configured
	if !chain.Match("any line") {
		t.Error("expected match-all with no filters")
	}
}

func TestBuildFilterChain_CSVFieldInvalidIndex(t *testing.T) {
	cfg := &config.Config{
		CSVFieldIndex: -1,
		CSVFieldValue: "foo",
	}
	_, err := config.BuildFilterChain(cfg)
	if err == nil {
		t.Fatal("expected error for negative CSV field index")
	}
}

func TestBuildFilterChain_CSVFieldAndRegex(t *testing.T) {
	cfg := &config.Config{
		Pattern:       "ERROR",
		CSVFieldIndex: 2,
		CSVFieldValue: "db",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !chain.Match("2024,ERROR,db") {
		t.Error("expected match with both filters satisfied")
	}
	if chain.Match("2024,INFO,db") {
		t.Error("expected no match: regex not satisfied")
	}
}
