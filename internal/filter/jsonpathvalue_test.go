package filter

import (
	"testing"
)

func TestNewJSONPathValueFilter_EmptyPath(t *testing.T) {
	_, err := NewJSONPathValueFilter("", "value", false)
	if err == nil {
		t.Fatal("expected error for empty path")
	}
}

func TestNewJSONPathValueFilter_EmptyExpected(t *testing.T) {
	_, err := NewJSONPathValueFilter("a.b", "", false)
	if err == nil {
		t.Fatal("expected error for empty expected")
	}
}

func TestNewJSONPathValueFilter_Valid(t *testing.T) {
	f, err := NewJSONPathValueFilter("level", "error", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Path() != "level" {
		t.Errorf("expected path 'level', got %q", f.Path())
	}
	if f.Expected() != "error" {
		t.Errorf("expected expected 'error', got %q", f.Expected())
	}
	if f.CaseInsensitive() {
		t.Error("expected case-sensitive")
	}
}

func TestJSONPathValueFilter_Match_TopLevel(t *testing.T) {
	f, _ := NewJSONPathValueFilter("level", "error", false)
	if !f.Match(`{"level":"error","msg":"oops"}`) {
		t.Error("expected match")
	}
	if f.Match(`{"level":"info","msg":"ok"}`) {
		t.Error("expected no match")
	}
}

func TestJSONPathValueFilter_Match_NestedPath(t *testing.T) {
	f, _ := NewJSONPathValueFilter("request.method", "GET", false)
	if !f.Match(`{"request":{"method":"GET","path":"/"}}`) {
		t.Error("expected match for nested path")
	}
	if f.Match(`{"request":{"method":"POST","path":"/"}}`) {
		t.Error("expected no match for wrong value")
	}
}

func TestJSONPathValueFilter_Match_CaseInsensitive(t *testing.T) {
	f, _ := NewJSONPathValueFilter("level", "ERROR", true)
	if !f.Match(`{"level":"error"}`) {
		t.Error("expected case-insensitive match")
	}
}

func TestJSONPathValueFilter_Match_InvalidJSON(t *testing.T) {
	f, _ := NewJSONPathValueFilter("level", "error", false)
	if f.Match(`not json`) {
		t.Error("expected no match for invalid JSON")
	}
}

func TestJSONPathValueFilter_Match_MissingPath(t *testing.T) {
	f, _ := NewJSONPathValueFilter("a.b.c", "val", false)
	if f.Match(`{"a":{"x":"val"}}`) {
		t.Error("expected no match when path is missing")
	}
}

func TestJSONPathValueFilter_Match_NumericValue(t *testing.T) {
	f, _ := NewJSONPathValueFilter("code", "200", false)
	if !f.Match(`{"code":200}`) {
		t.Error("expected match for numeric value rendered as string")
	}
}

func TestJSONPathValueFilter_InChain(t *testing.T) {
	f, _ := NewJSONPathValueFilter("env", "prod", false)
	chain, _ := NewChain(f)
	if !chain.Match(`{"env":"prod","svc":"api"}`) {
		t.Error("expected chain match")
	}
	if chain.Match(`{"env":"dev","svc":"api"}`) {
		t.Error("expected chain no match")
	}
}
