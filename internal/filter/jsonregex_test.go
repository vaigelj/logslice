package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewJSONRegexFilter_EmptyField(t *testing.T) {
	_, err := filter.NewJSONRegexFilter("", `\d+`)
	if err == nil {
		t.Fatal("expected error for empty field")
	}
}

func TestNewJSONRegexFilter_EmptyPattern(t *testing.T) {
	_, err := filter.NewJSONRegexFilter("level", "")
	if err == nil {
		t.Fatal("expected error for empty pattern")
	}
}

func TestNewJSONRegexFilter_InvalidPattern(t *testing.T) {
	_, err := filter.NewJSONRegexFilter("level", "[invalid")
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
}

func TestNewJSONRegexFilter_Valid(t *testing.T) {
	f, err := filter.NewJSONRegexFilter("level", `^(error|warn)$`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Field() != "level" {
		t.Errorf("expected field 'level', got %q", f.Field())
	}
	if f.Pattern() != `^(error|warn)$` {
		t.Errorf("unexpected pattern: %q", f.Pattern())
	}
}

func TestJSONRegexFilter_Match_StringField(t *testing.T) {
	f, _ := filter.NewJSONRegexFilter("level", `^(error|warn)$`)

	tests := []struct {
		line  string
		want  bool
	}{
		{`{"level":"error","msg":"oops"}`, true},
		{`{"level":"warn","msg":"careful"}`, true},
		{`{"level":"info","msg":"ok"}`, false},
		{`{"level":"debug","msg":"trace"}`, false},
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

func TestJSONRegexFilter_Match_NumericField(t *testing.T) {
	f, _ := filter.NewJSONRegexFilter("code", `^[45]`)

	tests := []struct {
		line string
		want bool
	}{
		{`{"code":404}`, true},
		{`{"code":500}`, true},
		{`{"code":200}`, false},
		{`{"code":301}`, false},
	}

	for _, tc := range tests {
		got := f.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}
