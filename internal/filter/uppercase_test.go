package filter_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/filter"
)

func TestNewUppercaseFilter_InvalidMode(t *testing.T) {
	_, err := filter.NewUppercaseFilter(filter.uppercaseMode(99))
	if err == nil {
		t.Fatal("expected error for invalid mode, got nil")
	}
}

func TestNewUppercaseFilter_ValidModes(t *testing.T) {
	modes := []filter.uppercaseMode{
		filter.UppercaseFull,
		filter.UppercaseLower,
		filter.UppercaseTitle,
	}
	for _, m := range modes {
		f, err := filter.NewUppercaseFilter(m)
		if err != nil {
			t.Fatalf("unexpected error for mode %d: %v", m, err)
		}
		if f == nil {
			t.Fatalf("expected non-nil filter for mode %d", m)
		}
	}
}

func TestUppercaseFilter_Match_AlwaysTrue(t *testing.T) {
	f, _ := filter.NewUppercaseFilter(filter.UppercaseFull)
	lines := []string{"", "hello", "WORLD", "  spaces  "}
	for _, line := range lines {
		if !f.Match(line) {
			t.Errorf("expected Match to return true for %q", line)
		}
	}
}

func TestUppercaseFilter_Transform_Full(t *testing.T) {
	f, _ := filter.NewUppercaseFilter(filter.UppercaseFull)
	got := f.Transform("hello world")
	want := "HELLO WORLD"
	if got != want {
		t.Errorf("Transform() = %q, want %q", got, want)
	}
}

func TestUppercaseFilter_Transform_Lower(t *testing.T) {
	f, _ := filter.NewUppercaseFilter(filter.UppercaseLower)
	got := f.Transform("HELLO WORLD")
	want := "hello world"
	if got != want {
		t.Errorf("Transform() = %q, want %q", got, want)
	}
}

func TestUppercaseFilter_Transform_Title(t *testing.T) {
	f, _ := filter.NewUppercaseFilter(filter.UppercaseTitle)
	got := f.Transform("hello world")
	want := "HELLO WORLD" // strings.ToTitle uppercases all letters
	if got != want {
		t.Errorf("Transform() = %q, want %q", got, want)
	}
}

func TestUppercaseFilter_Mode_Accessor(t *testing.T) {
	f, _ := filter.NewUppercaseFilter(filter.UppercaseLower)
	if f.Mode() != filter.UppercaseLower {
		t.Errorf("Mode() = %d, want %d", f.Mode(), filter.UppercaseLower)
	}
}

func TestUppercaseFilter_Transform_EmptyLine(t *testing.T) {
	f, _ := filter.NewUppercaseFilter(filter.UppercaseFull)
	got := f.Transform("")
	if got != "" {
		t.Errorf("Transform(\"\") = %q, want empty string", got)
	}
}
