package flourite

var ruby = []languagePattern{
	{expression: `(require|include)\s+'\w+(\.rb)?'`, patternType: metaImport, nearTop: true},
	{expression: `def\s+\w+\s*(\(.+\))?\s*\n`, patternType: keywordFunction},

	// instance variable
	{expression: `@\w+`, patternType: keywordOther},

	// boolean property
	{expression: `\.\w+\?`, patternType: constantBoolean},

	// Ruby print
	{expression: `puts\s+("|').+("|')`, patternType: keywordPrint},

	// inheriting class
	{expression: `class [A-Z]\w*\s*<\s*([A-Z]\w*(::)?)+`, patternType: keyword},

	{expression: `attr_accessor\s+(:\w+(,\s*)?)+`, patternType: keywordFunction},
	{expression: `\w+\.new\s+`, patternType: keyword},
	{expression: `elsif`, patternType: keywordControl},
	{expression: `\bmodule\s\S`, patternType: keywordOther},
	{expression: `\bBEGIN\s\{.*\}`, patternType: keywordOther},
	{expression: `\bEND\s\{.*\}`, patternType: keywordOther},
	{expression: `do\s*\|\w+(,\s*\w+)*\|`, patternType: keywordControl},
	{expression: `for (\w+|\(?\w+,\s*\w+\)?) in (.+)`, patternType: keywordControl},
	{expression: `nil`, patternType: constantNull},

	// scope operator
	{expression: `[A-Z]\w*::[A-Z]\w*`, patternType: macro},

	// termination code blocks `end`
	{expression: `^\s*end\b\s*(?:#.*)?$`, patternType: keyword},
}
