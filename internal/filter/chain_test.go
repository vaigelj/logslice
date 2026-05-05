package filter

import (
	"testing"
)

// alwaysFilter is a stub Filter for testing.
type alwaysFilter struct{ result bool }

func (a *alwaysFilter) Match(_ string) bool { return a.result }

func TestNewChain_Empty(t *testing.T) {
	c := NewChain()
	if c.Len() != 0 {
		t.Errorf("expected 0 filters, got %d", c.Len())
	}
	if !c.Match("any line") {
		t.Error("empty chain should match all lines")
	}
}

func TestNewChain_NilFiltersIgnored(t *testing.T) {
	c := NewChain(nil, nil)
	if c.Len() != 0 {
		t.Errorf("expected 0 active filters, got %d", c.Len())
	}
}

func TestChain_AllMatch(t *testing.T) {
	c := NewChain(&alwaysFilter{true}, &alwaysFilter{true})
	if !c.Match("some log line") {
		t.Error("expected chain to match when all filters pass")
	}
}

func TestChain_OneRejects(t *testing.T) {
	c := NewChain(&alwaysFilter{true}, &alwaysFilter{false})
	if c.Match("some log line") {
		t.Error("expected chain to reject when one filter fails")
	}
}

func TestChain_WithRealFilters(t *testing.T) {
	rf, err := NewRegexFilter("ERROR")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tf, err := NewTimeRangeFilter("2024-06-01T08:00:00Z", "2024-06-01T09:00:00Z", testLayout)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	c := NewChain(rf, tf)
	if c.Len() != 2 {
		t.Errorf("expected 2 filters, got %d", c.Len())
	}

	if !c.Match("2024-06-01T08:30:00 ERROR disk full") {
		t.Error("expected match for line within range containing ERROR")
	}
	if c.Match("2024-06-01T08:30:00 INFO disk ok") {
		t.Error("expected no match for INFO line")
	}
	if c.Match("2024-06-01T07:00:00 ERROR too early") {
		t.Error("expected no match for out-of-range ERROR line")
	}
}

func TestChain_Add(t *testing.T) {
	c := NewChain()
	c.Add(&alwaysFilter{true})
	if c.Len() != 1 {
		t.Errorf("expected 1 filter after Add, got %d", c.Len())
	}
	c.Add(nil)
	if c.Len() != 1 {
		t.Error("nil Add should not increase Len")
	}
}
