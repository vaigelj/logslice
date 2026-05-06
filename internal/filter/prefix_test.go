package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewPrefixFilter_EmptyPrefix(t *testing.T) {
	_, err := filter.NewPrefixFilter("", false)
	if err == nil {
		t.Fatal("expected error for empty prefix, got nil")
	}
}

func TestNewPrefixFilter_ValidPrefix(t *testing.T) {
	f, err := filter.NewPrefixFilter("ERROR", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Prefix() != "ERROR" {
		t.Errorf("expected prefix ERROR, got %q", f.Prefix())
	}
	if f.CaseFold() {
		t.Error("expected caseFold false")
	}
}

func TestPrefixFilter_Match_CaseSensitive(t *testing.T) {
	f, _ := filter.NewPrefixFilter("ERROR", false)
	tests := []struct {
		line string
		want bool
	}{
		{"ERROR something happened", true},
		{"error something happened", false},
		{"INFO not an error", false},
		{"ERROR", true},
		{"", false},
	}
	for _, tc := range tests {
		got := f.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestPrefixFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := filter.NewPrefixFilter("error", true)
	tests := []struct {
		line string
		want bool
	}{
		{"ERROR something", true},
		{"error something", true},
		{"Error something", true},
		{"INFO not matching", false},
		{"", false},
	}
	for _, tc := range tests {
		got := f.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestPrefixFilter_InChain(t *testing.T) {
	pf, _ := filter.NewPrefixFilter("2024", false)
	rf, _ := filter.NewRegexFilter(`ERROR`)
	chain := filter.NewChain(pf, rf)

	if !chain.Match("2024-01-01 ERROR something") {
		t.Error("expected chain to match line with prefix and regex")
	}
	if chain.Match("2024-01-01 INFO something") {
		t.Error("expected chain to reject line missing regex match")
	}
	if chain.Match("1999-01-01 ERROR something") {
		t.Error("expected chain to reject line missing prefix")
	}
}
