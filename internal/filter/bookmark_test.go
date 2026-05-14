package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewBookmarkFilter_EmptyStart(t *testing.T) {
	_, err := filter.NewBookmarkFilter("", "END")
	if err == nil {
		t.Fatal("expected error for empty start pattern")
	}
}

func TestNewBookmarkFilter_EmptyEnd(t *testing.T) {
	_, err := filter.NewBookmarkFilter("START", "")
	if err == nil {
		t.Fatal("expected error for empty end pattern")
	}
}

func TestNewBookmarkFilter_InvalidStartRegex(t *testing.T) {
	_, err := filter.NewBookmarkFilter("[", "END")
	if err == nil {
		t.Fatal("expected error for invalid start regex")
	}
}

func TestNewBookmarkFilter_InvalidEndRegex(t *testing.T) {
	_, err := filter.NewBookmarkFilter("START", "[")
	if err == nil {
		t.Fatal("expected error for invalid end regex")
	}
}

func TestNewBookmarkFilter_Valid(t *testing.T) {
	f, err := filter.NewBookmarkFilter("START", "END")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.StartPattern() != "START" {
		t.Errorf("expected StartPattern START, got %s", f.StartPattern())
	}
	if f.EndPattern() != "END" {
		t.Errorf("expected EndPattern END, got %s", f.EndPattern())
	}
}

func TestBookmarkFilter_Match_RegionIncluded(t *testing.T) {
	f, _ := filter.NewBookmarkFilter("BEGIN", "FINISH")

	lines := []struct {
		line string
		want bool
	}{
		{"before", false},
		{"BEGIN here", true},
		{"inside", true},
		{"FINISH now", true},
		{"after", false},
	}

	for _, tc := range lines {
		got := f.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestBookmarkFilter_Match_MultipleRegions(t *testing.T) {
	f, _ := filter.NewBookmarkFilter("START", "END")

	results := []bool{
		f.Match("START"),  // true
		f.Match("middle"), // true
		f.Match("END"),    // true
		f.Match("gap"),    // false
		f.Match("START"),  // true
		f.Match("END"),    // true
	}

	expected := []bool{true, true, true, false, true, true}
	for i, got := range results {
		if got != expected[i] {
			t.Errorf("step %d: got %v, want %v", i, got, expected[i])
		}
	}
}

func TestBookmarkFilter_Matched_Counter(t *testing.T) {
	f, _ := filter.NewBookmarkFilter("START", "END")
	f.Match("outside")
	f.Match("START")
	f.Match("inside")
	f.Match("END")
	f.Match("outside again")

	if f.Matched() != 3 {
		t.Errorf("expected Matched()=3, got %d", f.Matched())
	}
}
