package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
)

func TestBuildFilterChain_WithMinLength(t *testing.T) {
	cfg := &config.Config{
		MinLength: 10,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
	if !chain.Match("hello world!") {
		t.Error("expected line of length 12 to match min-length 10")
	}
	if chain.Match("short") {
		t.Error("expected line of length 5 to not match min-length 10")
	}
}

func TestBuildFilterChain_MinLengthZeroSkipped(t *testing.T) {
	cfg := &config.Config{
		MinLength: 0,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// no filter added, empty line should still pass
	if !chain.Match("") {
		t.Error("expected empty line to match when no filters are set")
	}
}

func TestBuildFilterChain_MinLengthAndMaxLength(t *testing.T) {
	cfg := &config.Config{
		MinLength: 5,
		MaxLength: 10,
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !chain.Match("hello") {
		t.Error("expected 'hello' (len 5) to match [5,10]")
	}
	if !chain.Match("helloworld") {
		t.Error("expected 'helloworld' (len 10) to match [5,10]")
	}
	if chain.Match("hi") {
		t.Error("expected 'hi' (len 2) to not match [5,10]")
	}
	if chain.Match("hello world!!") {
		t.Error("expected 'hello world!!' (len 13) to not match [5,10]")
	}
}

func TestBuildFilterChain_NoMinLength(t *testing.T) {
	cfg := &config.Config{}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
}
