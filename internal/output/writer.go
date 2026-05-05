package output

import (
	"bufio"
	"fmt"
	"io"
)

// Writer wraps a buffered writer and tracks the number of lines written.
type Writer struct {
	w     *bufio.Writer
	count int64
}

// NewWriter creates a new Writer wrapping the given io.Writer.
// bufSize <= 0 uses the default buffer size.
func NewWriter(w io.Writer, bufSize int) *Writer {
	if bufSize <= 0 {
		bufSize = 4096
	}
	return &Writer{
		w: bufio.NewWriterSize(w, bufSize),
	}
}

// WriteLine writes a single line followed by a newline character.
func (w *Writer) WriteLine(line string) error {
	_, err := fmt.Fprintln(w.w, line)
	if err != nil {
		return err
	}
	w.count++
	return nil
}

// Flush flushes any buffered data to the underlying writer.
func (w *Writer) Flush() error {
	return w.w.Flush()
}

// Count returns the number of lines successfully written.
func (w *Writer) Count() int64 {
	return w.count
}
