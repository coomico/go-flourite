package flourite

import "regexp"

var julia = []languagePattern{
	{expression: regexp.MustCompile(`(using)\s\w+`), patternType: metaImport},
	{expression: regexp.MustCompile(`(bare\s)?module`), patternType: metaModule},

	// avoiding Python's import
	{expression: regexp.MustCompile(`from\s.+import\s.+`), patternType: not},

	// stdout / print line
	{expression: regexp.MustCompile(`println\(.*\)`), patternType: keywordPrint},

	{expression: regexp.MustCompile(`(.*)!\(.*\)`), patternType: macro},
	{expression: regexp.MustCompile(`for\s(\w+)\s(in|=)\s`), patternType: keywordControl},

	// function isn't followed by {
	{expression: regexp.MustCompile(`function\s\w+\(.*\)\s\{`), patternType: not},

	// while loop isn't followed by {
	{expression: regexp.MustCompile(`while\s+\(.+\)\n`), patternType: not},

	{expression: regexp.MustCompile(`end\n?`), patternType: keyword},

	// struct with <: annotation
	{expression: regexp.MustCompile(`struct\s(.*)\s<:\s`), patternType: keywordOther},

	{expression: regexp.MustCompile(`(::)?(Int|Uint)(8|16|32|64|128)`), patternType: keywordVariable},
	{expression: regexp.MustCompile(`[0-9]+im`), patternType: keyword},

	// avoiding Rust confusion
	{expression: regexp.MustCompile(`\{:\?\}`), patternType: not},
	{expression: regexp.MustCompile(`fn\smain\(\)`), patternType: not},

	// avoiding Ruby confusion
	{expression: regexp.MustCompile(`def\s+\w+\s*(\(.+\))?\s*\n`), patternType: not},
	{expression: regexp.MustCompile(`puts\s+("|').+("|')`), patternType: not},
	{expression: regexp.MustCompile(`class\s`), patternType: not},

	// avoiding Lua confusion
	{expression: regexp.MustCompile(`local\s(function|\w+)`), patternType: not},
	{expression: regexp.MustCompile(`\bmodule\(.*\)`), patternType: not},

	// avoiding Kotlin confusion
	{expression: regexp.MustCompile(`fun main\((.*)?\) \{`), patternType: not},
	{expression: regexp.MustCompile(`fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)\{`), patternType: not},
}
