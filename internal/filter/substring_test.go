package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewSubstringFilter_EmptySubstring(t *testing.T) {
	_, err := filter.NewSubstringFilter("", false)
	if err == nil {
		t.Fatal("expected error for empty substring, got nil")
	}
}

func TestNewSubstringFilter_ValidSubstring(t *testing.T) {
	f, err := filter.NewSubstringFilter("ERROR", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Substring() != "ERROR" {
		t.Errorf("expected substring ERROR, got %s", f.Substring())
	}
	if f.CaseInsensitive() {
		t.Error("expected case-sensitive mode")
	}
}

func TestSubstringFilter_Match_CaseSensitive(t *testing.T) {
	f, _ := filter.NewSubstringFilter("ERROR", false)
	tests := []struct {
		line  string
		want  bool
	}{
		{"2024-01-01 ERROR something failed", true},
		{"2024-01-01 error something failed", false},
		{"2024-01-01 INFO all good", false},
		{"ERROR", true},
		{"", false},
	}
	for _, tt := range tests {
		got := f.Match(tt.line)
		if got != tt.want {
			t.Errorf("Match(%q) = %v, want %v", tt.line, got, tt.want)
		}
	}
}

func TestSubstringFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := filter.NewSubstringFilter("error", true)
	tests := []struct {
		line string
		want bool
	}{
		{"2024-01-01 ERROR something failed", true},
		{"2024-01-01 error something failed", true},
		{"2024-01-01 Error something failed", true},
		{"2024-01-01 INFO all good", false},
	}
	for _, tt := range tests {
		got := f.Match(tt.line)
		if got != tt.want {
			t.Errorf("Match(%q) = %v, want %v", tt.line, got, tt.want)
		}
	}
}

func TestSubstringFilter_InChain(t *testing.T) {
	sub, _ := filter.NewSubstringFilter("ERROR", false)
	chain := filter.NewChain(sub)
	if !chain.Match("ERROR: disk full") {
		t.Error("expected chain to match line containing ERROR")
	}
	if chain.Match("INFO: all good") {
		t.Error("expected chain to reject line without ERROR")
	}
}
