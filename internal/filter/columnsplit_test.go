package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewColumnSplitFilter_EmptyDelimiter(t *testing.T) {
	_, err := filter.NewColumnSplitFilter("", 0)
	if err == nil {
		t.Fatal("expected error for empty delimiter")
	}
}

func TestNewColumnSplitFilter_NegativeColumn(t *testing.T) {
	_, err := filter.NewColumnSplitFilter(",", -1)
	if err == nil {
		t.Fatal("expected error for negative column")
	}
}

func TestNewColumnSplitFilter_Valid(t *testing.T) {
	f, err := filter.NewColumnSplitFilter(",", 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Delimiter() != "," {
		t.Errorf("expected delimiter ',', got %q", f.Delimiter())
	}
	if f.Column() != 2 {
		t.Errorf("expected column 2, got %d", f.Column())
	}
}

func TestColumnSplitFilter_Match_ValidColumn(t *testing.T) {
	f, _ := filter.NewColumnSplitFilter(",", 1)
	line := "alpha,beta,gamma"
	if !f.Match(&line) {
		t.Fatal("expected Match to return true")
	}
	if line != "beta" {
		t.Errorf("expected 'beta', got %q", line)
	}
}

func TestColumnSplitFilter_Match_FirstColumn(t *testing.T) {
	f, _ := filter.NewColumnSplitFilter("|", 0)
	line := "one|two|three"
	if !f.Match(&line) {
		t.Fatal("expected Match to return true")
	}
	if line != "one" {
		t.Errorf("expected 'one', got %q", line)
	}
}

func TestColumnSplitFilter_Match_ColumnOutOfRange(t *testing.T) {
	f, _ := filter.NewColumnSplitFilter(",", 5)
	line := "a,b,c"
	if f.Match(&line) {
		t.Fatal("expected Match to return false for out-of-range column")
	}
}

func TestColumnSplitFilter_Match_NoDelimiterInLine(t *testing.T) {
	f, _ := filter.NewColumnSplitFilter(",", 1)
	line := "nodelmiter"
	if f.Match(&line) {
		t.Fatal("expected Match to return false when delimiter absent and column > 0")
	}
}

func TestColumnSplitFilter_InChain(t *testing.T) {
	col, _ := filter.NewColumnSplitFilter(" ", 2)
	chain, _ := filter.NewChain(col)
	line := "2024-01-01 ERROR something failed"
	if !chain.Match(&line) {
		t.Fatal("expected chain to match")
	}
	if line != "something" {
		t.Errorf("expected 'something', got %q", line)
	}
}
