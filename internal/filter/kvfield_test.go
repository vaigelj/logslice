package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewKVFieldFilter_EmptyKey(t *testing.T) {
	_, err := filter.NewKVFieldFilter("", "=", "value", false)
	if err == nil {
		t.Fatal("expected error for empty key")
	}
}

func TestNewKVFieldFilter_EmptyExpected(t *testing.T) {
	_, err := filter.NewKVFieldFilter("level", "=", "", false)
	if err == nil {
		t.Fatal("expected error for empty expected")
	}
}

func TestNewKVFieldFilter_DefaultSep(t *testing.T) {
	f, err := filter.NewKVFieldFilter("level", "", "info", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Sep() != "=" {
		t.Errorf("expected default sep '=', got %q", f.Sep())
	}
}

func TestNewKVFieldFilter_Valid(t *testing.T) {
	f, err := filter.NewKVFieldFilter("status", "=", "200", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Key() != "status" {
		t.Errorf("expected key 'status', got %q", f.Key())
	}
	if f.Expected() != "200" {
		t.Errorf("expected expected '200', got %q", f.Expected())
	}
}

func TestKVFieldFilter_Match_CaseSensitive(t *testing.T) {
	f, _ := filter.NewKVFieldFilter("level", "=", "INFO", false)
	if !f.Match("time=12:00 level=INFO msg=hello") {
		t.Error("expected match")
	}
	if f.Match("time=12:00 level=info msg=hello") {
		t.Error("expected no match (case mismatch)")
	}
}

func TestKVFieldFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := filter.NewKVFieldFilter("level", "=", "INFO", true)
	if !f.Match("time=12:00 level=info msg=hello") {
		t.Error("expected case-insensitive match")
	}
}

func TestKVFieldFilter_Match_MissingKey(t *testing.T) {
	f, _ := filter.NewKVFieldFilter("level", "=", "INFO", false)
	if f.Match("time=12:00 msg=hello") {
		t.Error("expected no match when key absent")
	}
}

func TestKVFieldFilter_Match_CustomSep(t *testing.T) {
	f, _ := filter.NewKVFieldFilter("level", ":", "warn", false)
	if !f.Match("level:warn region:us") {
		t.Error("expected match with custom separator")
	}
}

func TestKVFieldFilter_Transform_Passthrough(t *testing.T) {
	f, _ := filter.NewKVFieldFilter("level", "=", "info", false)
	line := "level=info msg=ok"
	if got := f.Transform(line); got != line {
		t.Errorf("expected passthrough, got %q", got)
	}
}
