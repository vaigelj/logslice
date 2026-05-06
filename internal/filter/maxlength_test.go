package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewMaxLengthFilter_NegativeMax(t *testing.T) {
	_, err := filter.NewMaxLengthFilter(-1)
	if err == nil {
		t.Fatal("expected error for negative max, got nil")
	}
}

func TestNewMaxLengthFilter_ZeroMax(t *testing.T) {
	f, err := filter.NewMaxLengthFilter(0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Max() != 0 {
		t.Errorf("expected max 0, got %d", f.Max())
	}
}

func TestNewMaxLengthFilter_Valid(t *testing.T) {
	f, err := filter.NewMaxLengthFilter(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Max() != 10 {
		t.Errorf("expected max 10, got %d", f.Max())
	}
}

func TestMaxLengthFilter_Match(t *testing.T) {
	tests := []struct {
		name  string
		max   int
		line  string
		want  bool
	}{
		{"exactly at limit", 5, "hello", true},
		{"under limit", 5, "hi", true},
		{"over limit", 5, "toolong", false},
		{"empty line zero max", 0, "", true},
		{"non-empty zero max", 0, "x", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f, err := filter.NewMaxLengthFilter(tc.max)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := f.Match(tc.line); got != tc.want {
				t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
			}
		})
	}
}

func TestMaxLengthFilter_InChain(t *testing.T) {
	min, err := filter.NewMinLengthFilter(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	max, err := filter.NewMaxLengthFilter(6)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	chain := filter.NewChain(min, max)

	if chain.Match("hi") {
		t.Error("expected 'hi' (len 2) to be rejected by min filter")
	}
	if !chain.Match("hello") {
		t.Error("expected 'hello' (len 5) to pass both filters")
	}
	if chain.Match("toolong!") {
		t.Error("expected 'toolong!' (len 8) to be rejected by max filter")
	}
}
