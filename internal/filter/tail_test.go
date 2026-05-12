package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewTailFilter_InvalidMax(t *testing.T) {
	for _, bad := range []int{0, -1, -100} {
		_, err := filter.NewTailFilter(bad)
		if err == nil {
			t.Errorf("expected error for max=%d", bad)
		}
	}
}

func TestNewTailFilter_Valid(t *testing.T) {
	f, err := filter.NewTailFilter(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Max() != 5 {
		t.Errorf("expected Max()=5, got %d", f.Max())
	}
	if f.Seen() != 0 {
		t.Errorf("expected Seen()=0, got %d", f.Seen())
	}
}

func TestTailFilter_Match_AlwaysTrue(t *testing.T) {
	f, _ := filter.NewTailFilter(3)
	for _, line := range []string{"a", "b", "c", "d"} {
		if !f.Match(line) {
			t.Errorf("Match(%q) should always return true", line)
		}
	}
}

func TestTailFilter_Lines_FewerThanMax(t *testing.T) {
	f, _ := filter.NewTailFilter(5)
	f.Match("line1")
	f.Match("line2")
	got := f.Lines()
	if len(got) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(got))
	}
	if got[0] != "line1" || got[1] != "line2" {
		t.Errorf("unexpected lines: %v", got)
	}
}

func TestTailFilter_Lines_ExactlyMax(t *testing.T) {
	f, _ := filter.NewTailFilter(3)
	f.Match("a")
	f.Match("b")
	f.Match("c")
	got := f.Lines()
	expected := []string{"a", "b", "c"}
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("pos %d: want %q got %q", i, v, got[i])
		}
	}
}

func TestTailFilter_Lines_MoreThanMax(t *testing.T) {
	f, _ := filter.NewTailFilter(3)
	for _, l := range []string{"1", "2", "3", "4", "5"} {
		f.Match(l)
	}
	got := f.Lines()
	if len(got) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(got))
	}
	expected := []string{"3", "4", "5"}
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("pos %d: want %q got %q", i, v, got[i])
		}
	}
}

func TestTailFilter_Lines_EmptyInput(t *testing.T) {
	f, _ := filter.NewTailFilter(3)
	if got := f.Lines(); got != nil {
		t.Errorf("expected nil for empty input, got %v", got)
	}
}

func TestTailFilter_Seen_Counter(t *testing.T) {
	f, _ := filter.NewTailFilter(2)
	for i := 0; i < 7; i++ {
		f.Match("x")
	}
	if f.Seen() != 7 {
		t.Errorf("expected Seen()=7, got %d", f.Seen())
	}
}
