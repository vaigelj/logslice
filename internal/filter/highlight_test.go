package filter_test

import (
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewHighlightFilter_EmptyPattern(t *testing.T) {
	_, err := filter.NewHighlightFilter("", "\033[31m")
	if err == nil {
		t.Fatal("expected error for empty pattern")
	}
}

func TestNewHighlightFilter_EmptyColor(t *testing.T) {
	_, err := filter.NewHighlightFilter("ERROR", "")
	if err == nil {
		t.Fatal("expected error for empty color")
	}
}

func TestNewHighlightFilter_InvalidPattern(t *testing.T) {
	_, err := filter.NewHighlightFilter("[invalid", "\033[31m")
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
}

func TestNewHighlightFilter_Valid(t *testing.T) {
	h, err := filter.NewHighlightFilter("ERROR", "\033[31m")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h.Pattern() != "ERROR" {
		t.Errorf("expected pattern ERROR, got %q", h.Pattern())
	}
	if h.Color() != "\033[31m" {
		t.Errorf("unexpected color value")
	}
}

func TestHighlightFilter_Match_AlwaysTrue(t *testing.T) {
	h, _ := filter.NewHighlightFilter("ERROR", "\033[31m")
	for _, line := range []string{"", "INFO ok", "ERROR bad"} {
		if !h.Match(line) {
			t.Errorf("Match(%q) = false, want true", line)
		}
	}
}

func TestHighlightFilter_Transform_WrapsMatch(t *testing.T) {
	const red = "\033[31m"
	h, _ := filter.NewHighlightFilter("ERROR", red)
	out := h.Transform("line with ERROR inside")
	if !strings.Contains(out, red+"ERROR"+"\033[0m") {
		t.Errorf("expected highlighted ERROR in output, got %q", out)
	}
}

func TestHighlightFilter_Transform_NoMatch(t *testing.T) {
	h, _ := filter.NewHighlightFilter("ERROR", "\033[31m")
	line := "everything is fine"
	if got := h.Transform(line); got != line {
		t.Errorf("expected unchanged line, got %q", got)
	}
}

func TestHighlightFilter_Transform_MultipleMatches(t *testing.T) {
	const yellow = "\033[33m"
	h, _ := filter.NewHighlightFilter("WARN", yellow)
	out := h.Transform("WARN first WARN second")
	count := strings.Count(out, yellow)
	if count != 2 {
		t.Errorf("expected 2 highlighted occurrences, got %d", count)
	}
}

func TestHighlightFilter_String(t *testing.T) {
	h, _ := filter.NewHighlightFilter("DEBUG", "\033[34m")
	if !strings.Contains(h.String(), "DEBUG") {
		t.Errorf("String() should contain pattern name, got %q", h.String())
	}
}
