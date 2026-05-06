package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewFieldMatchFilter_EmptyDelimiter(t *testing.T) {
	_, err := filter.NewFieldMatchFilter("", 0, "value", false)
	if err == nil {
		t.Fatal("expected error for empty delimiter")
	}
}

func TestNewFieldMatchFilter_NegativeIndex(t *testing.T) {
	_, err := filter.NewFieldMatchFilter(" ", -1, "value", false)
	if err == nil {
		t.Fatal("expected error for negative field index")
	}
}

func TestNewFieldMatchFilter_EmptyValue(t *testing.T) {
	_, err := filter.NewFieldMatchFilter(" ", 0, "", false)
	if err == nil {
		t.Fatal("expected error for empty value")
	}
}

func TestFieldMatchFilter_Match_CaseSensitive(t *testing.T) {
	f, err := filter.NewFieldMatchFilter(" ", 1, "ERROR", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	tests := []struct {
		line string
		want bool
	}{
		{"2024-01-01 ERROR something happened", true},
		{"2024-01-01 INFO nothing here", false},
		{"2024-01-01 error lowercase", false},
		{"singleword", false},
	}
	for _, tt := range tests {
		got := f.Match(tt.line)
		if got != tt.want {
			t.Errorf("Match(%q) = %v, want %v", tt.line, got, tt.want)
		}
	}
}

func TestFieldMatchFilter_Match_CaseInsensitive(t *testing.T) {
	f, err := filter.NewFieldMatchFilter(" ", 1, "error", true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !f.Match("2024-01-01 ERROR something") {
		t.Error("expected case-insensitive match for ERROR")
	}
	if !f.Match("2024-01-01 error something") {
		t.Error("expected match for error")
	}
}

func TestFieldMatchFilter_Match_FieldOutOfRange(t *testing.T) {
	f, err := filter.NewFieldMatchFilter("|", 5, "val", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Match("a|b|c") {
		t.Error("expected no match when field index out of range")
	}
}

func TestFieldMatchFilter_Accessors(t *testing.T) {
	f, _ := filter.NewFieldMatchFilter("|", 2, "WARN", true)
	if f.Delimiter() != "|" {
		t.Errorf("Delimiter() = %q, want %q", f.Delimiter(), "|")
	}
	if f.FieldIndex() != 2 {
		t.Errorf("FieldIndex() = %d, want 2", f.FieldIndex())
	}
	if f.Value() != "WARN" {
		t.Errorf("Value() = %q, want %q", f.Value(), "WARN")
	}
	if !f.CaseInsensitive() {
		t.Error("CaseInsensitive() should be true")
	}
}

func TestFieldMatchFilter_InChain(t *testing.T) {
	f1, _ := filter.NewFieldMatchFilter(" ", 1, "ERROR", false)
	f2, _ := filter.NewSubstringFilter("disk", false)
	chain, err := filter.NewChain(f1, f2)
	if err != nil {
		t.Fatalf("unexpected error building chain: %v", err)
	}
	if !chain.Match("2024-01-01 ERROR disk full") {
		t.Error("expected chain match")
	}
	if chain.Match("2024-01-01 ERROR memory full") {
		t.Error("expected chain rejection on missing substring")
	}
}
