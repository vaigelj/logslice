package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewExcludeAnyFilter_NoTerms(t *testing.T) {
	_, err := filter.NewExcludeAnyFilter([]string{}, false)
	if err == nil {
		t.Fatal("expected error for empty terms, got nil")
	}
}

func TestNewExcludeAnyFilter_Valid(t *testing.T) {
	f, err := filter.NewExcludeAnyFilter([]string{"ERROR", "WARN"}, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(f.Terms()) != 2 {
		t.Fatalf("expected 2 terms, got %d", len(f.Terms()))
	}
}

func TestExcludeAnyFilter_Match_CaseSensitive(t *testing.T) {
	f, _ := filter.NewExcludeAnyFilter([]string{"ERROR", "FATAL"}, false)

	tests := []struct {
		line string
		want bool
	}{
		{"INFO everything is fine", true},
		{"ERROR something broke", false},
		{"FATAL crash", false},
		{"error lowercase should pass", true}, // case-sensitive
	}
	for _, tc := range tests {
		if got := f.Match(tc.line); got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestExcludeAnyFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := filter.NewExcludeAnyFilter([]string{"error", "warn"}, true)

	tests := []struct {
		line string
		want bool
	}{
		{"INFO all good", true},
		{"ERROR something broke", false},
		{"Warning: disk low", false},
		{"no issues here", true},
	}
	for _, tc := range tests {
		if got := f.Match(tc.line); got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestExcludeAnyFilter_Accessors(t *testing.T) {
	f, _ := filter.NewExcludeAnyFilter([]string{"bad"}, true)
	if !f.CaseInsensitive() {
		t.Error("expected CaseInsensitive() to be true")
	}
	if f.Terms()[0] != "bad" {
		t.Errorf("expected term 'bad', got %q", f.Terms()[0])
	}
}

func TestExcludeAnyFilter_InChain(t *testing.T) {
	exclude, _ := filter.NewExcludeAnyFilter([]string{"DROP"}, false)
	chain, _ := filter.NewChain(exclude)

	if !chain.Match("KEEP this line") {
		t.Error("expected chain to accept line without excluded term")
	}
	if chain.Match("DROP this line") {
		t.Error("expected chain to reject line with excluded term")
	}
}
