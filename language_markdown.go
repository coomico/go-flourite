package flourite

import "regexp"

var markdown = []languagePattern{
	// heading
	{expression: regexp.MustCompile(`^#{2,6}\s.+`), patternType: keyword},

	// heading alternative syntax
	// [!] the original one isn't compatible with Golang regexp std library
	{expression: regexp.MustCompile(`^(?:=|-){2,}(?:>.+|$)`), patternType: metaModule},

	// images
	{expression: regexp.MustCompile(`!?\[.+\]\(.+\)`), patternType: keyword},

	// links
	{expression: regexp.MustCompile(`\[.+\]\[.+\]`), patternType: keyword},
	{expression: regexp.MustCompile(`^\[.+\]:\s?(<)?(http)?`), patternType: keyword},

	// blockquotes
	{expression: regexp.MustCompile(`^(> .*)+`), patternType: macro},

	// code blocks
	{expression: regexp.MustCompile(`^\x60\x60\x60([A-Za-z0-9#_]+)?$`), patternType: keyword},

	// front matter
	{expression: regexp.MustCompile(`^---$`), patternType: metaModule, nearTop: true},
}
