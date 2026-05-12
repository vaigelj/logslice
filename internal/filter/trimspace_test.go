package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewTrimSpaceFilter_BothFalse(t *testing.T) {
	_, err := filter.NewTrimSpaceFilter(false, false)
	if err == nil {
		t.Fatal("expected error when both trimLeft and trimRight are false")
	}
}

func TestNewTrimSpaceFilter_TrimLeft(t *testing.T) {
	f, err := filter.NewTrimSpaceFilter(true, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !f.TrimLeft() {
		t.Error("expected TrimLeft to be true")
	}
	if f.TrimRight() {
		t.Error("expected TrimRight to be false")
	}
}

func TestNewTrimSpaceFilter_TrimRight(t *testing.T) {
	f, err := filter.NewTrimSpaceFilter(false, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.TrimLeft() {
		t.Error("expected TrimLeft to be false")
	}
	if !f.TrimRight() {
		t.Error("expected TrimRight to be true")
	}
}

func TestTrimSpaceFilter_Match_AlwaysTrue(t *testing.T) {
	f, _ := filter.NewTrimSpaceFilter(true, true)
	for _, line := range []string{"", "   ", "hello", "  world  "} {
		if !f.Match(line) {
			t.Errorf("Match(%q) = false, want true", line)
		}
	}
}

func TestTrimSpaceFilter_Transform_BothSides(t *testing.T) {
	f, _ := filter.NewTrimSpaceFilter(true, true)
	cases := []struct{ in, want string }{
		{"  hello  ", "hello"},
		{"\t foo \t", "foo"},
		{"no-spaces", "no-spaces"},
		{"   ", ""},
	}
	for _, c := range cases {
		if got := f.Transform(c.in); got != c.want {
			t.Errorf("Transform(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestTrimSpaceFilter_Transform_LeftOnly(t *testing.T) {
	f, _ := filter.NewTrimSpaceFilter(true, false)
	if got := f.Transform("  hello  "); got != "hello  " {
		t.Errorf("got %q, want %q", got, "hello  ")
	}
}

func TestTrimSpaceFilter_Transform_RightOnly(t *testing.T) {
	f, _ := filter.NewTrimSpaceFilter(false, true)
	if got := f.Transform("  hello  "); got != "  hello" {
		t.Errorf("got %q, want %q", got, "  hello")
	}
}

func TestTrimSpaceFilter_InChain(t *testing.T) {
	trim, _ := filter.NewTrimSpaceFilter(true, true)
	regex, _ := filter.NewRegexFilter("^hello")
	chain, _ := filter.NewChain(trim, regex)

	if !chain.Match("  hello world") {
		t.Error("expected chain to match after trimming")
	}
	if chain.Match("  world hello") {
		t.Error("expected chain to reject non-matching line")
	}
}
