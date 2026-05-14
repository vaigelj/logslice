package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewRegexReplaceFilter_EmptyPattern(t *testing.T) {
	_, err := filter.NewRegexReplaceFilter("", "x")
	if err == nil {
		t.Fatal("expected error for empty pattern")
	}
}

func TestNewRegexReplaceFilter_InvalidPattern(t *testing.T) {
	_, err := filter.NewRegexReplaceFilter("[", "x")
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
}

func TestNewRegexReplaceFilter_Valid(t *testing.T) {
	f, err := filter.NewRegexReplaceFilter(`\d+`, "NUM")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Pattern() != `\d+` {
		t.Errorf("unexpected pattern: %s", f.Pattern())
	}
	if f.Replacement() != "NUM" {
		t.Errorf("unexpected replacement: %s", f.Replacement())
	}
}

func TestRegexReplaceFilter_Match_AlwaysTrue(t *testing.T) {
	f, _ := filter.NewRegexReplaceFilter(`\d+`, "NUM")
	for _, line := range []string{"", "abc", "123", "no digits here"} {
		if !f.Match(line) {
			t.Errorf("expected Match to return true for %q", line)
		}
	}
}

func TestRegexReplaceFilter_Transform_ReplacesAll(t *testing.T) {
	f, _ := filter.NewRegexReplaceFilter(`\d+`, "NUM")
	got := f.Transform("error 404 at line 12")
	want := "error NUM at line NUM"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRegexReplaceFilter_Transform_NoMatch(t *testing.T) {
	f, _ := filter.NewRegexReplaceFilter(`\d+`, "NUM")
	got := f.Transform("no numbers here")
	if got != "no numbers here" {
		t.Errorf("unexpected transform: %q", got)
	}
}

func TestRegexReplaceFilter_Transform_EmptyReplacement(t *testing.T) {
	f, _ := filter.NewRegexReplaceFilter(`\s+`, "")
	got := f.Transform("hello world foo")
	if got != "helloworldfoo" {
		t.Errorf("got %q, want %q", got, "helloworldfoo")
	}
}

func TestRegexReplaceFilter_Transform_CaptureGroup(t *testing.T) {
	f, _ := filter.NewRegexReplaceFilter(`(\w+)@(\w+)`, "[$1 at $2]")
	got := f.Transform("user@host logged in")
	want := "[user at host] logged in"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRegexReplaceFilter_InChain(t *testing.T) {
	rf, _ := filter.NewRegexReplaceFilter(`ERROR`, "WARN")
	chain, err := filter.NewChain(rf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !chain.Match("some ERROR line") {
		t.Error("chain should pass all lines")
	}
}
