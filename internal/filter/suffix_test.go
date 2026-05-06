package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewSuffixFilter_EmptySuffix(t *testing.T) {
	_, err := filter.NewSuffixFilter("", false)
	if err == nil {
		t.Fatal("expected error for empty suffix, got nil")
	}
}

func TestNewSuffixFilter_ValidSuffix(t *testing.T) {
	f, err := filter.NewSuffixFilter(".log", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f == nil {
		t.Fatal("expected non-nil filter")
	}
}

func TestSuffixFilter_Match_CaseSensitive(t *testing.T) {
	f, _ := filter.NewSuffixFilter("ERROR", false)

	cases := []struct {
		line  string
		want  bool
	}{
		{"something ERROR", true},
		{"something error", false},
		{"ERROR at start", false},
		{"noERROR", true},
	}
	for _, tc := range cases {
		got := f.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestSuffixFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := filter.NewSuffixFilter("ERROR", true)

	cases := []struct {
		line string
		want bool
	}{
		{"something ERROR", true},
		{"something error", true},
		{"something Error", true},
		{"ERROR at start", false},
	}
	for _, tc := range cases {
		got := f.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestSuffixFilter_Accessors(t *testing.T) {
	f, _ := filter.NewSuffixFilter("WARN", true)
	// stored lowercase for case-insensitive
	if f.Suffix() != "warn" {
		t.Errorf("Suffix() = %q, want %q", f.Suffix(), "warn")
	}
	if !f.CaseInsensitive() {
		t.Error("CaseInsensitive() = false, want true")
	}
}

func TestSuffixFilter_InChain(t *testing.T) {
	suffix, _ := filter.NewSuffixFilter(".go", false)
	chain := filter.NewChain(suffix)

	if !chain.Match("main.go") {
		t.Error("expected chain to match 'main.go'")
	}
	if chain.Match("main.py") {
		t.Error("expected chain to reject 'main.py'")
	}
}
