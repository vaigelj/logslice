package reader

import (
	"bufio"
	"io"
)

// LineReader reads lines from an io.Reader and emits them through a channel.
type LineReader struct {
	r       io.Reader
	bufSize int
}

// NewLineReader creates a new LineReader wrapping the given io.Reader.
// bufSize controls the scanner buffer size in bytes; if <= 0, a default of 64KB is used.
func NewLineReader(r io.Reader, bufSize int) *LineReader {
	if bufSize <= 0 {
		bufSize = 64 * 1024
	}
	return &LineReader{r: r, bufSize: bufSize}
}

// Lines returns a channel that yields each line of text (without the newline).
// The channel is closed when the reader is exhausted or an error occurs.
// If an error other than io.EOF is encountered, it is sent to the errCh channel.
func (lr *LineReader) Lines(errCh chan<- error) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		scanner := bufio.NewScanner(lr.r)
		buf := make([]byte, lr.bufSize)
		scanner.Buffer(buf, lr.bufSize)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			if errCh != nil {
				errCh <- err
			}
		}
	}()
	return ch
}

// Reader returns the underlying io.Reader.
func (lr *LineReader) Reader() io.Reader {
	return lr.r
}

// BufSize returns the configured buffer size.
func (lr *LineReader) BufSize() int {
	return lr.bufSize
}
