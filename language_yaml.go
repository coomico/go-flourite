package flourite

import "regexp"

var yaml = []languagePattern{
	// regular key: value
	{
		expression:  regexp.MustCompile(`^\s*[A-Za-z0-9_][A-Za-z0-9_\-. ]*:\s*["']?[^;\r\n]*["']?$`),
		patternType: keywordVariable, // take precedence over Markdown front matter
	},

	// regular array - key: value
	{expression: regexp.MustCompile(`^\s*-\s+[A-Za-z0-9_. ]+:\s*["']?.*["']?$`), patternType: keyword},

	// regular array - value
	{expression: regexp.MustCompile(`^\s*-\s+\S+.*$`), patternType: keyword},

	// binary tag
	{expression: regexp.MustCompile(`^\s*[A-Za-z0-9_. ]+:\s*!!binary\s*\|?\s*$`), patternType: constantType},

	// literal multiline block
	{expression: regexp.MustCompile(`^\s*[A-Za-z0-9_. ]+:\s*\|$`), patternType: keyword},

	// folded multiline style
	{expression: regexp.MustCompile(`^\s*[A-Za-z0-9_. ]+:\s*>$`), patternType: keyword},

	// set types
	{expression: regexp.MustCompile(`^\s*\?\s+.*$`), patternType: keyword},

	// complex key / multiline key
	{expression: regexp.MustCompile(`^\s*\?\s*\|$`), patternType: constantType},

	// merge key
	{expression: regexp.MustCompile(`^\s*<<:\s*\*.*$`), patternType: constantType},

	// avoiding CSS confusion
	{expression: regexp.MustCompile(`^( )*([A-Za-z0-9_. ]+):(.*)?( )?{$`), patternType: not},
	{expression: regexp.MustCompile(`^( )*([A-Za-z0-9_. ]+):(.*)?( )?,$`), patternType: not},
}
