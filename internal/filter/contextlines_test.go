package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewContextLinesFilter_NilAnchor(t *testing.T) {
	_, err := filter.NewContextLinesFilter(nil, 1, 1)
	if err == nil {
		t.Fatal("expected error for nil anchor")
	}
}

func TestNewContextLinesFilter_NegativeBefore(t *testing.T) {
	anchor, _ := filter.NewSubstringFilter("ERROR", false)
	_, err := filter.NewContextLinesFilter(anchor, -1, 0)
	if err == nil {
		t.Fatal("expected error for negative before")
	}
}

func TestNewContextLinesFilter_NegativeAfter(t *testing.T) {
	anchor, _ := filter.NewSubstringFilter("ERROR", false)
	_, err := filter.NewContextLinesFilter(anchor, 0, -1)
	if err == nil {
		t.Fatal("expected error for negative after")
	}
}

func TestContextLinesFilter_Accessors(t *testing.T) {
	anchor, _ := filter.NewSubstringFilter("ERROR", false)
	f, err := filter.NewContextLinesFilter(anchor, 2, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Before() != 2 {
		t.Errorf("Before() = %d, want 2", f.Before())
	}
	if f.After() != 3 {
		t.Errorf("After() = %d, want 3", f.After())
	}
	if f.Anchor() != anchor {
		t.Error("Anchor() did not return the original anchor filter")
	}
}

func TestContextLinesFilter_AfterContext(t *testing.T) {
	anchor, _ := filter.NewSubstringFilter("ERROR", false)
	f, _ := filter.NewContextLinesFilter(anchor, 0, 2)

	lines := []string{"info a", "ERROR here", "after1", "after2", "after3", "normal"}
	want := []bool{false, true, true, true, false, false}

	for i, line := range lines {
		got := f.Match(line)
		if got != want[i] {
			t.Errorf("line %d %q: Match()=%v want %v", i, line, got, want[i])
		}
	}
}

func TestContextLinesFilter_MatchOnlyAnchor(t *testing.T) {
	anchor, _ := filter.NewSubstringFilter("HIT", false)
	f, _ := filter.NewContextLinesFilter(anchor, 0, 0)

	lines := []string{"miss", "HIT", "miss", "HIT", "miss"}
	want := []bool{false, true, false, true, false}

	for i, line := range lines {
		if got := f.Match(line); got != want[i] {
			t.Errorf("line %d %q: Match()=%v want %v", i, line, got, want[i])
		}
	}
}

func TestContextLinesFilter_ConsecutiveMatches(t *testing.T) {
	anchor, _ := filter.NewSubstringFilter("ERR", false)
	f, _ := filter.NewContextLinesFilter(anchor, 0, 1)

	lines := []string{"ERR1", "ERR2", "tail", "skip"}
	want := []bool{true, true, true, false}

	for i, line := range lines {
		if got := f.Match(line); got != want[i] {
			t.Errorf("line %d %q: Match()=%v want %v", i, line, got, want[i])
		}
	}
}
