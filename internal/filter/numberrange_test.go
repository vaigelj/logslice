package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewNumberRangeFilter_EmptyDelimiter(t *testing.T) {
	_, err := filter.NewNumberRangeFilter("", 0, 0, 100)
	if err == nil {
		t.Fatal("expected error for empty delimiter")
	}
}

func TestNewNumberRangeFilter_NegativeIndex(t *testing.T) {
	_, err := filter.NewNumberRangeFilter(",", -1, 0, 100)
	if err == nil {
		t.Fatal("expected error for negative field index")
	}
}

func TestNewNumberRangeFilter_MinExceedsMax(t *testing.T) {
	_, err := filter.NewNumberRangeFilter(",", 0, 50, 10)
	if err == nil {
		t.Fatal("expected error when min > max")
	}
}

func TestNewNumberRangeFilter_Valid(t *testing.T) {
	f, err := filter.NewNumberRangeFilter(",", 2, 1.5, 9.5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Delimiter() != "," {
		t.Errorf("expected delimiter ',', got %q", f.Delimiter())
	}
	if f.FieldIndex() != 2 {
		t.Errorf("expected field index 2, got %d", f.FieldIndex())
	}
	if f.Min() != 1.5 || f.Max() != 9.5 {
		t.Errorf("unexpected min/max: %g/%g", f.Min(), f.Max())
	}
}

func TestNumberRangeFilter_Match_InRange(t *testing.T) {
	f, _ := filter.NewNumberRangeFilter(",", 1, 10, 20)
	if !f.Match("foo,15,bar") {
		t.Error("expected match for value 15 in [10,20]")
	}
}

func TestNumberRangeFilter_Match_BoundaryValues(t *testing.T) {
	f, _ := filter.NewNumberRangeFilter(",", 0, 5, 10)
	if !f.Match("5") {
		t.Error("expected match at lower boundary")
	}
	if !f.Match("10") {
		t.Error("expected match at upper boundary")
	}
}

func TestNumberRangeFilter_Match_OutOfRange(t *testing.T) {
	f, _ := filter.NewNumberRangeFilter(",", 0, 5, 10)
	if f.Match("4.9") {
		t.Error("expected no match for value below range")
	}
	if f.Match("10.1") {
		t.Error("expected no match for value above range")
	}
}

func TestNumberRangeFilter_Match_FieldMissing(t *testing.T) {
	f, _ := filter.NewNumberRangeFilter(",", 5, 0, 100)
	if f.Match("a,b,c") {
		t.Error("expected no match when field index out of bounds")
	}
}

func TestNumberRangeFilter_Match_NonNumeric(t *testing.T) {
	f, _ := filter.NewNumberRangeFilter(",", 1, 0, 100)
	if f.Match("foo,bar,baz") {
		t.Error("expected no match for non-numeric field")
	}
}

func TestNumberRangeFilter_Transform_Unchanged(t *testing.T) {
	f, _ := filter.NewNumberRangeFilter(",", 0, 0, 100)
	line := "42,hello"
	if got := f.Transform(line); got != line {
		t.Errorf("expected Transform to return line unchanged, got %q", got)
	}
}
