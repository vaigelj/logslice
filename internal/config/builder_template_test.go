package config_test

import (
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/config"
	"github.com/yourorg/logslice/internal/filter"
)

func TestBuildFilterChain_WithTemplate(t *testing.T) {
	cfg := &config.Config{
		TemplatePattern: `(\d+)`,
		TemplateText:    "NUM:{{.Match}}",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == nil {
		t.Fatal("expected non-nil chain")
	}
}

func TestBuildFilterChain_TemplateEmptyPatternSkipped(t *testing.T) {
	cfg := &config.Config{
		TemplatePattern: "",
		TemplateText:    "{{.Line}}",
	}
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBuildFilterChain_TemplateEmptyTextSkipped(t *testing.T) {
	cfg := &config.Config{
		TemplatePattern: `\d+`,
		TemplateText:    "",
	}
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBuildFilterChain_TemplateInvalidPattern(t *testing.T) {
	cfg := &config.Config{
		TemplatePattern: `[bad`,
		TemplateText:    "{{.Line}}",
	}
	_, err := config.BuildFilterChain(cfg)
	if err == nil {
		t.Fatal("expected error for invalid pattern")
	}
	if !strings.Contains(err.Error(), "template filter") {
		t.Errorf("error should mention 'template filter', got: %v", err)
	}
}

func TestBuildFilterChain_TemplateAndRegex(t *testing.T) {
	cfg := &config.Config{
		Regex:           `ERROR`,
		TemplatePattern: `(\d+)`,
		TemplateText:    "code:{{.Match}}",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_ = chain
	_ = filter.NewChain // ensure filter package used
}
