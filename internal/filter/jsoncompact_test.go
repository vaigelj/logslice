package filter

import (
	"testing"
)

func TestNewJSONCompactFilter_Passthrough(t *testing.T) {
	f, err := NewJSONCompactFilter(true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !f.Passthrough() {
		t.Error("expected Passthrough() == true")
	}
}

func TestNewJSONCompactFilter_NoPassthrough(t *testing.T) {
	f, err := NewJSONCompactFilter(false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Passthrough() {
		t.Error("expected Passthrough() == false")
	}
}

func TestJSONCompactFilter_Match_ValidJSON(t *testing.T) {
	f, _ := NewJSONCompactFilter(false)
	if !f.Match(`{ "key": "value" }`) {
		t.Error("expected Match to return true for valid JSON")
	}
}

func TestJSONCompactFilter_Match_InvalidJSON_Passthrough(t *testing.T) {
	f, _ := NewJSONCompactFilter(true)
	if !f.Match("not json at all") {
		t.Error("expected Match to return true when passthrough=true")
	}
}

func TestJSONCompactFilter_Match_InvalidJSON_NoPassthrough(t *testing.T) {
	f, _ := NewJSONCompactFilter(false)
	if f.Match("not json at all") {
		t.Error("expected Match to return false when passthrough=false")
	}
}

func TestJSONCompactFilter_Transform_CompactsJSON(t *testing.T) {
	f, _ := NewJSONCompactFilter(false)
	input := `{  "a":  1,  "b":  [1, 2, 3]  }`
	got := f.Transform(input)
	want := `{"a":1,"b":[1,2,3]}`
	if got != want {
		t.Errorf("Transform() = %q, want %q", got, want)
	}
}

func TestJSONCompactFilter_Transform_AlreadyCompact(t *testing.T) {
	f, _ := NewJSONCompactFilter(false)
	input := `{"x":true}`
	got := f.Transform(input)
	if got != input {
		t.Errorf("Transform() = %q, want %q", got, input)
	}
}

func TestJSONCompactFilter_Transform_NonJSON_Passthrough(t *testing.T) {
	f, _ := NewJSONCompactFilter(true)
	input := "plain text line"
	got := f.Transform(input)
	if got != input {
		t.Errorf("Transform() = %q, want unchanged %q", got, input)
	}
}

func TestJSONCompactFilter_InChain(t *testing.T) {
	compact, _ := NewJSONCompactFilter(false)
	chain, err := NewChain(compact)
	if err != nil {
		t.Fatalf("NewChain error: %v", err)
	}
	if !chain.Match(`{"ok":true}`) {
		t.Error("chain should match valid JSON")
	}
	if chain.Match("garbage") {
		t.Error("chain should reject non-JSON when passthrough=false")
	}
}
