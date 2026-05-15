package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewJSONLevelFilter_EmptyField(t *testing.T) {
	_, err := filter.NewJSONLevelFilter("", []string{"error"}, false)
	if err == nil {
		t.Fatal("expected error for empty field")
	}
}

func TestNewJSONLevelFilter_EmptyLevels(t *testing.T) {
	_, err := filter.NewJSONLevelFilter("level", nil, false)
	if err == nil {
		t.Fatal("expected error for empty levels")
	}
}

func TestNewJSONLevelFilter_Valid(t *testing.T) {
	f, err := filter.NewJSONLevelFilter("level", []string{"error", "warn"}, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Field() != "level" {
		t.Errorf("expected field 'level', got %q", f.Field())
	}
	if f.CaseInsensitive() {
		t.Error("expected case-sensitive")
	}
}

func TestJSONLevelFilter_Match_CaseSensitive(t *testing.T) {
	f, _ := filter.NewJSONLevelFilter("level", []string{"error", "warn"}, false)
	tests := []struct {
		line string
		want bool
	}{
		{`{"level":"error","msg":"oops"}`, true},
		{`{"level":"warn","msg":"careful"}`, true},
		{`{"level":"info","msg":"ok"}`, false},
		{`{"level":"ERROR","msg":"oops"}`, false},
		{`{"msg":"no level field"}`, false},
		{`not json`, false},
	}
	for _, tc := range tests {
		if got := f.Match(tc.line); got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestJSONLevelFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := filter.NewJSONLevelFilter("level", []string{"error", "warn"}, true)
	tests := []struct {
		line string
		want bool
	}{
		{`{"level":"ERROR"}`, true},
		{`{"level":"Warn"}`, true},
		{`{"level":"INFO"}`, false},
	}
	for _, tc := range tests {
		if got := f.Match(tc.line); got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestJSONLevelFilter_Match_NonStringField(t *testing.T) {
	f, _ := filter.NewJSONLevelFilter("level", []string{"1"}, false)
	if f.Match(`{"level":1}`) {
		t.Error("expected false for non-string JSON field")
	}
}

func TestJSONLevelFilter_InChain(t *testing.T) {
	f, _ := filter.NewJSONLevelFilter("severity", []string{"critical"}, false)
	chain, _ := filter.NewChain(f)
	if !chain.Match(`{"severity":"critical","msg":"down"}`) {
		t.Error("expected chain match for critical severity")
	}
	if chain.Match(`{"severity":"debug","msg":"trace"}`) {
		t.Error("expected chain reject for debug severity")
	}
}
