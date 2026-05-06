package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
	"github.com/yourorg/logslice/internal/filter"
)

func TestBuildFilterChain_WithSuffix(t *testing.T) {
	cfg := &config.Config{
		Suffix: ".log",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
	if !chain.Match("app.log") {
		t.Error("expected chain to match 'app.log'")
	}
	if chain.Match("app.txt") {
		t.Error("expected chain to reject 'app.txt'")
	}
}

func TestBuildFilterChain_SuffixCaseInsensitive(t *testing.T) {
	cfg := &config.Config{
		Suffix:          "ERROR",
		CaseInsensitive: true,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !chain.Match("something error") {
		t.Error("expected case-insensitive match for 'something error'")
	}
}

func TestBuildFilterChain_SuffixAndPrefix(t *testing.T) {
	cfg := &config.Config{
		Prefix: "INFO",
		Suffix: "done",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_ = chain // both filters applied; just ensure no error building
}

func TestBuildFilterChain_NoSuffix(t *testing.T) {
	cfg := &config.Config{}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// empty chain matches everything
	var _ filter.Filter = chain
}
