package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewTemplateFilter_EmptyPattern(t *testing.T) {
	_, err := filter.NewTemplateFilter("", "{{.Line}}")
	if err == nil {
		t.Fatal("expected error for empty pattern")
	}
}

func TestNewTemplateFilter_EmptyTemplate(t *testing.T) {
	_, err := filter.NewTemplateFilter(`\d+`, "")
	if err == nil {
		t.Fatal("expected error for empty template")
	}
}

func TestNewTemplateFilter_InvalidPattern(t *testing.T) {
	_, err := filter.NewTemplateFilter(`[invalid`, "{{.Line}}")
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
}

func TestNewTemplateFilter_InvalidTemplate(t *testing.T) {
	_, err := filter.NewTemplateFilter(`\d+`, "{{.Unclosed")
	if err == nil {
		t.Fatal("expected error for invalid template")
	}
}

func TestNewTemplateFilter_Valid(t *testing.T) {
	f, err := filter.NewTemplateFilter(`(\w+)`, "MATCH:{{.Match}}")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Pattern() != `(\w+)` {
		t.Errorf("unexpected pattern: %s", f.Pattern())
	}
	if f.Template() != "MATCH:{{.Match}}" {
		t.Errorf("unexpected template: %s", f.Template())
	}
}

func TestTemplateFilter_Match_AlwaysTrue(t *testing.T) {
	f, _ := filter.NewTemplateFilter(`\d+`, "{{.Line}}")
	if !f.Match("no digits here") {
		t.Error("Match should always return true")
	}
	if !f.Match("123") {
		t.Error("Match should always return true")
	}
}

func TestTemplateFilter_Transform_NonMatchingPassthrough(t *testing.T) {
	f, _ := filter.NewTemplateFilter(`\d+`, "NUM:{{.Match}}")
	result := f.Transform("no numbers")
	if result != "no numbers" {
		t.Errorf("expected passthrough, got %q", result)
	}
}

func TestTemplateFilter_Transform_MatchingLine(t *testing.T) {
	f, _ := filter.NewTemplateFilter(`(\d+)`, "found:{{.Match}}")
	result := f.Transform("error 404 not found")
	if result != "found:404" {
		t.Errorf("unexpected result: %q", result)
	}
}

func TestTemplateFilter_Transform_FullLine(t *testing.T) {
	f, _ := filter.NewTemplateFilter(`ERROR`, "[ALERT] {{.Line}}")
	result := f.Transform("ERROR: disk full")
	if result != "[ALERT] ERROR: disk full" {
		t.Errorf("unexpected result: %q", result)
	}
}

func TestTemplateFilter_Transform_MultipleSubmatches(t *testing.T) {
	f, _ := filter.NewTemplateFilter(`(\w+)=(\w+)`, `{{index .Matches 0}}->{{index .Matches 1}}`)
	result := f.Transform("key=value")
	if result != "key->value" {
		t.Errorf("unexpected result: %q", result)
	}
}
