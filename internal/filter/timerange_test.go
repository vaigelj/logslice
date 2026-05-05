package filter

import (
	"testing"
	"time"
)

const testLayout = "2006-01-02T15:04:05"

func TestNewTimeRangeFilter_EmptyLayout(t *testing.T) {
	_, err := NewTimeRangeFilter("", "", "")
	if err == nil {
		t.Fatal("expected error for empty layout")
	}
}

func TestNewTimeRangeFilter_InvalidStart(t *testing.T) {
	_, err := NewTimeRangeFilter("not-a-time", "", testLayout)
	if err == nil {
		t.Fatal("expected error for invalid start time")
	}
}

func TestNewTimeRangeFilter_InvalidEnd(t *testing.T) {
	_, err := NewTimeRangeFilter("", "not-a-time", testLayout)
	if err == nil {
		t.Fatal("expected error for invalid end time")
	}
}

func TestNewTimeRangeFilter_EndBeforeStart(t *testing.T) {
	_, err := NewTimeRangeFilter("2024-01-02T10:00:00Z", "2024-01-01T10:00:00Z", testLayout)
	if err == nil {
		t.Fatal("expected error when end is before start")
	}
}

func TestTimeRangeFilter_Match(t *testing.T) {
	f, err := NewTimeRangeFilter("2024-06-01T08:00:00Z", "2024-06-01T09:00:00Z", testLayout)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		line  string
		want  bool
	}{
		{"2024-06-01T08:30:00 INFO server started", true},
		{"2024-06-01T07:59:59 DEBUG early message", false},
		{"2024-06-01T09:00:01 WARN late message", false},
		{"no timestamp here", false},
		{"2024-06-01T08:00:00 INFO exact start", true},
		{"2024-06-01T09:00:00 INFO exact end", true},
	}

	for _, tc := range tests {
		got := f.Match(tc.line)
		if got != tc.want {
			t.Errorf("Match(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestTimeRangeFilter_OpenBounds(t *testing.T) {
	f, err := NewTimeRangeFilter("", "", testLayout)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !f.Match("2024-06-01T08:30:00 INFO anything goes") {
		t.Error("expected open-bounds filter to match any parseable line")
	}
}

func TestTimeRangeFilter_Accessors(t *testing.T) {
	start := "2024-06-01T08:00:00Z"
	end := "2024-06-01T09:00:00Z"
	f, _ := NewTimeRangeFilter(start, end, testLayout)

	if f.Layout() != testLayout {
		t.Errorf("Layout() = %q, want %q", f.Layout(), testLayout)
	}
	if f.Start().IsZero() {
		t.Error("expected non-zero Start()")
	}
	if f.End().IsZero() {
		t.Error("expected non-zero End()")
	}
	_ = f.Start().Equal(time.Time{})
}
