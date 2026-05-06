package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
)

func TestParse_MinLengthFlag(t *testing.T) {
	cfg, err := config.Parse([]string{"-min-length", "20"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.MinLength != 20 {
		t.Errorf("expected MinLength=20, got %d", cfg.MinLength)
	}
}

func TestParse_MinLengthDefault(t *testing.T) {
	cfg, err := config.Parse([]string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.MinLength != 0 {
		t.Errorf("expected default MinLength=0, got %d", cfg.MinLength)
	}
}

func TestParse_MinLengthNegative(t *testing.T) {
	_, err := config.Parse([]string{"-min-length", "-5"})
	if err == nil {
		t.Fatal("expected error for negative min-length, got nil")
	}
}

func TestParse_MinLengthAndMaxLength(t *testing.T) {
	cfg, err := config.Parse([]string{"-min-length", "3", "-max-length", "50"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.MinLength != 3 {
		t.Errorf("expected MinLength=3, got %d", cfg.MinLength)
	}
	if cfg.MaxLength != 50 {
		t.Errorf("expected MaxLength=50, got %d", cfg.MaxLength)
	}
}
