package flourite

import "regexp"

var markdown = []languagePattern{
	// horizontal rules
	{expression: regexp.MustCompile(`^[\*\-_]{3,}\s*$`), patternType: macro},

	// YAML front matter keys
	{expression: regexp.MustCompile(`^\s*[a-zA-Z_][a-zA-Z0-9_-]*:\s*.+$`), patternType: keyword, nearTop: true},

	// heading
	{expression: regexp.MustCompile(`^#{1,6}\s+.+`), patternType: keyword},

	// heading alternative syntax
	// [!] the original one isn't compatible with Golang regexp std library
	{expression: regexp.MustCompile(`^(?:=|-){2,}(?:>.+|$)`), patternType: metaModule},

	// images
	{expression: regexp.MustCompile(`!?\[.+\]\(.+\)`), patternType: keyword},

	// links
	{expression: regexp.MustCompile(`\[.+\]\[.+\]`), patternType: keywordVariable},
	{expression: regexp.MustCompile(`^\[[^\]]+\]:\s*<?[a-zA-Z0-9:/.#_-]+>?`), patternType: keywordVariable},

	// blockquotes
	{expression: regexp.MustCompile(`^(> .*)+`), patternType: macro},

	// code blocks
	{expression: regexp.MustCompile(`^\x60{3}[A-Za-z0-9#_+-]*$`), patternType: keyword},

	// inline code
	{expression: regexp.MustCompile(`\x60[^\x60\n]+\x60`), patternType: keyword},

	// unordered list
	{expression: regexp.MustCompile(`^\s*[-*+]\s+.+`), patternType: keyword},

	// ordered list
	{expression: regexp.MustCompile(`^\s*\d+\.\s+.+`), patternType: keyword},

	// bold and italic emphasis
	{expression: regexp.MustCompile(`\*{1,3}[^\s\*]([^\*]*[^\s\*])?\*{1,3}`), patternType: keyword},
	{expression: regexp.MustCompile(`(?:^|[^\w])_{1,3}[^\s_]([^_]*[^\s_])?_{1,3}(?:$|[^\w])`), patternType: keyword},
}
