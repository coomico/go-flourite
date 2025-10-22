package flourite

import "regexp"

var goLang = []languagePattern{
	{expression: regexp.MustCompile(`package\s+[a-z]+\n`), patternType: metaModule, nearTop: true},
	{expression: regexp.MustCompile(`(import\s*\(\s*\n)|(import\s+"[a-z0-9/.]+")`), patternType: metaImport, nearTop: true},

	// error handling
	{expression: regexp.MustCompile(`if.+err\s*!=\s*nil.+\{`), patternType: keywordFunction},

	{expression: regexp.MustCompile(`fmt\.Print(f|ln)?\(.*\)`), patternType: keywordPrint},
	{expression: regexp.MustCompile(`func(\s+\w+\s*)?\(.*\).*\{`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`\w+\s*:=\s*.+[^;\n]`), patternType: keywordVariable},
	{expression: regexp.MustCompile(`(var|const)\s+\w+\s+[\w*]+(\n|\s*=|$)`), patternType: keywordVariable},
	{expression: regexp.MustCompile(`(}\s*else\s*)?if[^()]+\{`), patternType: keywordControl},
	{expression: regexp.MustCompile(`nil`), patternType: keyword},

	// public access identifier
	{expression: regexp.MustCompile(`[a-z]+\.[A-Z]\w*`), patternType: macro},

	// single quote with multichar
	{expression: regexp.MustCompile(`'.{2,}'`), patternType: not},

	// avoiding C# confusion
	{expression: regexp.MustCompile(`Console\.(WriteLine|Write)(\s*)?\(`), patternType: not},
	{expression: regexp.MustCompile(`using\sSystem(\..*)?(;)?`), patternType: not},
	{expression: regexp.MustCompile(`(public|private|protected|internal)\s`), patternType: not},
}
