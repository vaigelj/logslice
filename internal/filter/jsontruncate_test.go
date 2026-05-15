package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewJSONTruncateFilter_EmptyField(t *testing.T) {
	_, err := filter.NewJSONTruncateFilter("", 10, "...")
	if err == nil {
		t.Fatal("expected error for empty field")
	}
}

func TestNewJSONTruncateFilter_InvalidMaxLen(t *testing.T) {
	_, err := filter.NewJSONTruncateFilter("msg", 0, "")
	if err == nil {
		t.Fatal("expected error for zero maxLen")
	}
}

func TestNewJSONTruncateFilter_SuffixTooLong(t *testing.T) {
	_, err := filter.NewJSONTruncateFilter("msg", 3, "...")
	if err == nil {
		t.Fatal("expected error when suffix length >= maxLen")
	}
}

func TestNewJSONTruncateFilter_Valid(t *testing.T) {
	f, err := filter.NewJSONTruncateFilter("msg", 10, "...")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Field() != "msg" {
		t.Errorf("expected field 'msg', got %q", f.Field())
	}
	if f.MaxLen() != 10 {
		t.Errorf("expected maxLen 10, got %d", f.MaxLen())
	}
	if f.Suffix() != "..." {
		t.Errorf("expected suffix '...', got %q", f.Suffix())
	}
}

func TestJSONTruncateFilter_Match_AlwaysTrue(t *testing.T) {
	f, _ := filter.NewJSONTruncateFilter("msg", 10, "")
	if !f.Match(`{"msg":"hello"}`) {
		t.Error("expected Match to always return true")
	}
}

func TestJSONTruncateFilter_Transform_ShortValue(t *testing.T) {
	f, _ := filter.NewJSONTruncateFilter("msg", 20, "...")
	input := `{"msg":"short"}`
	out := f.Transform(input)
	if !containsSubstr(out, `"short"`) {
		t.Errorf("expected unchanged value, got %q", out)
	}
}

func TestJSONTruncateFilter_Transform_LongValue(t *testing.T) {
	f, _ := filter.NewJSONTruncateFilter("msg", 10, "...")
	input := `{"msg":"this is a very long message"}`
	out := f.Transform(input)
	if !containsSubstr(out, `"this is..."`) {
		t.Errorf("expected truncated value, got %q", out)
	}
}

func TestJSONTruncateFilter_Transform_NonStringField(t *testing.T) {
	f, _ := filter.NewJSONTruncateFilter("count", 5, "")
	input := `{"count":42}`
	out := f.Transform(input)
	if out != input {
		t.Errorf("expected unchanged line for non-string field, got %q", out)
	}
}

func TestJSONTruncateFilter_Transform_InvalidJSON(t *testing.T) {
	f, _ := filter.NewJSONTruncateFilter("msg", 10, "...")
	input := `not json`
	out := f.Transform(input)
	if out != input {
		t.Errorf("expected unchanged line for invalid JSON, got %q", out)
	}
}

func containsSubstr(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(s) > 0 && containsSubstrHelper(s, sub))
}

func containsSubstrHelper(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
