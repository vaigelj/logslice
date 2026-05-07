package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewSampleFilter_InvalidN(t *testing.T) {
	for _, n := range []int{0, -1, -100} {
		_, err := filter.NewSampleFilter(n)
		if err == nil {
			t.Errorf("expected error for n=%d", n)
		}
	}
}

func TestNewSampleFilter_ValidN(t *testing.T) {
	f, err := filter.NewSampleFilter(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.N() != 3 {
		t.Errorf("expected N=3, got %d", f.N())
	}
}

func TestSampleFilter_EveryLine(t *testing.T) {
	f, _ := filter.NewSampleFilter(1)
	for i := 0; i < 5; i++ {
		if !f.Match("line") {
			t.Errorf("expected Match=true for n=1 at iteration %d", i)
		}
	}
}

func TestSampleFilter_EveryOtherLine(t *testing.T) {
	f, _ := filter.NewSampleFilter(2)
	results := make([]bool, 6)
	for i := range results {
		results[i] = f.Match("line")
	}
	// Expect: false, true, false, true, false, true
	expected := []bool{false, true, false, true, false, true}
	for i, got := range results {
		if got != expected[i] {
			t.Errorf("index %d: expected %v, got %v", i, expected[i], got)
		}
	}
}

func TestSampleFilter_EveryThirdLine(t *testing.T) {
	f, _ := filter.NewSampleFilter(3)
	matches := 0
	for i := 0; i < 9; i++ {
		if f.Match("line") {
			matches++
		}
	}
	if matches != 3 {
		t.Errorf("expected 3 matches in 9 lines with n=3, got %d", matches)
	}
}

func TestSampleFilter_InChain(t *testing.T) {
	regex, _ := filter.NewRegexFilter("error")
	sample, _ := filter.NewSampleFilter(2)
	chain, _ := filter.NewChain(regex, sample)

	lines := []string{
		"error one",
		"error two",
		"error three",
		"info message",
		"error four",
	}
	matches := 0
	for _, l := range lines {
		if chain.Match(l) {
			matches++
		}
	}
	// "error one", "error two", "error three", "error four" pass regex (4 lines)
	// sample n=2 passes every 2nd: "error two", "error four" => 2 matches
	if matches != 2 {
		t.Errorf("expected 2 matches, got %d", matches)
	}
}
