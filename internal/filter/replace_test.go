package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewReplaceFilter_EmptyPattern(t *testing.T) {
	_, err := filter.NewReplaceFilter("", "x", false, false)
	if err == nil {
		t.Fatal("expected error for empty pattern")
	}
}

func TestNewReplaceFilter_InvalidRegex(t *testing.T) {
	_, err := filter.NewReplaceFilter("[", "x", false, true)
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
}

func TestNewReplaceFilter_Valid(t *testing.T) {
	f, err := filter.NewReplaceFilter("foo", "bar", false, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Pattern() != "foo" {
		t.Errorf("expected pattern 'foo', got %q", f.Pattern())
	}
	if f.Replacement() != "bar" {
		t.Errorf("expected replacement 'bar', got %q", f.Replacement())
	}
}

func TestReplaceFilter_Match_AlwaysTrue(t *testing.T) {
	f, _ := filter.NewReplaceFilter("x", "y", false, false)
	for _, line := range []string{"", "hello", "xyz"} {
		if !f.Match(line) {
			t.Errorf("Match(%q) = false, want true", line)
		}
	}
}

func TestReplaceFilter_Transform_Literal(t *testing.T) {
	f, _ := filter.NewReplaceFilter("foo", "bar", false, false)
	got := f.Transform("foo and foo")
	if got != "bar and bar" {
		t.Errorf("got %q, want %q", got, "bar and bar")
	}
}

func TestReplaceFilter_Transform_CaseInsensitiveLiteral(t *testing.T) {
	f, _ := filter.NewReplaceFilter("foo", "bar", true, false)
	got := f.Transform("FOO and Foo")
	if got != "bar and bar" {
		t.Errorf("got %q, want %q", got, "bar and bar")
	}
}

func TestReplaceFilter_Transform_Regex(t *testing.T) {
	f, _ := filter.NewReplaceFilter(`\d+`, "NUM", false, true)
	got := f.Transform("error 404 and 500")
	if got != "error NUM and NUM" {
		t.Errorf("got %q, want %q", got, "error NUM and NUM")
	}
}

func TestReplaceFilter_Transform_RegexCaseInsensitive(t *testing.T) {
	f, _ := filter.NewReplaceFilter("error", "ERR", true, true)
	got := f.Transform("ERROR and Error")
	if got != "ERR and ERR" {
		t.Errorf("got %q, want %q", got, "ERR and ERR")
	}
}

func TestReplaceFilter_Transform_NoMatch(t *testing.T) {
	f, _ := filter.NewReplaceFilter("foo", "bar", false, false)
	got := f.Transform("nothing here")
	if got != "nothing here" {
		t.Errorf("got %q, want unchanged", got)
	}
}
