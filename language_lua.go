package flourite

import "regexp"

var lua = []languagePattern{
	// multiline string
	{expression: regexp.MustCompile(`(\[\[.*\]\])`), patternType: constantString},

	{expression: regexp.MustCompile(`local\s([a-zA-Z0-9_]+)(\s*=)?`), patternType: keywordVariable},
	{expression: regexp.MustCompile(`(local\s)?function\s*([a-zA-Z0-9_]*)?\(\)`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`for\s+([a-zA-Z]+)\s*=\s*([a-zA-Z0-9_]+),\s*([a-zA-Z0-9_]+)\s+do`), patternType: keywordControl},
	{expression: regexp.MustCompile(`while\s(.*)\sdo`), patternType: keywordControl},
	{
		expression:  regexp.MustCompile(`\s+(and|break|do|else|elseif|end|false|function|if|in|not|or|local|repeat|return|then|true|until|pairs|ipairs|in|yield)`),
		patternType: keywordOther,
	},
	{expression: regexp.MustCompile(`nil`), patternType: constantNull},

	// length operator
	{expression: regexp.MustCompile(`#([a-zA-Z_{}]+)`), patternType: keywordOperator},

	// metatables
	{expression: regexp.MustCompile(`((get|set)metatable|raw(get|set|equal))\(.*\)`), patternType: keywordOther},

	// metamethods
	{expression: regexp.MustCompile(`__(index|newindex|call|sub|mul|div|mod|pow|unm|eq|le|lt)`), patternType: keywordOther},

	// method invocation
	{expression: regexp.MustCompile(`(\(.+\)|([a-zA-Z_]+)):([a-zA-Z_])\(.*\)`), patternType: keywordOther},

	// array-like table
	{expression: regexp.MustCompile(`\{\s*[^\s;,]+([;,]\s*[^\s;,]+)*,?\s*\}`), patternType: constantArray},

	// map-like table
	{
		expression:  regexp.MustCompile(`\{\s*([^\s;,=]+\s*=\s*[^\s;,=]+)(\s*[;,=]\s*[^\s;,=]+\s*=\s*[^\s;,=]+)*\s*,?\s*\}`),
		patternType: constantDictionary,
	},

	// built-in methods
	{expression: regexp.MustCompile(`math\.(.*)\([0-9]*\)`), patternType: macro},
	{expression: regexp.MustCompile(`table\.(.*)\(.*\)`), patternType: macro},
	{expression: regexp.MustCompile(`io\.(.*)\(.*\)`), patternType: macro},

	// built-in functions
	{expression: regexp.MustCompile(`(require|dofile)\((.*)\)`), patternType: metaImport},
	{expression: regexp.MustCompile(`(pcall|xpcall|unpack|pack|coroutine)`), patternType: keywordOther},

	{expression: regexp.MustCompile(`--(\[\[)?.*`), patternType: commentLine},

	// rest args
	{expression: regexp.MustCompile(`\.\.\.`), patternType: keywordOther},

	{expression: regexp.MustCompile(`\bmodule\s*\(.*\)`), patternType: keywordOther},

	// invalid comments
	{expression: regexp.MustCompile(`(//|/\*)`), patternType: not},

	// avoiding C confusion
	{expression: regexp.MustCompile(`(#(include|define)|printf|\s+int\s+)`), patternType: not},

	// avoiding JavaScript confusion
	{expression: regexp.MustCompile(`\s+(let|const|var)\s+`), patternType: not},

	// avoiding PHP/Python confusion
	{expression: regexp.MustCompile(`\s+(echo|die|\$(.*))\s+`), patternType: not},

	// avoiding Python confusion
	{expression: regexp.MustCompile(`(def|len|from|import)`), patternType: not},

	// avoiding SQL confusion
	{expression: regexp.MustCompile(`(SELECT|FROM|INSERT|ALTER)`), patternType: not},

	// avoiding Ruby confusion
	{expression: regexp.MustCompile(`(puts)`), patternType: not},
	{expression: regexp.MustCompile(`\bmodule\s\S`), patternType: not},

	// avoiding Julia confusion
	{expression: regexp.MustCompile(`(([a-zA-Z0-9]+)::([a-zA-Z0-9]+)|using|(.*)!\(.*\)|(\|\|))`), patternType: not},
}
