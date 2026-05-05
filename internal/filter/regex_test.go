package filter

import (
	"testing"
)

func TestNewRegexFilter_EmptyPattern(t *testing.T) {
	_, err := NewRegexFilter("", false)
	if err == nil {
		t.Fatal("expected error for empty pattern, got nil")
	}
}

func TestNewRegexFilter_InvalidPattern(t *testing.T) {
	_, err := NewRegexFilter("[invalid", false)
	if err == nil {
		t.Fatal("expected error for invalid regex, got nil")
	}
}

func TestRegexFilter_Match(t *testing.T) {
	tests := []struct {
		pattern string
		invert  bool
		line    string
		want    bool
	}{
		{pattern: "ERROR", invert: false, line: "2024-01-01 ERROR something broke", want: true},
		{pattern: "ERROR", invert: false, line: "2024-01-01 INFO all good", want: false},
		{pattern: "ERROR", invert: true, line: "2024-01-01 INFO all good", want: true},
		{pattern: "ERROR", invert: true, line: "2024-01-01 ERROR something broke", want: false},
		{pattern: `\d{4}-\d{2}-\d{2}`, invert: false, line: "2024-06-15 DEBUG msg", want: true},
		{pattern: `\d{4}-\d{2}-\d{2}`, invert: false, line: "no date here", want: false},
	}

	for _, tc := range tests {
		f, err := NewRegexFilter(tc.pattern, tc.invert)
		if err != nil {
			t.Fatalf("NewRegexFilter(%q, %v) unexpected error: %v", tc.pattern, tc.invert, err)
		}
		got := f.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) with pattern=%q invert=%v: got %v, want %v",
				tc.line, tc.pattern, tc.invert, got, tc.want)
		}
	}
}

func TestRegexFilter_Accessors(t *testing.T) {
	f, err := NewRegexFilter("WARN", true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Pattern() != "WARN" {
		t.Errorf("Pattern() = %q, want %q", f.Pattern(), "WARN")
	}
	if !f.Inverted() {
		t.Error("Inverted() = false, want true")
	}
}
