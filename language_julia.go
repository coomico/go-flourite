package flourite

import "regexp"

var julia = []languagePattern{
	{expression: regexp.MustCompile(`\busing\s+\w+`), patternType: metaImport},
	{expression: regexp.MustCompile(`\b(bare\s+)?module\b`), patternType: metaModule},
	{expression: regexp.MustCompile(`function\s+\w+\s*\([^)]*\)`), patternType: keywordFunction},

	// new{T} constructor
	{expression: regexp.MustCompile(`\bnew\{[^}]+\}`), patternType: keywordOther},

	// stdout / print line
	{expression: regexp.MustCompile(`println\(.*\)`), patternType: keywordPrint},

	{expression: regexp.MustCompile(`\w+!\s*\(`), patternType: macro},
	{expression: regexp.MustCompile(`for\s+\w+\s+(in|=)\s+`), patternType: keywordControl},
	{expression: regexp.MustCompile(`while\s+[^{]+`), patternType: keywordControl},
	{expression: regexp.MustCompile(`\breturn\s+`), patternType: keywordControl},

	{expression: regexp.MustCompile(`Union\{[^}]+\}`), patternType: keywordVariable},

	{expression: regexp.MustCompile(`end\n?`), patternType: keyword},

	// struct with <: annotation
	{expression: regexp.MustCompile(`(mutable\s+)?struct\s+\w+(\{[^}]+\})?`), patternType: keywordOther},

	{expression: regexp.MustCompile(`(::)?(Int|Uint)(8|16|32|64|128)`), patternType: keywordVariable},
	{expression: regexp.MustCompile(`\b(::)?Float(16|32|64)\b`), patternType: keywordVariable},
	{expression: regexp.MustCompile(`[0-9]+im\b`), patternType: keyword},
	{expression: regexp.MustCompile(`(?i)\bnothing\b`), patternType: constantNull},
	{expression: regexp.MustCompile(`\bSubString\s*\(`), patternType: keywordOther},
	{expression: regexp.MustCompile(`\blength\s*\(`), patternType: keywordOther},

	// array comprehension and indexing
	{expression: regexp.MustCompile(`\[\w+\s+for\s+\w+\s+in\s+`), patternType: keywordControl},
	{expression: regexp.MustCompile(`\w+\[(\d+|end|\d+:\d+|:|\d+:end|end-\d+)\]`), patternType: keywordOther},

	// avoiding Python's import
	{expression: regexp.MustCompile(`from\s.+import\s.+`), patternType: not},

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
