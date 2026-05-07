package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewContainsAnyFilter_NoTerms(t *testing.T) {
	_, err := filter.NewContainsAnyFilter([]string{}, false)
	if err == nil {
		t.Fatal("expected error for empty terms, got nil")
	}
}

func TestNewContainsAnyFilter_Valid(t *testing.T) {
	f, err := filter.NewContainsAnyFilter([]string{"error", "warn"}, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f == nil {
		t.Fatal("expected non-nil filter")
	}
}

func TestContainsAnyFilter_Match_CaseSensitive(t *testing.T) {
	f, _ := filter.NewContainsAnyFilter([]string{"ERROR", "WARN"}, false)
	tests := []struct {
		line  string
		want  bool
	}{
		{"2024/01/01 ERROR something failed", true},
		{"2024/01/01 WARN low disk", true},
		{"2024/01/01 INFO all good", false},
		{"error in lowercase", false},
	}
	for _, tc := range tests {
		got := f.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestContainsAnyFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := filter.NewContainsAnyFilter([]string{"error", "warn"}, true)
	tests := []struct {
		line string
		want bool
	}{
		{"2024/01/01 ERROR something failed", true},
		{"2024/01/01 WARN low disk", true},
		{"mixed Error case", true},
		{"2024/01/01 INFO all good", false},
	}
	for _, tc := range tests {
		got := f.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestContainsAnyFilter_Accessors(t *testing.T) {
	terms := []string{"foo", "bar"}
	f, _ := filter.NewContainsAnyFilter(terms, true)
	if f.CaseInsensitive() != true {
		t.Error("expected CaseInsensitive() == true")
	}
	if len(f.Terms()) != 2 {
		t.Errorf("expected 2 terms, got %d", len(f.Terms()))
	}
}

func TestContainsAnyFilter_InChain(t *testing.T) {
	caf, _ := filter.NewContainsAnyFilter([]string{"error"}, false)
	chain, _ := filter.NewChain(caf)
	if chain.Match("critical error occurred") != true {
		t.Error("chain should match line containing 'error'")
	}
	if chain.Match("everything is fine") != false {
		t.Error("chain should not match line without 'error'")
	}
}
