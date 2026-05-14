package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewJSONKeysFilter_EmptyKeys(t *testing.T) {
	_, err := filter.NewJSONKeysFilter([]string{}, false)
	if err == nil {
		t.Fatal("expected error for empty keys list")
	}
}

func TestNewJSONKeysFilter_EmptyKeyEntry(t *testing.T) {
	_, err := filter.NewJSONKeysFilter([]string{"valid", ""}, false)
	if err == nil {
		t.Fatal("expected error for empty key entry")
	}
}

func TestNewJSONKeysFilter_Valid(t *testing.T) {
	f, err := filter.NewJSONKeysFilter([]string{"level", "msg"}, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(f.Keys()) != 2 {
		t.Errorf("expected 2 keys, got %d", len(f.Keys()))
	}
}

func TestJSONKeysFilter_Match_AllPresent(t *testing.T) {
	f, _ := filter.NewJSONKeysFilter([]string{"level", "msg"}, false)
	if !f.Match(`{"level":"info","msg":"hello"}`) {
		t.Error("expected match when all keys present")
	}
}

func TestJSONKeysFilter_Match_MissingKey(t *testing.T) {
	f, _ := filter.NewJSONKeysFilter([]string{"level", "msg", "ts"}, false)
	if f.Match(`{"level":"info","msg":"hello"}`) {
		t.Error("expected no match when key 'ts' is missing")
	}
}

func TestJSONKeysFilter_Match_InvalidJSON(t *testing.T) {
	f, _ := filter.NewJSONKeysFilter([]string{"level"}, false)
	if f.Match(`not json`) {
		t.Error("expected no match for invalid JSON")
	}
}

func TestJSONKeysFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := filter.NewJSONKeysFilter([]string{"Level", "MSG"}, true)
	if !f.Match(`{"level":"info","msg":"hello"}`) {
		t.Error("expected case-insensitive match")
	}
}

func TestJSONKeysFilter_Match_CaseSensitiveFails(t *testing.T) {
	f, _ := filter.NewJSONKeysFilter([]string{"Level"}, false)
	if f.Match(`{"level":"info"}`) {
		t.Error("expected no match with case-sensitive key mismatch")
	}
}

func TestJSONKeysFilter_Accessors(t *testing.T) {
	f, _ := filter.NewJSONKeysFilter([]string{"a", "b"}, true)
	if !f.CaseInsensitive() {
		t.Error("expected CaseInsensitive to return true")
	}
	if len(f.Keys()) != 2 {
		t.Errorf("expected 2 keys, got %d", len(f.Keys()))
	}
}

func TestJSONKeysFilter_Match_EmptyObject(t *testing.T) {
	f, _ := filter.NewJSONKeysFilter([]string{"key"}, false)
	if f.Match(`{}`) {
		t.Error("expected no match for empty JSON object")
	}
}
