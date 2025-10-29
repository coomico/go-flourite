package flourite

import "regexp"

var goLang = []languagePattern{
	{expression: regexp.MustCompile(`package\s+[a-zA-Z_][a-zA-Z0-9_]*`), patternType: metaModule, nearTop: true},
	{
		expression: regexp.MustCompile(`^\s*import\s*(?:\(|[a-zA-Z_.]*\s*"[^"]+")`), patternType: metaImport,
		nearTop: true,
	},

	// error handling
	{expression: regexp.MustCompile(`if.+err\s*!=\s*nil.+\{`), patternType: keywordFunction},

	{expression: regexp.MustCompile(`\b(?:fmt|log|slog)\.\w+\(`), patternType: keywordPrint},
	{
		expression:  regexp.MustCompile(`func(?:\s+\([^)]*\))?(?:\s+\w+)?(?:\s*\[[^\]]+\])?\s*\([^)]*\)`),
		patternType: keywordFunction,
	},
	{expression: regexp.MustCompile(`\w+\s*:=\s*.+[^;\n]`), patternType: keywordVariable},
	{
		expression:  regexp.MustCompile(`^\s*(?:var|const)\s+(?:\(|[\w_]+(?:\s*,\s*[\w_]+)*(?:\s+[^=\n]+?)?\s*=?)`),
		patternType: keywordVariable,
	},
	{expression: regexp.MustCompile(`(}\s*else\s*)?if[^()]+\{`), patternType: keywordControl},
	{expression: regexp.MustCompile(`switch\s+.*\{`), patternType: keywordControl},
	{expression: regexp.MustCompile(`case\s+.*:`), patternType: keywordControl},
	{expression: regexp.MustCompile(`for\s*(\w+\s*:=\s*.*;.*|.*;.*|.*\{|.*range.*\{|\{)`), patternType: keywordControl},
	{
		expression:  regexp.MustCompile(`\b(defer|go|chan|select|fallthrough|break|continue|goto)\b`),
		patternType: keyword,
	},
	{expression: regexp.MustCompile(`\bnil\b`), patternType: constantNull},
	{expression: regexp.MustCompile(`\bmake\s*\(\s*(map\[|\[\]|chan\s)`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`^\s*type\s+\w+\s+(struct|interface|[\w*\[\]]+)\s*\{?`), patternType: constantType},

	// referencing
	{expression: regexp.MustCompile(`&\w+`), patternType: keywordOperator},

	// slice literal
	{expression: regexp.MustCompile(`\[\][\w*\[\]]+\{`), patternType: constantArray},

	{expression: regexp.MustCompile(`map\[\w+\][\w*&\[\]]+\{?`), patternType: constantDictionary},

	// public access identifier
	{expression: regexp.MustCompile(`[a-z]+\.[A-Z]\w*`), patternType: macro},

	// channel operations
	{expression: regexp.MustCompile(`(<-\s*chan|chan\s*<-|<-\s*\w+|\w+\s*<-)`), patternType: keywordOperator},

	// single quote with multichar
	{expression: regexp.MustCompile(`'.{2,}'`), patternType: not},

	// avoiding C# confusion
	{expression: regexp.MustCompile(`Console\.(WriteLine|Write)(\s*)?\(`), patternType: not},
	{expression: regexp.MustCompile(`using\sSystem(\..*)?(;)?`), patternType: not},
	{expression: regexp.MustCompile(`(public|private|protected|internal)\s`), patternType: not},
}
