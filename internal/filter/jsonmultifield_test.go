package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewJSONMultiFieldFilter_NoPairs(t *testing.T) {
	_, err := filter.NewJSONMultiFieldFilter(nil, false)
	if err == nil {
		t.Fatal("expected error for empty pairs")
	}
}

func TestNewJSONMultiFieldFilter_InvalidPairFormat(t *testing.T) {
	_, err := filter.NewJSONMultiFieldFilter([]string{"noequals"}, false)
	if err == nil {
		t.Fatal("expected error for pair without '='")
	}
}

func TestNewJSONMultiFieldFilter_EmptyValue(t *testing.T) {
	_, err := filter.NewJSONMultiFieldFilter([]string{"field="}, false)
	if err == nil {
		t.Fatal("expected error for empty value")
	}
}

func TestNewJSONMultiFieldFilter_Valid(t *testing.T) {
	f, err := filter.NewJSONMultiFieldFilter([]string{"level=error", "service=api"}, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f == nil {
		t.Fatal("expected non-nil filter")
	}
}

func TestJSONMultiFieldFilter_Match_AllMatch(t *testing.T) {
	f, _ := filter.NewJSONMultiFieldFilter([]string{"level=error", "service=api"}, false)
	line := `{"level":"error","service":"api","msg":"oops"}`
	if !f.Match(line) {
		t.Error("expected match when all fields match")
	}
}

func TestJSONMultiFieldFilter_Match_OneMismatch(t *testing.T) {
	f, _ := filter.NewJSONMultiFieldFilter([]string{"level=error", "service=api"}, false)
	line := `{"level":"error","service":"worker"}`
	if f.Match(line) {
		t.Error("expected no match when one field differs")
	}
}

func TestJSONMultiFieldFilter_Match_MissingField(t *testing.T) {
	f, _ := filter.NewJSONMultiFieldFilter([]string{"level=error", "service=api"}, false)
	line := `{"level":"error"}`
	if f.Match(line) {
		t.Error("expected no match when a field is absent")
	}
}

func TestJSONMultiFieldFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := filter.NewJSONMultiFieldFilter([]string{"level=ERROR"}, true)
	line := `{"level":"error"}`
	if !f.Match(line) {
		t.Error("expected case-insensitive match")
	}
}

func TestJSONMultiFieldFilter_Match_InvalidJSON(t *testing.T) {
	f, _ := filter.NewJSONMultiFieldFilter([]string{"level=error"}, false)
	if f.Match("not json") {
		t.Error("expected no match for invalid JSON")
	}
}

func TestJSONMultiFieldFilter_Transform_Passthrough(t *testing.T) {
	f, _ := filter.NewJSONMultiFieldFilter([]string{"k=v"}, false)
	line := `{"k":"v"}`
	if got := f.Transform(line); got != line {
		t.Errorf("expected passthrough, got %q", got)
	}
}
