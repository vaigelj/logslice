package output_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/output"
)

func TestNewWriter_DefaultBufSize(t *testing.T) {
	var buf bytes.Buffer
	w := output.NewWriter(&buf, 0)
	if w == nil {
		t.Fatal("expected non-nil Writer")
	}
}

func TestNewWriter_CustomBufSize(t *testing.T) {
	var buf bytes.Buffer
	w := output.NewWriter(&buf, 8192)
	if w == nil {
		t.Fatal("expected non-nil Writer")
	}
}

func TestWriter_WriteLineAndCount(t *testing.T) {
	var buf bytes.Buffer
	w := output.NewWriter(&buf, 0)

	lines := []string{"line one", "line two", "line three"}
	for _, l := range lines {
		if err := w.WriteLine(l); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if err := w.Flush(); err != nil {
		t.Fatalf("flush error: %v", err)
	}

	if w.Count() != int64(len(lines)) {
		t.Errorf("expected count %d, got %d", len(lines), w.Count())
	}

	got := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	if len(got) != len(lines) {
		t.Fatalf("expected %d lines in output, got %d", len(lines), len(got))
	}
	for i, l := range lines {
		if got[i] != l {
			t.Errorf("line %d: expected %q, got %q", i, l, got[i])
		}
	}
}

func TestWriter_CountStartsAtZero(t *testing.T) {
	var buf bytes.Buffer
	w := output.NewWriter(&buf, 0)
	if w.Count() != 0 {
		t.Errorf("expected initial count 0, got %d", w.Count())
	}
}

func TestWriter_FlushEmpty(t *testing.T) {
	var buf bytes.Buffer
	w := output.NewWriter(&buf, 0)
	if err := w.Flush(); err != nil {
		t.Errorf("unexpected flush error on empty writer: %v", err)
	}
}
