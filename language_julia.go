package flourite

var julia = []languagePattern{
	{expression: `(using)\s\w+`, patternType: metaImport},
	{expression: `(bare\s)?module`, patternType: metaModule},

	// avoiding Python's import
	{expression: `from\s.+import\s.+`, patternType: not},

	// stdout / print line
	{expression: `println\(.*\)`, patternType: keywordPrint},

	{expression: `(.*)!\(.*\)`, patternType: macro},
	{expression: `for\s(\w+)\s(in|=)\s`, patternType: keywordControl},

	// function isn't followed by {
	{expression: `function\s\w+\(.*\)\s\{`, patternType: not},

	// while loop isn't followed by {
	{expression: `while\s+\(.+\)\n`, patternType: not},

	{expression: `end\n?`, patternType: keyword},

	// struct with <: annotation
	{expression: `struct\s(.*)\s<:\s`, patternType: keywordOther},

	{expression: `(::)?(Int|Uint)(8|16|32|64|128)`, patternType: keywordVariable},
	{expression: `[0-9]+im`, patternType: keyword},

	// avoiding Rust confusion
	{expression: `\{:\?\}`, patternType: not},
	{expression: `fn\smain\(\)`, patternType: not},

	// avoiding Ruby confusion
	{expression: `def\s+\w+\s*(\(.+\))?\s*\n`, patternType: not},
	{expression: `puts\s+("|').+("|')`, patternType: not},
	{expression: `class\s`, patternType: not},

	// avoiding Lua confusion
	{expression: `local\s(function|\w+)`, patternType: not},
	{expression: `\bmodule\(.*\)`, patternType: not},

	// avoiding Kotlin confusion
	{expression: `fun main\((.*)?\) \{`, patternType: not},
	{expression: `fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)\{`, patternType: not},
}
