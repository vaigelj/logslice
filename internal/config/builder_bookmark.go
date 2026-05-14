package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addBookmarkFilter appends a BookmarkFilter to the chain when both
// --bookmark-start and --bookmark-end flags are provided.
func addBookmarkFilter(cfg *Config, chain *filter.Chain) error {
	start := cfg.BookmarkStart
	end := cfg.BookmarkEnd

	if start == "" && end == "" {
		return nil
	}
	if start == "" || end == "" {
		return fmt.Errorf("--bookmark-start and --bookmark-end must both be set")
	}

	f, err := filter.NewBookmarkFilter(start, end)
	if err != nil {
		return fmt.Errorf("bookmark filter: %w", err)
	}
	chain.Add(f)
	return nil
}
