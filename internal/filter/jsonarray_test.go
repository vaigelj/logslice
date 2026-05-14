package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewJSONArrayFilter_EmptyField(t *testing.T) {
	_, err := filter.NewJSONArrayFilter("", "val", false)
	if err == nil {
		t.Fatal("expected error for empty field")
	}
}

func TestNewJSONArrayFilter_EmptyValue(t *testing.T) {
	_, err := filter.NewJSONArrayFilter("tags", "", false)
	if err == nil {
		t.Fatal("expected error for empty value")
	}
}

func TestNewJSONArrayFilter_Valid(t *testing.T) {
	f, err := filter.NewJSONArrayFilter("tags", "error", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Field() != "tags" {
		t.Errorf("Field() = %q, want %q", f.Field(), "tags")
	}
	if f.Value() != "error" {
		t.Errorf("Value() = %q, want %q", f.Value(), "error")
	}
	if f.IgnoreCase() {
		t.Error("IgnoreCase() should be false")
	}
}

func TestJSONArrayFilter_Match_Found(t *testing.T) {
	f, _ := filter.NewJSONArrayFilter("tags", "error", false)
	line := `{"tags":["info","error","warn"],"msg":"test"}`
	if !f.Match(line) {
		t.Error("expected match")
	}
}

func TestJSONArrayFilter_Match_NotFound(t *testing.T) {
	f, _ := filter.NewJSONArrayFilter("tags", "debug", false)
	line := `{"tags":["info","error"],"msg":"test"}`
	if f.Match(line) {
		t.Error("expected no match")
	}
}

func TestJSONArrayFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := filter.NewJSONArrayFilter("tags", "ERROR", true)
	line := `{"tags":["info","error"],"msg":"test"}`
	if !f.Match(line) {
		t.Error("expected case-insensitive match")
	}
}

func TestJSONArrayFilter_Match_MissingField(t *testing.T) {
	f, _ := filter.NewJSONArrayFilter("labels", "error", false)
	line := `{"tags":["error"],"msg":"test"}`
	if f.Match(line) {
		t.Error("expected no match when field is absent")
	}
}

func TestJSONArrayFilter_Match_InvalidJSON(t *testing.T) {
	f, _ := filter.NewJSONArrayFilter("tags", "error", false)
	if f.Match("not json at all") {
		t.Error("expected no match for invalid JSON")
	}
}

func TestJSONArrayFilter_Match_FieldNotArray(t *testing.T) {
	f, _ := filter.NewJSONArrayFilter("tags", "error", false)
	line := `{"tags":"error"}`
	if f.Match(line) {
		t.Error("expected no match when field is not an array")
	}
}

func TestJSONArrayFilter_InChain(t *testing.T) {
	f, _ := filter.NewJSONArrayFilter("tags", "warn", false)
	chain, _ := filter.NewChain(f)
	if !chain.Match(`{"tags":["warn","info"]}`) {
		t.Error("expected chain match")
	}
	if chain.Match(`{"tags":["info"]}`) {
		t.Error("expected chain no-match")
	}
}
