package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewJSONRangeFilter_EmptyField(t *testing.T) {
	_, err := filter.NewJSONRangeFilter("", 0, 100)
	if err == nil {
		t.Fatal("expected error for empty field")
	}
}

func TestNewJSONRangeFilter_MinExceedsMax(t *testing.T) {
	_, err := filter.NewJSONRangeFilter("latency", 200, 100)
	if err == nil {
		t.Fatal("expected error when min > max")
	}
}

func TestNewJSONRangeFilter_Valid(t *testing.T) {
	f, err := filter.NewJSONRangeFilter("latency", 10, 500)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Field() != "latency" {
		t.Errorf("expected field 'latency', got %q", f.Field())
	}
	if f.Min() != 10 {
		t.Errorf("expected min 10, got %v", f.Min())
	}
	if f.Max() != 500 {
		t.Errorf("expected max 500, got %v", f.Max())
	}
}

func TestJSONRangeFilter_Match_InRange(t *testing.T) {
	f, _ := filter.NewJSONRangeFilter("status", 200, 299)
	if !f.Match(`{"status": 200}`) {
		t.Error("expected match for value at lower bound")
	}
	if !f.Match(`{"status": 250}`) {
		t.Error("expected match for value in middle")
	}
	if !f.Match(`{"status": 299}`) {
		t.Error("expected match for value at upper bound")
	}
}

func TestJSONRangeFilter_Match_OutOfRange(t *testing.T) {
	f, _ := filter.NewJSONRangeFilter("status", 200, 299)
	if f.Match(`{"status": 199}`) {
		t.Error("expected no match for value below range")
	}
	if f.Match(`{"status": 300}`) {
		t.Error("expected no match for value above range")
	}
}

func TestJSONRangeFilter_Match_StringNumericField(t *testing.T) {
	f, _ := filter.NewJSONRangeFilter("score", 1, 10)
	if !f.Match(`{"score": "7"}`) {
		t.Error("expected match for numeric string value in range")
	}
	if f.Match(`{"score": "11"}`) {
		t.Error("expected no match for numeric string value out of range")
	}
}

func TestJSONRangeFilter_Match_MissingField(t *testing.T) {
	f, _ := filter.NewJSONRangeFilter("latency", 0, 100)
	if f.Match(`{"other": 50}`) {
		t.Error("expected no match when field is missing")
	}
}

func TestJSONRangeFilter_Match_InvalidJSON(t *testing.T) {
	f, _ := filter.NewJSONRangeFilter("latency", 0, 100)
	if f.Match(`not json`) {
		t.Error("expected no match for invalid JSON")
	}
}

func TestJSONRangeFilter_Match_NonNumericField(t *testing.T) {
	f, _ := filter.NewJSONRangeFilter("level", 0, 100)
	if f.Match(`{"level": "info"}`) {
		t.Error("expected no match for non-numeric string value")
	}
}

func TestJSONRangeFilter_EqualMinMax(t *testing.T) {
	f, err := filter.NewJSONRangeFilter("code", 404, 404)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !f.Match(`{"code": 404}`) {
		t.Error("expected match when value equals min==max")
	}
	if f.Match(`{"code": 200}`) {
		t.Error("expected no match for value outside single-point range")
	}
}
