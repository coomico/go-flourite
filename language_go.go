package flourite

var goLang = []languagePattern{
	{expression: `package\s+[a-z]+\n`, patternType: metaModule, nearTop: true},
	{expression: `(import\s*\(\s*\n)|(import\s+"[a-z0-9/.]+")`, patternType: metaImport, nearTop: true},

	// error handling
	{expression: `if.+err\s*!=\s*nil.+\{`, patternType: keywordFunction},

	{expression: `fmt\.Print(f|ln)?\(.*\)`, patternType: keywordPrint},
	{expression: `func(\s+\w+\s*)?\(.*\).*\{`, patternType: keywordFunction},
	{expression: `\w+\s*:=\s*.+[^;\n]`, patternType: keywordVariable},
	{expression: `(var|const)\s+\w+\s+[\w*]+(\n|\s*=|$)`, patternType: keywordVariable},
	{expression: `(}\s*else\s*)?if[^()]+\{`, patternType: keywordControl},
	{expression: `nil`, patternType: keyword},

	// public access identifier
	{expression: `[a-z]+\.[A-Z]\w*`, patternType: macro},

	// single quote with multichar
	{expression: `'.{2,}'`, patternType: not},

	// avoiding C# confusion
	{expression: `Console\.(WriteLine|Write)(\s*)?\(`, patternType: not},
	{expression: `using\sSystem(\..*)?(;)?`, patternType: not},
	{expression: `(public|private|protected|internal)\s`, patternType: not},
}
