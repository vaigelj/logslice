package config

import (
	"fmt"

	"github.com/yourorg/logslice/internal/filter"
)

// addRegexReplaceFilter appends a RegexReplaceFilter to the chain when
// cfg.RegexReplacePattern is non-empty. cfg.RegexReplaceWith may be empty
// (which deletes all matches).
func addRegexReplaceFilter(cfg *Config, chain *filter.Chain) error {
	if cfg.RegexReplacePattern == "" {
		return nil
	}
	f, err := filter.NewRegexReplaceFilter(cfg.RegexReplacePattern, cfg.RegexReplaceWith)
	if err != nil {
		return fmt.Errorf("regex-replace filter: %w", err)
	}
	chain.Add(f)
	return nil
}

func init() {
	registerRegexReplaceFlags()
}

func registerRegexReplaceFlags() {
	// --regex-replace-pattern and --regex-replace-with are registered here so
	// they appear in the global flag set used by config.Parse.
	globalFlags.StringVar(
		&defaultConfig.RegexReplacePattern,
		"regex-replace-pattern",
		"",
		"replace all occurrences of this regex in each line",
	)
	globalFlags.StringVar(
		&defaultConfig.RegexReplaceWith,
		"regex-replace-with",
		"",
		"replacement string for --regex-replace-pattern (supports capture groups)",
	)
}
