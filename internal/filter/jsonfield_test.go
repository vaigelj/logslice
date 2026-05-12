package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewJSONFieldFilter_EmptyField(t *testing.T) {
	_, err := filter.NewJSONFieldFilter("", "value", false)
	if err == nil {
		t.Fatal("expected error for empty field, got nil")
	}
}

func TestNewJSONFieldFilter_EmptyExpected(t *testing.T) {
	_, err := filter.NewJSONFieldFilter("level", "", false)
	if err == nil {
		t.Fatal("expected error for empty expected value, got nil")
	}
}

func TestNewJSONFieldFilter_Valid(t *testing.T) {
	f, err := filter.NewJSONFieldFilter("level", "error", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Field() != "level" {
		t.Errorf("Field() = %q, want %q", f.Field(), "level")
	}
	if f.Expected() != "error" {
		t.Errorf("Expected() = %q, want %q", f.Expected(), "error")
	}
	if f.CaseInsensitive() {
		t.Error("CaseInsensitive() should be false")
	}
}

func TestJSONFieldFilter_Match_CaseSensitive(t *testing.T) {
	f, _ := filter.NewJSONFieldFilter("level", "error", false)

	tests := []struct {
		line string
		want bool
	}{
		{`{"level":"error","msg":"oops"}`, true},
		{`{"level":"ERROR","msg":"oops"}`, false},
		{`{"level":"info","msg":"ok"}`, false},
		{`{"msg":"no level field"}`, false},
		{`not json at all`, false},
		{`{}`, false},
	}
	for _, tc := range tests {
		got := f.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestJSONFieldFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := filter.NewJSONFieldFilter("level", "error", true)

	tests := []struct {
		line string
		want bool
	}{
		{`{"level":"error"}`, true},
		{`{"level":"ERROR"}`, true},
		{`{"level":"Error"}`, true},
		{`{"level":"info"}`, false},
	}
	for _, tc := range tests {
		got := f.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestJSONFieldFilter_Transform_ReturnsUnchanged(t *testing.T) {
	f, _ := filter.NewJSONFieldFilter("level", "error", false)
	line := `{"level":"error"}`
	if got := f.Transform(line); got != line {
		t.Errorf("Transform() = %q, want %q", got, line)
	}
}

func TestJSONFieldFilter_Match_NumericField(t *testing.T) {
	f, _ := filter.NewJSONFieldFilter("code", "404", false)
	if !f.Match(`{"code":404,"msg":"not found"}`) {
		t.Error("expected match for numeric JSON field")
	}
	if f.Match(`{"code":200,"msg":"ok"}`) {
		t.Error("expected no match for different numeric value")
	}
}
