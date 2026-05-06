package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewLineNumFilter_NegativeMin(t *testing.T) {
	_, err := filter.NewLineNumFilter(-1, 10)
	if err == nil {
		t.Fatal("expected error for negative min")
	}
}

func TestNewLineNumFilter_NegativeMax(t *testing.T) {
	_, err := filter.NewLineNumFilter(1, -5)
	if err == nil {
		t.Fatal("expected error for negative max")
	}
}

func TestNewLineNumFilter_MinExceedsMax(t *testing.T) {
	_, err := filter.NewLineNumFilter(10, 5)
	if err == nil {
		t.Fatal("expected error when min > max")
	}
}

func TestNewLineNumFilter_Valid(t *testing.T) {
	f, err := filter.NewLineNumFilter(2, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Min() != 2 || f.Max() != 5 {
		t.Errorf("unexpected bounds: min=%d max=%d", f.Min(), f.Max())
	}
}

func TestLineNumFilter_Match_Range(t *testing.T) {
	f, _ := filter.NewLineNumFilter(2, 4)
	lines := []string{"a", "b", "c", "d", "e"}
	expected := []bool{false, true, true, true, false}

	for i, line := range lines {
		got := f.Match(line)
		if got != expected[i] {
			t.Errorf("line %d: expected %v got %v", i+1, expected[i], got)
		}
	}
}

func TestLineNumFilter_Match_NoMin(t *testing.T) {
	f, _ := filter.NewLineNumFilter(0, 3)
	for i := 0; i < 3; i++ {
		if !f.Match("x") {
			t.Errorf("line %d should match", i+1)
		}
	}
	if f.Match("x") {
		t.Error("line 4 should not match")
	}
}

func TestLineNumFilter_Match_NoMax(t *testing.T) {
	f, _ := filter.NewLineNumFilter(3, 0)
	if f.Match("x") || f.Match("x") {
		t.Error("lines 1-2 should not match")
	}
	for i := 0; i < 5; i++ {
		if !f.Match("x") {
			t.Errorf("line %d should match", i+3)
		}
	}
}

func TestLineNumFilter_Current(t *testing.T) {
	f, _ := filter.NewLineNumFilter(1, 10)
	f.Match("a")
	f.Match("b")
	f.Match("c")
	if f.Current() != 3 {
		t.Errorf("expected current=3, got %d", f.Current())
	}
}
