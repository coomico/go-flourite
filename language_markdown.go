package flourite

var markdown = []languagePattern{
	// heading
	{expression: `^#{2,6}\s.+`, patternType: keyword},

	// heading alternative syntax
	// [!] the original one isn't compatible with Golang regexp std library
	{expression: `^(?:=|-){2,}(?:>.+|$)`, patternType: metaModule},

	// images
	{expression: `!?\[.+\]\(.+\)`, patternType: keyword},

	// links
	{expression: `\[.+\]\[.+\]`, patternType: keyword},
	{expression: `^\[.+\]:\s?(<)?(http)?`, patternType: keyword},

	// blockquotes
	{expression: `^(> .*)+`, patternType: macro},

	// code blocks
	{expression: `^\x60\x60\x60([A-Za-z0-9#_]+)?$`, patternType: keyword},

	// front matter
	{expression: `^---$`, patternType: metaModule, nearTop: true},
}
