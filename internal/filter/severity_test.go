package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewSeverityFilter_EmptyLevel(t *testing.T) {
	_, err := filter.NewSeverityFilter("")
	if err == nil {
		t.Fatal("expected error for empty level")
	}
}

func TestNewSeverityFilter_UnknownLevel(t *testing.T) {
	_, err := filter.NewSeverityFilter("verbose")
	if err == nil {
		t.Fatal("expected error for unknown level")
	}
}

func TestNewSeverityFilter_Valid(t *testing.T) {
	f, err := filter.NewSeverityFilter("warn")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.MinLevel() != "warn" {
		t.Errorf("expected minLevel=warn, got %s", f.MinLevel())
	}
}

func TestSeverityFilter_Match_AboveThreshold(t *testing.T) {
	f, _ := filter.NewSeverityFilter("warn")
	lines := []string{
		"2024-01-01 ERROR something failed",
		"2024-01-01 FATAL system crash",
		"level=warn msg=disk full",
	}
	for _, line := range lines {
		if !f.Match(line) {
			t.Errorf("expected match for line: %s", line)
		}
	}
}

func TestSeverityFilter_Match_BelowThreshold(t *testing.T) {
	f, _ := filter.NewSeverityFilter("warn")
	lines := []string{
		"2024-01-01 DEBUG connecting to db",
		"2024-01-01 INFO server started",
		"level=trace msg=enter function",
	}
	for _, line := range lines {
		if f.Match(line) {
			t.Errorf("expected no match for line: %s", line)
		}
	}
}

func TestSeverityFilter_Match_NoKeyword(t *testing.T) {
	f, _ := filter.NewSeverityFilter("info")
	if f.Match("plain log line without level") {
		t.Error("expected no match for line without severity keyword")
	}
}

func TestSeverityFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := filter.NewSeverityFilter("error")
	if !f.Match("[ERROR] something bad happened") {
		t.Error("expected match for uppercase ERROR")
	}
	if !f.Match("[error] lowercase") {
		t.Error("expected match for lowercase error")
	}
}

func TestSeverityFilter_InChain(t *testing.T) {
	sev, _ := filter.NewSeverityFilter("error")
	rx, _ := filter.NewRegexFilter(`database`)
	chain := filter.NewChain(sev, rx)
	if !chain.Match("ERROR database connection lost") {
		t.Error("expected chain match")
	}
	if chain.Match("ERROR network timeout") {
		t.Error("expected chain rejection when regex fails")
	}
}
