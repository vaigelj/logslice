package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewMinLengthFilter_NegativeMin(t *testing.T) {
	_, err := filter.NewMinLengthFilter(-1)
	if err == nil {
		t.Fatal("expected error for negative minLen, got nil")
	}
}

func TestNewMinLengthFilter_ZeroMin(t *testing.T) {
	f, err := filter.NewMinLengthFilter(0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.MinLen() != 0 {
		t.Errorf("expected MinLen 0, got %d", f.MinLen())
	}
}

func TestNewMinLengthFilter_Valid(t *testing.T) {
	f, err := filter.NewMinLengthFilter(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.MinLen() != 10 {
		t.Errorf("expected MinLen 10, got %d", f.MinLen())
	}
}

func TestMinLengthFilter_Match(t *testing.T) {
	tests := []struct {
		name   string
		minLen int
		line   string
		want   bool
	}{
		{"exact length", 5, "hello", true},
		{"longer than min", 3, "hello world", true},
		{"shorter than min", 10, "hi", false},
		{"empty line zero min", 0, "", true},
		{"empty line nonzero min", 1, "", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f, err := filter.NewMinLengthFilter(tc.minLen)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := f.Match(tc.line); got != tc.want {
				t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
			}
		})
	}
}

func TestMinLengthFilter_InChain(t *testing.T) {
	minF, _ := filter.NewMinLengthFilter(5)
	chain, err := filter.NewChain(minF)
	if err != nil {
		t.Fatalf("NewChain error: %v", err)
	}
	if chain.Match("hi") {
		t.Error("expected short line to be rejected by chain")
	}
	if !chain.Match("hello world") {
		t.Error("expected long line to pass chain")
	}
}
