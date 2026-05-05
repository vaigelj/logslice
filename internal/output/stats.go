package output

import (
	"fmt"
	"io"
	"time"
)

// Stats holds processing statistics for a log slicing run.
type Stats struct {
	LinesRead    int64
	LinesWritten int64
	Duration     time.Duration
}

// Print writes a human-readable summary of the stats to w.
func (s Stats) Print(w io.Writer) {
	fmt.Fprintf(w, "Lines read:    %d\n", s.LinesRead)
	fmt.Fprintf(w, "Lines written: %d\n", s.LinesWritten)
	fmt.Fprintf(w, "Lines dropped: %d\n", s.LinesRead-s.LinesWritten)
	if s.LinesRead > 0 {
		pct := float64(s.LinesWritten) / float64(s.LinesRead) * 100.0
		fmt.Fprintf(w, "Match rate:    %.1f%%\n", pct)
	}
	fmt.Fprintf(w, "Duration:      %s\n", s.Duration.Round(time.Millisecond))
}
