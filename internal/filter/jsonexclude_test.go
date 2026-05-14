package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewJSONExcludeFilter_EmptyField(t *testing.T) {
	_, err := filter.NewJSONExcludeFilter("", "error", false)
	if err == nil {
		t.Fatal("expected error for empty field")
	}
}

func TestNewJSONExcludeFilter_EmptyExpected(t *testing.T) {
	_, err := filter.NewJSONExcludeFilter("level", "", false)
	if err == nil {
		t.Fatal("expected error for empty expected value")
	}
}

func TestNewJSONExcludeFilter_Valid(t *testing.T) {
	f, err := filter.NewJSONExcludeFilter("level", "debug", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Field() != "level" {
		t.Errorf("expected field 'level', got %q", f.Field())
	}
	if f.Expected() != "debug" {
		t.Errorf("expected expected 'debug', got %q", f.Expected())
	}
	if f.CaseInsensitive() {
		t.Error("expected case-sensitive")
	}
}

func TestJSONExcludeFilter_Match_ExcludesMatchingField(t *testing.T) {
	f, _ := filter.NewJSONExcludeFilter("level", "debug", false)
	if f.Match(`{"level":"debug","msg":"hi"}`) {
		t.Error("expected line to be excluded (Match=false)")
	}
}

func TestJSONExcludeFilter_Match_AllowsNonMatchingField(t *testing.T) {
	f, _ := filter.NewJSONExcludeFilter("level", "debug", false)
	if !f.Match(`{"level":"error","msg":"hi"}`) {
		t.Error("expected line to be allowed (Match=true)")
	}
}

func TestJSONExcludeFilter_Match_FieldAbsent(t *testing.T) {
	f, _ := filter.NewJSONExcludeFilter("level", "debug", false)
	if !f.Match(`{"msg":"no level here"}`) {
		t.Error("expected line without field to be allowed")
	}
}

func TestJSONExcludeFilter_Match_InvalidJSON(t *testing.T) {
	f, _ := filter.NewJSONExcludeFilter("level", "debug", false)
	if !f.Match(`not json`) {
		t.Error("expected non-JSON line to be allowed")
	}
}

func TestJSONExcludeFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := filter.NewJSONExcludeFilter("level", "debug", true)
	if f.Match(`{"level":"DEBUG"}`) {
		t.Error("expected case-insensitive match to exclude line")
	}
	if !f.Match(`{"level":"INFO"}`) {
		t.Error("expected non-matching line to be allowed")
	}
}

func TestJSONExcludeFilter_InChain(t *testing.T) {
	exclude, _ := filter.NewJSONExcludeFilter("level", "debug", false)
	chain, _ := filter.NewChain(exclude)
	if chain.Match(`{"level":"debug"}`) {
		t.Error("chain should exclude debug lines")
	}
	if !chain.Match(`{"level":"info"}`) {
		t.Error("chain should allow info lines")
	}
}
