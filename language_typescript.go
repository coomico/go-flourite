package flourite

var typescript = []languagePattern{
	{
		expression:  `(Readonly<|ReadonlyArray<|Array<|Record<|Pick<|Omit<|Exclude<|Extract<)`,
		patternType: constantDictionary,
	},
	{expression: `\w+\[\]`, patternType: keywordOther},
	{expression: `:\s*\w+\s*<`, patternType: keywordControl},
	{
		expression:  `\??\s*:\s*(any|string|boolean|undefined|null|number|void|never|symbol|bigint)\s*`,
		patternType: constantType,
	},
	{
		expression:  `(let|const|var)*\s*\w*:\s*(any|string|boolean|undefined|null|number|void|never|symbol|bigint)\s*=`,
		patternType: constantType,
	},
	{expression: `console\.log\s*\(`, patternType: keywordPrint},
	{expression: `interface\s*(\w+)\s*\{`, patternType: constantType},
	{expression: `enum\s*\w+\s*=`, patternType: constantType},
	{expression: `type\s*(\w+)\s*=`, patternType: constantType},
	{expression: `function\s+\w+\(.*\):\s\w+\s*\{`, patternType: keywordFunction},

	// arrow function
	{expression: `\(.*\):\s\w+\s*=>\s*\{`, patternType: keywordFunction},

	{expression: `(typeof|declare)\s+`, patternType: keyword},
	{expression: `\s+as\s+`, patternType: keyword},

	// Rust types
	{expression: `usize`, patternType: not},

	// Kotlin
	{expression: `Array<String>`, patternType: not},
}
