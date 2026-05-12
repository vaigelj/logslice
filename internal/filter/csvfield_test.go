package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewCSVFieldFilter_NegativeIndex(t *testing.T) {
	_, err := filter.NewCSVFieldFilter(-1, "foo", false)
	if err == nil {
		t.Fatal("expected error for negative index")
	}
}

func TestNewCSVFieldFilter_EmptyExpected(t *testing.T) {
	_, err := filter.NewCSVFieldFilter(0, "", false)
	if err == nil {
		t.Fatal("expected error for empty expected value")
	}
}

func TestNewCSVFieldFilter_Valid(t *testing.T) {
	f, err := filter.NewCSVFieldFilter(1, "bar", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Index() != 1 {
		t.Errorf("expected index 1, got %d", f.Index())
	}
	if f.Expected() != "bar" {
		t.Errorf("expected 'bar', got %q", f.Expected())
	}
}

func TestCSVFieldFilter_Match_CaseSensitive(t *testing.T) {
	f, _ := filter.NewCSVFieldFilter(1, "hello", false)
	if !f.Match("foo,hello,baz") {
		t.Error("expected match")
	}
	if f.Match("foo,HELLO,baz") {
		t.Error("expected no match (case sensitive)")
	}
}

func TestCSVFieldFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := filter.NewCSVFieldFilter(0, "ERROR", true)
	if !f.Match("error,some message") {
		t.Error("expected case-insensitive match")
	}
}

func TestCSVFieldFilter_Match_IndexOutOfBounds(t *testing.T) {
	f, _ := filter.NewCSVFieldFilter(5, "x", false)
	if f.Match("a,b,c") {
		t.Error("expected no match when index out of bounds")
	}
}

func TestCSVFieldFilter_Match_QuotedField(t *testing.T) {
	f, _ := filter.NewCSVFieldFilter(1, "hello world", false)
	if !f.Match(`foo,"hello world",baz`) {
		t.Error("expected match for quoted CSV field")
	}
}

func TestCSVFieldFilter_Transform_ReturnsLine(t *testing.T) {
	f, _ := filter.NewCSVFieldFilter(0, "x", false)
	line := "x,y,z"
	if got := f.Transform(line); got != line {
		t.Errorf("expected %q, got %q", line, got)
	}
}

func TestCSVFieldFilter_InChain(t *testing.T) {
	f1, _ := filter.NewCSVFieldFilter(0, "INFO", true)
	f2, _ := filter.NewCSVFieldFilter(2, "db", false)
	chain, _ := filter.NewChain(f1, f2)
	if !chain.Match("info,2024-01-01,db") {
		t.Error("expected chain match")
	}
	if chain.Match("info,2024-01-01,web") {
		t.Error("expected chain no-match on second filter")
	}
}
