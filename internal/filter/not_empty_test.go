package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewNotEmptyFilter_DefaultFalse(t *testing.T) {
	f := filter.NewNotEmptyFilter(false)
	if f == nil {
		t.Fatal("expected non-nil filter")
	}
	if f.TrimWhitespace() {
		t.Error("expected TrimWhitespace to be false")
	}
}

func TestNewNotEmptyFilter_TrimTrue(t *testing.T) {
	f := filter.NewNotEmptyFilter(true)
	if !f.TrimWhitespace() {
		t.Error("expected TrimWhitespace to be true")
	}
}

func TestNotEmptyFilter_Match_NoTrim(t *testing.T) {
	f := filter.NewNotEmptyFilter(false)

	cases := []struct {
		line string
		want bool
	}{
		{"", false},
		{" ", true},  // whitespace only but not trimmed
		{"\t", true}, // tab only but not trimmed
		{"hello", true},
		{"  spaces  ", true},
	}

	for _, tc := range cases {
		got := f.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestNotEmptyFilter_Match_WithTrim(t *testing.T) {
	f := filter.NewNotEmptyFilter(true)

	cases := []struct {
		line string
		want bool
	}{
		{"", false},
		{" ", false},
		{"\t", false},
		{"\t   \t", false},
		{"hello", true},
		{"  spaces  ", true},
		{" x ", true},
	}

	for _, tc := range cases {
		got := f.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestNotEmptyFilter_InChain(t *testing.T) {
	chain := filter.NewChain(
		filter.NewNotEmptyFilter(true),
	)

	if chain.Match("") {
		t.Error("empty line should be rejected by chain")
	}
	if !chain.Match("data") {
		t.Error("non-empty line should pass chain")
	}
}
