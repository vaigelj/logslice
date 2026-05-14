package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewWordCountFilter_NegativeMin(t *testing.T) {
	_, err := filter.NewWordCountFilter(-1, 0)
	if err == nil {
		t.Fatal("expected error for negative min")
	}
}

func TestNewWordCountFilter_MaxLessThanMin(t *testing.T) {
	_, err := filter.NewWordCountFilter(5, 3)
	if err == nil {
		t.Fatal("expected error when max < min")
	}
}

func TestNewWordCountFilter_Valid(t *testing.T) {
	f, err := filter.NewWordCountFilter(2, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Min() != 2 || f.Max() != 5 {
		t.Fatalf("accessors: got min=%d max=%d", f.Min(), f.Max())
	}
}

func TestWordCountFilter_Match_InRange(t *testing.T) {
	f, _ := filter.NewWordCountFilter(2, 4)
	cases := []struct {
		line string
		want bool
	}{
		{"one", false},
		{"one two", true},
		{"one two three", true},
		{"one two three four", true},
		{"one two three four five", false},
		{"", false},
	}
	for _, c := range cases {
		if got := f.Match(c.line); got != c.want {
			t.Errorf("Match(%q) = %v, want %v", c.line, got, c.want)
		}
	}
}

func TestWordCountFilter_Match_UnboundedMax(t *testing.T) {
	f, _ := filter.NewWordCountFilter(1, 0)
	if !f.Match("a b c d e f g") {
		t.Error("expected match for unbounded max")
	}
	if f.Match("") {
		t.Error("empty line should not match min=1")
	}
}

func TestWordCountFilter_Transform_Passthrough(t *testing.T) {
	f, _ := filter.NewWordCountFilter(0, 0)
	line := "hello world"
	if got := f.Transform(line); got != line {
		t.Errorf("Transform changed line: got %q", got)
	}
}

func TestWordCountFilter_InChain(t *testing.T) {
	wc, _ := filter.NewWordCountFilter(2, 3)
	chain, _ := filter.NewChain(wc)
	if !chain.Match("foo bar") {
		t.Error("expected chain to match 2-word line")
	}
	if chain.Match("only") {
		t.Error("expected chain to reject 1-word line")
	}
}
