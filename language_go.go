package flourite

var goLang = []LanguagePattern{
	{Pattern: `package\s+[a-z]+\n`, Type: MetaModule, NearTop: true},
	{Pattern: `(import\s*\(\s*\n)|(import\s+"[a-z0-9/.]+")`, Type: MetaImport, NearTop: true},

	// error handling
	{Pattern: `if.+err\s*!=\s*nil.+{`, Type: KeywordFunction},

	{Pattern: `fmt\.Print(f|ln)?\(.*\)`, Type: KeywordPrint},
	{Pattern: `func(\s+\w+\s*)?\(.*\).*{`, Type: KeywordFunction},
	{Pattern: `\w+\s*:=\s*.+[^;\n]`, Type: KeywordVariable},
	{Pattern: `(var|const)\s+\w+\s+[\w*]+(\n|\s*=|$)`, Type: KeywordVariable},
	{Pattern: `(}\s*else\s*)?if[^()]+{`, Type: KeywordControl},
	{Pattern: `nil`, Type: Keyword},

	// public access identifier
	{Pattern: `[a-z]+\.[A-Z]\w*`, Type: Macro},

	// single quote with multichar
	{Pattern: `'.{2,}'`, Type: Not},

	// avoiding C# confusion
	{Pattern: `Console\.(WriteLine|Write)(\s*)?\(`, Type: Not},
	{Pattern: `using\sSystem(\..*)?(;)?`, Type: Not},
	{Pattern: `(public|private|protected|internal)\s`, Type: Not},
}
