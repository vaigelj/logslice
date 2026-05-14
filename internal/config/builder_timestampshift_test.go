package config_test

import (
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/config"
	"github.com/yourorg/logslice/internal/filter"
)

func TestBuildFilterChain_WithTimestampShift(t *testing.T) {
	cfg := &config.Config{
		TimestampShiftLayout:   "2006-01-02T15:04:05",
		TimestampShiftDuration: time.Hour,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
	line := "2024-03-01T10:00:00 hello"
	matched, transformed := chain.Apply(line)
	if !matched {
		t.Error("expected line to match")
	}
	want := "2024-03-01T11:00:00 hello"
	if transformed != want {
		t.Errorf("Apply() transformed = %q; want %q", transformed, want)
	}
}

func TestBuildFilterChain_TimestampShiftEmptyLayoutSkipped(t *testing.T) {
	cfg := &config.Config{
		TimestampShiftLayout:   "",
		TimestampShiftDuration: time.Hour,
	}
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBuildFilterChain_TimestampShiftZeroDurationSkipped(t *testing.T) {
	cfg := &config.Config{
		TimestampShiftLayout:   "2006-01-02T15:04:05",
		TimestampShiftDuration: 0,
	}
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBuildFilterChain_TimestampShiftAndRegex(t *testing.T) {
	cfg := &config.Config{
		Pattern:                "hello",
		TimestampShiftLayout:   "2006-01-02T15:04:05",
		TimestampShiftDuration: -time.Hour,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// line that matches regex and has a timestamp
	line := "2024-06-01T12:00:00 hello world"
	matched, transformed := chain.Apply(line)
	if !matched {
		t.Error("expected line to match")
	}
	want := "2024-06-01T11:00:00 hello world"
	if transformed != want {
		t.Errorf("transformed = %q; want %q", transformed, want)
	}
	// line that does not match regex
	line2 := "2024-06-01T12:00:00 goodbye world"
	matched2, _ := chain.Apply(line2)
	if matched2 {
		t.Error("expected non-matching line to be rejected")
	}
	_ = filter.NewChain // keep import used
}
