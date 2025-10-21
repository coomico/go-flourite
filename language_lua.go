package flourite

var lua = []languagePattern{
	// multiline string
	{expression: `(\[\[.*\]\])`, patternType: constantString},

	{expression: `local\s([a-zA-Z0-9_]+)(\s*=)?`, patternType: keywordVariable},
	{expression: `(local\s)?function\s*([a-zA-Z0-9_]*)?\(\)`, patternType: keywordFunction},
	{expression: `for\s+([a-zA-Z]+)\s*=\s*([a-zA-Z0-9_]+),\s*([a-zA-Z0-9_]+)\s+do`, patternType: keywordControl},
	{expression: `while\s(.*)\sdo`, patternType: keywordControl},
	{
		expression:  `\s+(and|break|do|else|elseif|end|false|function|if|in|not|or|local|repeat|return|then|true|until|pairs|ipairs|in|yield)`,
		patternType: keywordOther,
	},
	{expression: `nil`, patternType: constantNull},

	// length operator
	{expression: `#([a-zA-Z_{}]+)`, patternType: keywordOperator},

	// metatables
	{expression: `((get|set)metatable|raw(get|set|equal))\(.*\)`, patternType: keywordOther},

	// metamethods
	{expression: `__(index|newindex|call|sub|mul|div|mod|pow|unm|eq|le|lt)`, patternType: keywordOther},

	// method invocation
	{expression: `(\(.+\)|([a-zA-Z_]+)):([a-zA-Z_])\(.*\)`, patternType: keywordOther},

	// array-like table
	{expression: `{\s*[^\s;,]+([;,]\s*[^\s;,]+)*,?\s*}`, patternType: constantArray},

	// map-like table
	{
		expression:  `{\s*([^\s;,=]+\s*=\s*[^\s;,=]+)(\s*[;,=]\s*[^\s;,=]+\s*=\s*[^\s;,=]+)*\s*,?\s*}`,
		patternType: constantDictionary,
	},

	// built-in methods
	{expression: `math\.(.*)\([0-9]*\)`, patternType: macro},
	{expression: `table\.(.*)\(.*\)`, patternType: macro},
	{expression: `io\.(.*)\(.*\)`, patternType: macro},

	// built-in functions
	{expression: `(require|dofile)\((.*)\)`, patternType: metaImport},
	{expression: `(pcall|xpcall|unpack|pack|coroutine)`, patternType: keywordOther},

	{expression: `--(\[\[)?.*`, patternType: commentLine},

	// rest args
	{expression: `\.\.\.`, patternType: keywordOther},

	{expression: `\bmodule\s*\(.*\)`, patternType: keywordOther},

	// invalid comments
	{expression: `(\/\/|\/\*)`, patternType: not},

	// avoiding C confusion
	{expression: `(#(include|define)|printf|\s+int\s+)`, patternType: not},

	// avoiding JavaScript confusion
	{expression: `\s+(let|const|var)\s+`, patternType: not},

	// avoiding PHP/Python confusion
	{expression: `\s+(echo|die|\$(.*))\s+`, patternType: not},

	// avoiding Python confusion
	{expression: `(def|len|from|import)`, patternType: not},

	// avoiding SQL confusion
	{expression: `(SELECT|FROM|INSERT|ALTER)`, patternType: not},

	// avoiding Ruby confusion
	{expression: `(puts)`, patternType: not},
	{expression: `\bmodule\s\S`, patternType: not},

	// avoiding Julia confusion
	{expression: `(([a-zA-Z0-9]+)::([a-zA-Z0-9]+)|using|(.*)!\(.*\)|(\|\|))`, patternType: not},
}
