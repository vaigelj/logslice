package config

import "flag"

func init() {
	registerTemplateFlags()
}

func registerTemplateFlags() {
	flag.StringVar(
		&defaultConfig.TemplatePattern,
		"template-pattern",
		"",
		"regex pattern to match for template transformation (requires --template-text)",
	)
	flag.StringVar(
		&defaultConfig.TemplateText,
		"template-text",
		"",
		"Go template applied to matching lines; variables: .Line, .Match, .Matches",
	)
}
