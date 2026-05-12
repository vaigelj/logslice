package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewSkipLinesFilter_InvalidN(t *testing.T) {
	for _, n := range []int{0, -1, -100} {
		_, err := filter.NewSkipLinesFilter(n)
		if err == nil {
			t.Errorf("expected error for n=%d, got nil", n)
		}
	}
}

func TestNewSkipLinesFilter_Valid(t *testing.T) {
	f, err := filter.NewSkipLinesFilter(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Skip() != 3 {
		t.Errorf("expected skip=3, got %d", f.Skip())
	}
	if f.Seen() != 0 {
		t.Errorf("expected seen=0, got %d", f.Seen())
	}
}

func TestSkipLinesFilter_Match_SkipsFirst(t *testing.T) {
	f, _ := filter.NewSkipLinesFilter(2)

	if f.Match("line1") {
		t.Error("expected line1 to be skipped")
	}
	if f.Match("line2") {
		t.Error("expected line2 to be skipped")
	}
	if !f.Match("line3") {
		t.Error("expected line3 to pass")
	}
	if !f.Match("line4") {
		t.Error("expected line4 to pass")
	}
}

func TestSkipLinesFilter_Seen_Counter(t *testing.T) {
	f, _ := filter.NewSkipLinesFilter(1)
	for i := 0; i < 5; i++ {
		f.Match("x")
	}
	if f.Seen() != 5 {
		t.Errorf("expected seen=5, got %d", f.Seen())
	}
}

func TestSkipLinesFilter_SkipOne(t *testing.T) {
	f, _ := filter.NewSkipLinesFilter(1)
	if f.Match("first") {
		t.Error("expected first line to be skipped")
	}
	if !f.Match("second") {
		t.Error("expected second line to pass")
	}
}

func TestSkipLinesFilter_InChain(t *testing.T) {
	skip, _ := filter.NewSkipLinesFilter(1)
	regex, _ := filter.NewRegexFilter("keep")
	chain, _ := filter.NewChain(skip, regex)

	lines := []string{"keep this skipped", "keep this", "drop this", "keep another"}
	expected := []bool{false, true, false, true}

	for i, line := range lines {
		got := chain.Match(line)
		if got != expected[i] {
			t.Errorf("line %d %q: expected %v, got %v", i, line, expected[i], got)
		}
	}
}
