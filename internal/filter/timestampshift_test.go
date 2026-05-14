package filter_test

import (
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/filter"
)

const tsLayout = "2006-01-02T15:04:05"

func TestNewTimestampShiftFilter_EmptyLayout(t *testing.T) {
	_, err := filter.NewTimestampShiftFilter("", time.Hour)
	if err == nil {
		t.Fatal("expected error for empty layout")
	}
}

func TestNewTimestampShiftFilter_Valid(t *testing.T) {
	f, err := filter.NewTimestampShiftFilter(tsLayout, time.Hour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Layout() != tsLayout {
		t.Errorf("Layout() = %q; want %q", f.Layout(), tsLayout)
	}
	if f.Shift() != time.Hour {
		t.Errorf("Shift() = %v; want %v", f.Shift(), time.Hour)
	}
}

func TestTimestampShiftFilter_Match_AlwaysTrue(t *testing.T) {
	f, _ := filter.NewTimestampShiftFilter(tsLayout, time.Minute)
	for _, line := range []string{"", "no timestamp", "2024-01-01T00:00:00 hello"} {
		if !f.Match(line) {
			t.Errorf("Match(%q) = false; want true", line)
		}
	}
}

func TestTimestampShiftFilter_Transform_ShiftsTimestamp(t *testing.T) {
	f, _ := filter.NewTimestampShiftFilter(tsLayout, time.Hour)
	input := "2024-03-01T10:00:00 some log message"
	got := f.Transform(input)
	want := "2024-03-01T11:00:00 some log message"
	if got != want {
		t.Errorf("Transform() = %q; want %q", got, want)
	}
}

func TestTimestampShiftFilter_Transform_NegativeShift(t *testing.T) {
	f, _ := filter.NewTimestampShiftFilter(tsLayout, -30*time.Minute)
	input := "2024-03-01T10:00:00 warn: disk usage high"
	got := f.Transform(input)
	want := "2024-03-01T09:30:00 warn: disk usage high"
	if got != want {
		t.Errorf("Transform() = %q; want %q", got, want)
	}
}

func TestTimestampShiftFilter_Transform_NoTimestamp(t *testing.T) {
	f, _ := filter.NewTimestampShiftFilter(tsLayout, time.Hour)
	input := "no timestamp here"
	got := f.Transform(input)
	if got != input {
		t.Errorf("Transform() = %q; want %q", got, input)
	}
}

func TestTimestampShiftFilter_Transform_ShortLine(t *testing.T) {
	f, _ := filter.NewTimestampShiftFilter(tsLayout, time.Hour)
	input := "2024"
	got := f.Transform(input)
	if got != input {
		t.Errorf("Transform() = %q; want %q", got, input)
	}
}

func TestTimestampShiftFilter_Transform_ZeroShift(t *testing.T) {
	f, _ := filter.NewTimestampShiftFilter(tsLayout, 0)
	input := "2024-06-15T08:30:00 info: startup complete"
	got := f.Transform(input)
	if got != input {
		t.Errorf("Transform() = %q; want %q", got, input)
	}
}
