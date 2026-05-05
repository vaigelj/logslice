package output_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/output"
)

func TestStats_Print_AllLines(t *testing.T) {
	s := output.Stats{
		LinesRead:    100,
		LinesWritten: 100,
		Duration:     250 * time.Millisecond,
	}
	var buf bytes.Buffer
	s.Print(&buf)
	out := buf.String()

	if !strings.Contains(out, "Lines read:    100") {
		t.Errorf("missing lines read in output: %s", out)
	}
	if !strings.Contains(out, "Lines written: 100") {
		t.Errorf("missing lines written in output: %s", out)
	}
	if !strings.Contains(out, "Lines dropped: 0") {
		t.Errorf("missing lines dropped in output: %s", out)
	}
	if !strings.Contains(out, "Match rate:    100.0%") {
		t.Errorf("missing match rate in output: %s", out)
	}
}

func TestStats_Print_PartialMatch(t *testing.T) {
	s := output.Stats{
		LinesRead:    200,
		LinesWritten: 50,
		Duration:     500 * time.Millisecond,
	}
	var buf bytes.Buffer
	s.Print(&buf)
	out := buf.String()

	if !strings.Contains(out, "Lines dropped: 150") {
		t.Errorf("expected 150 dropped lines, output: %s", out)
	}
	if !strings.Contains(out, "Match rate:    25.0%") {
		t.Errorf("expected 25.0%% match rate, output: %s", out)
	}
}

func TestStats_Print_ZeroRead(t *testing.T) {
	s := output.Stats{
		LinesRead:    0,
		LinesWritten: 0,
		Duration:     0,
	}
	var buf bytes.Buffer
	// Should not panic or print match rate when no lines were read.
	s.Print(&buf)
	out := buf.String()
	if strings.Contains(out, "Match rate") {
		t.Errorf("should not print match rate when LinesRead is 0, got: %s", out)
	}
}
