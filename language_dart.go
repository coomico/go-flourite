package flourite

var dart = []languagePattern{
	{
		expression:  `^\s*(const|final|var|dynamic|late)?\s*(int|double|String|bool|List<[A-Za-z \[\](),]+>|HashMap<[A-Za-z \[\](),]+>|Iterator<[A-Za-z \[\](),]+>|Set<[A-Za-z \[\](),]+>)?(\?)?\s\w+(\s=\s.+)?(;|,)$`,
		patternType: keywordVariable,
	},
	{expression: `\bstdout\.write\(.+\);`, patternType: keywordPrint},
	{expression: `\bprint\(.+\);`, patternType: keywordPrint},
	{expression: `^\s*import\s("|')dart:\w+("|')`, patternType: metaImport, nearTop: true},
	{expression: `^\s*import\s("|')package:\w+("|')`, patternType: metaImport, nearTop: true},
	{expression: `^\s*library\s\w+;`, patternType: metaModule, nearTop: true},
	{expression: `^\s*(void\s)?main\(\)\s(async\s)?\{`, patternType: keywordFunction},

	// function with type definition
	{
		expression:  `^\s*(List<[A-Za-z \[\](),]+>|HashMap<[A-Za-z \[\](),]+>|int|double|String|bool|void|Iterator<[A-Za-z \[\](),]+>|Set<[A-Za-z \[\](),]+>)\s\w+\(.+\)\s*\{$`,
		patternType: keywordFunction,
	},

	// arrow function
	{
		expression:  `^\s*(int|double|String|bool|List<[A-Za-z \[\](),]+>|HashMap<[A-Za-z \[\](),]+>|Iterator<[A-Za-z \[\](),]+>|Set<[A-Za-z \[\](),]+>)\s\w+\(.+\)\s=>`,
		patternType: keywordFunction,
	},

	{expression: `\bnew\s(List|Map|Iterator|HashMap|Set)<\w+>\(\);$`, patternType: keywordVariable},
	{
		expression:  `^(abstract\s)?class\s\w+\s(extends\s\w+\s)?(with\s\w+\s)?(implements\s\w+\s)?\{(\})?$`,
		patternType: keywordControl,
	},
	{expression: `\bget\s\w+=>\w+`, patternType: keywordControl},
	{expression: `^\s*@override$`, patternType: keywordControl},
	{expression: `\bset\s\w+\(.+\)`, patternType: keywordControl},
	{expression: `^\s*Future<\w+>\s\w+\(.*\)\s+async`, patternType: keywordControl},
	{expression: `^\s*await\sfor`, patternType: keywordControl},
	{expression: `^\s*typedef\s.+\s=`, patternType: keywordControl},

	// avoiding C confusion
	{expression: `\blong\s`, patternType: not},
	{expression: `\s*function\b`, patternType: not},

	// avoiding Java confusion
	{expression: `\bArrayList`, patternType: not},
}
