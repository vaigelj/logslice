package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewHeadFilter_InvalidMax(t *testing.T) {
	_, err := filter.NewHeadFilter(0)
	if err == nil {
		t.Fatal("expected error for maxLines=0")
	}
	_, err = filter.NewHeadFilter(-5)
	if err == nil {
		t.Fatal("expected error for maxLines=-5")
	}
}

func TestNewHeadFilter_Valid(t *testing.T) {
	f, err := filter.NewHeadFilter(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.MaxLines() != 3 {
		t.Errorf("expected MaxLines=3, got %d", f.MaxLines())
	}
	if f.Seen() != 0 {
		t.Errorf("expected Seen=0, got %d", f.Seen())
	}
}

func TestHeadFilter_Match_AllowsUpToMax(t *testing.T) {
	f, _ := filter.NewHeadFilter(3)
	lines := []string{"a", "b", "c", "d", "e"}
	expected := []bool{true, true, true, false, false}
	for i, line := range lines {
		got := f.Match(line)
		if got != expected[i] {
			t.Errorf("line %d %q: expected %v, got %v", i, line, expected[i], got)
		}
	}
}

func TestHeadFilter_Match_ExactlyOne(t *testing.T) {
	f, _ := filter.NewHeadFilter(1)
	if !f.Match("first") {
		t.Error("expected first line to match")
	}
	if f.Match("second") {
		t.Error("expected second line to not match")
	}
	if f.Seen() != 1 {
		t.Errorf("expected Seen=1, got %d", f.Seen())
	}
}

func TestHeadFilter_Seen_Counter(t *testing.T) {
	f, _ := filter.NewHeadFilter(5)
	for i := 0; i < 3; i++ {
		f.Match("line")
	}
	if f.Seen() != 3 {
		t.Errorf("expected Seen=3, got %d", f.Seen())
	}
}

func TestHeadFilter_InChain(t *testing.T) {
	head, _ := filter.NewHeadFilter(2)
	chain := filter.NewChain(head)
	if !chain.Match("line1") {
		t.Error("expected line1 to match")
	}
	if !chain.Match("line2") {
		t.Error("expected line2 to match")
	}
	if chain.Match("line3") {
		t.Error("expected line3 to not match")
	}
}
