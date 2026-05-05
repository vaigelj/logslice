// Package output provides utilities for writing filtered log lines
// and reporting processing statistics.
//
// Writer wraps a buffered io.Writer and exposes a simple WriteLine
// method along with a line counter, making it easy to track how many
// log lines were emitted during a slicing run.
//
// Stats collects summary information (lines read, lines written, and
// elapsed duration) and can print a human-readable report to any
// io.Writer.
package output
