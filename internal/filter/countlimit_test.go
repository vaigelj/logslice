package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewCountLimitFilter_InvalidMax(t *testing.T) {
	_, err := filter.NewCountLimitFilter(0)
	if err == nil {
		t.Fatal("expected error for maxCount=0, got nil")
	}
	_, err = filter.NewCountLimitFilter(-5)
	if err == nil {
		t.Fatal("expected error for maxCount=-5, got nil")
	}
}

func TestNewCountLimitFilter_Valid(t *testing.T) {
	f, err := filter.NewCountLimitFilter(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.MaxCount() != 3 {
		t.Errorf("expected MaxCount=3, got %d", f.MaxCount())
	}
}

func TestCountLimitFilter_Match_AllowsUpToMax(t *testing.T) {
	f, _ := filter.NewCountLimitFilter(3)
	lines := []string{"a", "b", "c", "d", "e"}
	expected := []bool{true, true, true, false, false}
	for i, line := range lines {
		got := f.Match(line)
		if got != expected[i] {
			t.Errorf("line %d: expected Match=%v, got %v", i, expected[i], got)
		}
	}
}

func TestCountLimitFilter_Match_ExactlyOne(t *testing.T) {
	f, _ := filter.NewCountLimitFilter(1)
	if !f.Match("first") {
		t.Error("expected first line to match")
	}
	if f.Match("second") {
		t.Error("expected second line to not match")
	}
}

func TestCountLimitFilter_Matched_Counter(t *testing.T) {
	f, _ := filter.NewCountLimitFilter(5)
	for i := 0; i < 3; i++ {
		f.Match("line")
	}
	if f.Matched() != 3 {
		t.Errorf("expected Matched=3, got %d", f.Matched())
	}
}

func TestCountLimitFilter_InChain(t *testing.T) {
	regex, _ := filter.NewRegexFilter("error")
	limit, _ := filter.NewCountLimitFilter(2)
	chain, _ := filter.NewChain(regex, limit)

	matches := 0
	for _, line := range []string{"error one", "error two", "error three", "info msg"} {
		if chain.Match(line) {
			matches++
		}
	}
	if matches != 2 {
		t.Errorf("expected 2 matches via chain, got %d", matches)
	}
}
