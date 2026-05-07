package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewTruncateFilter_InvalidMaxLen(t *testing.T) {
	_, err := filter.NewTruncateFilter(0, "")
	if err == nil {
		t.Fatal("expected error for maxLen=0")
	}
	_, err = filter.NewTruncateFilter(-5, "...")
	if err == nil {
		t.Fatal("expected error for negative maxLen")
	}
}

func TestNewTruncateFilter_SuffixTooLong(t *testing.T) {
	_, err := filter.NewTruncateFilter(2, "...")
	if err == nil {
		t.Fatal("expected error when suffix length exceeds maxLen")
	}
}

func TestNewTruncateFilter_Valid(t *testing.T) {
	f, err := filter.NewTruncateFilter(10, "...")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.MaxLen() != 10 {
		t.Errorf("expected MaxLen 10, got %d", f.MaxLen())
	}
	if f.Suffix() != "..." {
		t.Errorf("expected suffix '...', got %q", f.Suffix())
	}
}

func TestTruncateFilter_Transform_ShortLine(t *testing.T) {
	f, _ := filter.NewTruncateFilter(20, "...")
	out := f.Transform("hello world")
	if out != "hello world" {
		t.Errorf("expected unchanged line, got %q", out)
	}
}

func TestTruncateFilter_Transform_ExactLength(t *testing.T) {
	f, _ := filter.NewTruncateFilter(5, "...")
	out := f.Transform("hello")
	if out != "hello" {
		t.Errorf("expected unchanged line, got %q", out)
	}
}

func TestTruncateFilter_Transform_LongLine(t *testing.T) {
	f, _ := filter.NewTruncateFilter(10, "...")
	out := f.Transform("this is a very long log line")
	if out != "this is..." {
		t.Errorf("expected 'this is...', got %q", out)
	}
	if len(out) != 10 {
		t.Errorf("expected length 10, got %d", len(out))
	}
}

func TestTruncateFilter_Transform_NoSuffix(t *testing.T) {
	f, _ := filter.NewTruncateFilter(5, "")
	out := f.Transform("abcdefgh")
	if out != "abcde" {
		t.Errorf("expected 'abcde', got %q", out)
	}
}

func TestTruncateFilter_Match_AlwaysTrue(t *testing.T) {
	f, _ := filter.NewTruncateFilter(5, "...")
	lines := []string{"", "hi", "a very long line that exceeds the limit"}
	for _, l := range lines {
		if !f.Match(l) {
			t.Errorf("Match(%q) expected true", l)
		}
	}
}
