package flourite

import "regexp"

var lua = []languagePattern{
	{expression: regexp.MustCompile(`\bmodule\s*\(\s*["'][^"']+["']\s*[,)]`), patternType: metaModule},

	// multiline string
	{expression: regexp.MustCompile(`(\[\[.*\]\])`), patternType: constantString},

	{expression: regexp.MustCompile(`\blocal\s+[a-zA-Z_][a-zA-Z0-9_]*(\s*=)?`), patternType: keywordVariable},
	{expression: regexp.MustCompile(`(local\s+)?function\s+([a-zA-Z0-9_.]+)?\s*\([^)]*\)`), patternType: keywordFunction},
	{
		expression:  regexp.MustCompile(`\bfor\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*=\s*[a-zA-Z0-9_]+\s*,\s*[a-zA-Z0-9_]+\s+do\b`),
		patternType: keywordControl,
	},
	{expression: regexp.MustCompile(`\bwhile\s+.+\s+do\b`), patternType: keywordControl},
	{expression: regexp.MustCompile(`\b(if|elseif)\s+.+\s+then\b`), patternType: keywordControl},

	{
		expression:  regexp.MustCompile(`\s+(and|break|do|else|end|false|function|in|not|or|local|repeat|return|then|true|until|pairs|ipairs|in|yield)`),
		patternType: keywordOther,
	},
	{expression: regexp.MustCompile(`nil`), patternType: constantNull},

	// length operator
	{expression: regexp.MustCompile(`#([a-zA-Z_{}]+)`), patternType: keywordOperator},

	// metatables
	{expression: regexp.MustCompile(`((get|set)metatable|raw(get|set|equal))\(.*\)`), patternType: keywordOther},

	// metamethods
	{
		expression:  regexp.MustCompile(`__(index|newindex|call|add|sub|mul|div|mod|pow|unm|concat|len|eq|lt|le|tostring|metatable|gc|mode)`),
		patternType: keywordOther,
	},

	// method invocation
	{expression: regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*:[a-zA-Z_][a-zA-Z0-9_]*\s*\(`), patternType: keywordOther},

	// array-like table
	{expression: regexp.MustCompile(`\{\s*[^\s;,]+([;,]\s*[^\s;,]+)*,?\s*\}`), patternType: constantArray},

	// map-like table
	{
		expression:  regexp.MustCompile(`\{\s*([^\s;,=]+\s*=\s*[^\s;,=]+)(\s*[;,=]\s*[^\s;,=]+\s*=\s*[^\s;,=]+)*\s*,?\s*\}`),
		patternType: constantDictionary,
	},

	// built-in methods
	{expression: regexp.MustCompile(`\bmath\.[a-zA-Z_][a-zA-Z0-9_]*\s*\(`), patternType: macro},
	{expression: regexp.MustCompile(`\btable\.[a-zA-Z_][a-zA-Z0-9_]*\s*\(`), patternType: macro},
	{expression: regexp.MustCompile(`\bio\.[a-zA-Z_][a-zA-Z0-9_]*\s*\(`), patternType: macro},
	{expression: regexp.MustCompile(`\bstring\.[a-zA-Z_][a-zA-Z0-9_]*\s*\(`), patternType: macro},
	{expression: regexp.MustCompile(`\bos\.[a-zA-Z_][a-zA-Z0-9_]*\s*\(`), patternType: macro},

	// built-in functions
	{expression: regexp.MustCompile(`\b(require|dofile|loadfile|load)\s*\(`), patternType: metaImport},
	{expression: regexp.MustCompile(`(pcall|xpcall|unpack|pack|coroutine)`), patternType: keywordOther},

	{expression: regexp.MustCompile(`--(\[\[)?.*`), patternType: commentLine},

	// rest args
	{expression: regexp.MustCompile(`\.\.\.`), patternType: keywordOther},

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
