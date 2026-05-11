package filter

import (
	"bytes"
	"fmt"
	"regexp"
	"text/template"
)

// TemplateFilter transforms matching lines using a Go text/template.
// The template receives a map with keys: "Line", "Match" (first submatch), "Matches" (all submatches).
type TemplateFilter struct {
	pattern  string
	re       *regexp.Regexp
	tmpl     *template.Template
	tmplText string
}

// NewTemplateFilter creates a TemplateFilter that applies tmplText to lines matching pattern.
func NewTemplateFilter(pattern, tmplText string) (*TemplateFilter, error) {
	if pattern == "" {
		return nil, ErrEmptyPattern
	}
	if tmplText == "" {
		return nil, fmt.Errorf("%w: template text must not be empty", ErrInvalidArgument)
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidPattern, err)
	}
	tmpl, err := template.New("line").Parse(tmplText)
	if err != nil {
		return nil, fmt.Errorf("%w: template parse error: %s", ErrInvalidArgument, err)
	}
	return &TemplateFilter{pattern: pattern, re: re, tmpl: tmpl, tmplText: tmplText}, nil
}

// Match always returns true; non-matching lines are passed through unchanged.
func (f *TemplateFilter) Match(line string) bool { return true }

// Transform applies the template to matching lines; non-matching lines are returned as-is.
func (f *TemplateFilter) Transform(line string) string {
	matches := f.re.FindStringSubmatch(line)
	if matches == nil {
		return line
	}
	first := ""
	if len(matches) > 1 {
		first = matches[1]
	}
	data := map[string]interface{}{
		"Line":    line,
		"Match":   first,
		"Matches": matches[1:],
	}
	var buf bytes.Buffer
	if err := f.tmpl.Execute(&buf, data); err != nil {
		return line
	}
	return buf.String()
}

// Pattern returns the compiled regex pattern string.
func (f *TemplateFilter) Pattern() string { return f.pattern }

// Template returns the raw template text.
func (f *TemplateFilter) Template() string { return f.tmplText }
