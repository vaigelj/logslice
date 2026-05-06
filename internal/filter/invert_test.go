package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewInvertFilter_NilInner(t *testing.T) {
	_, err := filter.NewInvertFilter(nil)
	if err == nil {
		t.Fatal("expected error for nil inner filter, got nil")
	}
}

func TestInvertFilter_MatchNegates(t *testing.T) {
	regex, err := filter.NewRegexFilter("ERROR")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	inv, err := filter.NewInvertFilter(regex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		line     string
		wantMatch bool
	}{
		{"2024-01-01 ERROR something failed", false},
		{"2024-01-01 INFO all good", true},
		{"ERROR", false},
		{"", true},
	}

	for _, tc := range tests {
		got := inv.Match([]byte(tc.line))
		if got != tc.wantMatch {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.wantMatch)
		}
	}
}

func TestInvertFilter_Inner(t *testing.T) {
	regex, err := filter.NewRegexFilter("WARN")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	inv, err := filter.NewInvertFilter(regex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if inv.Inner() != regex {
		t.Error("Inner() did not return the wrapped filter")
	}
}

func TestInvertFilter_InChain(t *testing.T) {
	// Only lines that contain INFO but NOT DEBUG should pass.
	infoFilter, _ := filter.NewRegexFilter("INFO")
	debugFilter, _ := filter.NewRegexFilter("DEBUG")
	notDebug, _ := filter.NewInvertFilter(debugFilter)

	chain := filter.NewChain(infoFilter, notDebug)

	if !chain.Match([]byte("INFO user logged in")) {
		t.Error("expected match for INFO-only line")
	}
	if chain.Match([]byte("INFO DEBUG verbose")) {
		t.Error("expected no match for line containing both INFO and DEBUG")
	}
	if chain.Match([]byte("ERROR something")) {
		t.Error("expected no match for ERROR line")
	}
}
