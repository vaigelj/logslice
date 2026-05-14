package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewJSONPresentFilter_EmptyField(t *testing.T) {
	_, err := filter.NewJSONPresentFilter("", false)
	if err == nil {
		t.Fatal("expected error for empty field")
	}
}

func TestNewJSONPresentFilter_Valid(t *testing.T) {
	f, err := filter.NewJSONPresentFilter("level", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Field() != "level" {
		t.Errorf("expected field 'level', got %q", f.Field())
	}
	if f.Inverted() {
		t.Error("expected inverted=false")
	}
}

func TestJSONPresentFilter_Match_FieldPresent(t *testing.T) {
	f, _ := filter.NewJSONPresentFilter("level", false)
	if !f.Match(`{"level":"info","msg":"ok"}`) {
		t.Error("expected match when field is present")
	}
}

func TestJSONPresentFilter_Match_FieldAbsent(t *testing.T) {
	f, _ := filter.NewJSONPresentFilter("level", false)
	if f.Match(`{"msg":"ok"}`) {
		t.Error("expected no match when field is absent")
	}
}

func TestJSONPresentFilter_Match_Inverted_FieldAbsent(t *testing.T) {
	f, _ := filter.NewJSONPresentFilter("level", true)
	if !f.Match(`{"msg":"ok"}`) {
		t.Error("expected match when field is absent and filter is inverted")
	}
}

func TestJSONPresentFilter_Match_Inverted_FieldPresent(t *testing.T) {
	f, _ := filter.NewJSONPresentFilter("level", true)
	if f.Match(`{"level":"info"}`) {
		t.Error("expected no match when field is present and filter is inverted")
	}
}

func TestJSONPresentFilter_Match_NestedField(t *testing.T) {
	f, _ := filter.NewJSONPresentFilter("meta.user", false)
	if !f.Match(`{"meta":{"user":"alice"}}`) {
		t.Error("expected match for nested field")
	}
}

func TestJSONPresentFilter_Match_NestedField_Absent(t *testing.T) {
	f, _ := filter.NewJSONPresentFilter("meta.user", false)
	if f.Match(`{"meta":{"host":"srv1"}}`) {
		t.Error("expected no match when nested field is absent")
	}
}

func TestJSONPresentFilter_Match_InvalidJSON(t *testing.T) {
	f, _ := filter.NewJSONPresentFilter("level", false)
	if f.Match(`not json`) {
		t.Error("expected no match for invalid JSON")
	}
}

func TestJSONPresentFilter_InChain(t *testing.T) {
	present, _ := filter.NewJSONPresentFilter("level", false)
	chain := filter.NewChain(present)
	if !chain.Match(`{"level":"warn"}`) {
		t.Error("expected chain to match when field is present")
	}
	if chain.Match(`{"msg":"no level"}`) {
		t.Error("expected chain to reject when field is absent")
	}
}
