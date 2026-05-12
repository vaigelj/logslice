package filter

import (
	"testing"
	"time"
)

func TestNewRateLimitFilter_InvalidMax(t *testing.T) {
	_, err := NewRateLimitFilter(0, time.Second)
	if err == nil {
		t.Fatal("expected error for maxPerWindow=0")
	}
	_, err = NewRateLimitFilter(-1, time.Second)
	if err == nil {
		t.Fatal("expected error for maxPerWindow=-1")
	}
}

func TestNewRateLimitFilter_InvalidWindow(t *testing.T) {
	_, err := NewRateLimitFilter(5, 0)
	if err == nil {
		t.Fatal("expected error for window=0")
	}
	_, err = NewRateLimitFilter(5, -time.Second)
	if err == nil {
		t.Fatal("expected error for negative window")
	}
}

func TestNewRateLimitFilter_Valid(t *testing.T) {
	f, err := NewRateLimitFilter(3, time.Minute)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.MaxPerWindow() != 3 {
		t.Errorf("expected MaxPerWindow=3, got %d", f.MaxPerWindow())
	}
	if f.Window() != time.Minute {
		t.Errorf("expected Window=1m, got %s", f.Window())
	}
	if f.Count() != 0 {
		t.Errorf("expected Count=0, got %d", f.Count())
	}
}

func TestRateLimitFilter_AllowsUpToMax(t *testing.T) {
	f, _ := NewRateLimitFilter(3, time.Hour)
	for i := 0; i < 3; i++ {
		if !f.Match("line") {
			t.Errorf("expected Match=true on call %d", i+1)
		}
	}
	if f.Count() != 3 {
		t.Errorf("expected Count=3, got %d", f.Count())
	}
}

func TestRateLimitFilter_RejectsOverMax(t *testing.T) {
	f, _ := NewRateLimitFilter(2, time.Hour)
	f.Match("a")
	f.Match("b")
	if f.Match("c") {
		t.Error("expected Match=false after exceeding limit")
	}
}

func TestRateLimitFilter_WindowReset(t *testing.T) {
	// Use a very short window so we can trigger a reset.
	f, _ := NewRateLimitFilter(1, 50*time.Millisecond)
	if !f.Match("first") {
		t.Fatal("expected first line to match")
	}
	if f.Match("second") {
		t.Fatal("expected second line to be rejected within window")
	}
	time.Sleep(60 * time.Millisecond)
	if !f.Match("third") {
		t.Fatal("expected match after window reset")
	}
	if f.Count() != 1 {
		t.Errorf("expected Count=1 after reset, got %d", f.Count())
	}
}

func TestRateLimitFilter_InChain(t *testing.T) {
	rl, _ := NewRateLimitFilter(2, time.Hour)
	chain := NewChain(rl)
	if !chain.Match("x") {
		t.Error("expected chain match 1")
	}
	if !chain.Match("y") {
		t.Error("expected chain match 2")
	}
	if chain.Match("z") {
		t.Error("expected chain reject after limit")
	}
}
