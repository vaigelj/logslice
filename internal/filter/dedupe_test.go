package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewDedupeFilter_InitialSeenZero(t *testing.T) {
	f := filter.NewDedupeFilter()
	if got := f.Seen(); got != 0 {
		t.Fatalf("expected 0 seen, got %d", got)
	}
}

func TestDedupeFilter_FirstOccurrenceMatches(t *testing.T) {
	f := filter.NewDedupeFilter()
	if !f.Match("hello world") {
		t.Fatal("expected first occurrence to match")
	}
}

func TestDedupeFilter_DuplicateRejected(t *testing.T) {
	f := filter.NewDedupeFilter()
	f.Match("hello world")
	if f.Match("hello world") {
		t.Fatal("expected duplicate to be rejected")
	}
}

func TestDedupeFilter_DifferentLinesAllMatch(t *testing.T) {
	f := filter.NewDedupeFilter()
	lines := []string{"alpha", "beta", "gamma"}
	for _, l := range lines {
		if !f.Match(l) {
			t.Fatalf("expected %q to match on first occurrence", l)
		}
	}
	if got := f.Seen(); got != 3 {
		t.Fatalf("expected 3 seen, got %d", got)
	}
}

func TestDedupeFilter_SeenCountAccurate(t *testing.T) {
	f := filter.NewDedupeFilter()
	f.Match("line1")
	f.Match("line1")
	f.Match("line2")
	f.Match("line2")
	f.Match("line3")
	if got := f.Seen(); got != 3 {
		t.Fatalf("expected 3 unique lines, got %d", got)
	}
}

func TestDedupeFilter_InChain(t *testing.T) {
	dedup := filter.NewDedupeFilter()
	chain, err := filter.NewChain(dedup)
	if err != nil {
		t.Fatalf("unexpected error building chain: %v", err)
	}
	if !chain.Match("unique") {
		t.Fatal("expected unique line to pass chain")
	}
	if chain.Match("unique") {
		t.Fatal("expected duplicate to be rejected by chain")
	}
}
