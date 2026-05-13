package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewMultiRegexFilter_NoPatterns(t *testing.T) {
	_, err := filter.NewMultiRegexFilter([]string{})
	if err == nil {
		t.Fatal("expected error for empty pattern slice")
	}
}

func TestNewMultiRegexFilter_EmptyPattern(t *testing.T) {
	_, err := filter.NewMultiRegexFilter([]string{"foo", ""})
	if err == nil {
		t.Fatal("expected error for empty pattern string")
	}
}

func TestNewMultiRegexFilter_InvalidPattern(t *testing.T) {
	_, err := filter.NewMultiRegexFilter([]string{"valid", "["})
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
}

func TestNewMultiRegexFilter_Valid(t *testing.T) {
	f, err := filter.NewMultiRegexFilter([]string{`\d+`, `error`})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(f.Patterns()) != 2 {
		t.Fatalf("expected 2 patterns, got %d", len(f.Patterns()))
	}
}

func TestMultiRegexFilter_Match_AllMatch(t *testing.T) {
	f, _ := filter.NewMultiRegexFilter([]string{`error`, `\d+`})
	if !f.Match("error code 42") {
		t.Error("expected match when all patterns present")
	}
}

func TestMultiRegexFilter_Match_OneAbsent(t *testing.T) {
	f, _ := filter.NewMultiRegexFilter([]string{`error`, `\d+`})
	if f.Match("error occurred") {
		t.Error("expected no match when one pattern is absent")
	}
}

func TestMultiRegexFilter_Match_NoneMatch(t *testing.T) {
	f, _ := filter.NewMultiRegexFilter([]string{`error`, `warn`})
	if f.Match("info message") {
		t.Error("expected no match")
	}
}

func TestMultiRegexFilter_Transform_Passthrough(t *testing.T) {
	f, _ := filter.NewMultiRegexFilter([]string{`.*`})
	line := "hello world"
	if got := f.Transform(line); got != line {
		t.Errorf("expected %q, got %q", line, got)
	}
}

func TestMultiRegexFilter_InChain(t *testing.T) {
	mf, _ := filter.NewMultiRegexFilter([]string{`error`, `critical`})
	chain, _ := filter.NewChain(mf)
	if !chain.Match("critical error detected") {
		t.Error("expected chain match")
	}
	if chain.Match("error only") {
		t.Error("expected chain reject when second pattern absent")
	}
}
