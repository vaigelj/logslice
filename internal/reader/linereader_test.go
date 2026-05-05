package reader

import (
	"strings"
	"testing"
)

func TestNewLineReader_DefaultBufSize(t *testing.T) {
	r := strings.NewReader("hello")
	lr := NewLineReader(r, 0)
	if lr.BufSize() != 64*1024 {
		t.Errorf("expected default buf size %d, got %d", 64*1024, lr.BufSize())
	}
}

func TestNewLineReader_CustomBufSize(t *testing.T) {
	r := strings.NewReader("hello")
	lr := NewLineReader(r, 4096)
	if lr.BufSize() != 4096 {
		t.Errorf("expected buf size 4096, got %d", lr.BufSize())
	}
}

func TestNewLineReader_ReaderAccessor(t *testing.T) {
	r := strings.NewReader("hello")
	lr := NewLineReader(r, 0)
	if lr.Reader() != r {
		t.Error("expected Reader() to return the original io.Reader")
	}
}

func TestLines_EmptyInput(t *testing.T) {
	lr := NewLineReader(strings.NewReader(""), 0)
	errCh := make(chan error, 1)
	var lines []string
	for line := range lr.Lines(errCh) {
		lines = append(lines, line)
	}
	if len(lines) != 0 {
		t.Errorf("expected 0 lines, got %d", len(lines))
	}
}

func TestLines_MultipleLines(t *testing.T) {
	input := "first\nsecond\nthird"
	lr := NewLineReader(strings.NewReader(input), 0)
	errCh := make(chan error, 1)
	var lines []string
	for line := range lr.Lines(errCh) {
		lines = append(lines, line)
	}
	expected := []string{"first", "second", "third"}
	if len(lines) != len(expected) {
		t.Fatalf("expected %d lines, got %d", len(expected), len(lines))
	}
	for i, l := range lines {
		if l != expected[i] {
			t.Errorf("line %d: expected %q, got %q", i, expected[i], l)
		}
	}
}

func TestLines_TrailingNewline(t *testing.T) {
	input := "alpha\nbeta\n"
	lr := NewLineReader(strings.NewReader(input), 0)
	errCh := make(chan error, 1)
	var lines []string
	for line := range lr.Lines(errCh) {
		lines = append(lines, line)
	}
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}
}

func TestLines_NilErrCh(t *testing.T) {
	input := "only line"
	lr := NewLineReader(strings.NewReader(input), 0)
	var lines []string
	for line := range lr.Lines(nil) {
		lines = append(lines, line)
	}
	if len(lines) != 1 || lines[0] != "only line" {
		t.Errorf("unexpected lines: %v", lines)
	}
}
