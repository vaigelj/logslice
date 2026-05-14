package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
)

func TestBuildFilterChain_WithJSONPresent(t *testing.T) {
	cfg := &config.Config{JSONPresentField: "level"}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !chain.Match(`{"level":"info"}`) {
		t.Error("expected match when field is present")
	}
	if chain.Match(`{"msg":"no level"}`) {
		t.Error("expected no match when field is absent")
	}
}

func TestBuildFilterChain_WithJSONAbsent(t *testing.T) {
	cfg := &config.Config{JSONAbsentField: "debug"}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !chain.Match(`{"level":"info"}`) {
		t.Error("expected match when field is absent")
	}
	if chain.Match(`{"debug":true}`) {
		t.Error("expected no match when field is present")
	}
}

func TestBuildFilterChain_JSONPresentEmptySkipped(t *testing.T) {
	cfg := &config.Config{}
	_, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error when both fields are empty: %v", err)
	}
}

func TestBuildFilterChain_JSONPresentNestedField(t *testing.T) {
	cfg := &config.Config{JSONPresentField: "meta.user"}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !chain.Match(`{"meta":{"user":"alice"}}`) {
		t.Error("expected match for nested field")
	}
	if chain.Match(`{"meta":{"host":"srv1"}}`) {
		t.Error("expected no match when nested field is absent")
	}
}

func TestBuildFilterChain_JSONPresentAndRegex(t *testing.T) {
	cfg := &config.Config{
		JSONPresentField: "level",
		Pattern:          "info",
	}
	chain, err := config.BuildFilterChain(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !chain.Match(`{"level":"info","msg":"started"}`) {
		t.Error("expected match when both conditions are satisfied")
	}
	if chain.Match(`{"level":"warn","msg":"started"}`) {
		t.Error("expected no match when regex does not match")
	}
}
