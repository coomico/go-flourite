package flourite

import "regexp"

var cs = []languagePattern{
	{expression: regexp.MustCompile(`^\s*using\s+System(\.\w+)*(;|\s)`), patternType: metaImport},
	{expression: regexp.MustCompile(`Console\.(WriteLine|Write)(\s*)?\(`), patternType: keywordPrint},
	{expression: regexp.MustCompile(`Console\.(ReadLine|Read)\(\)`), patternType: keywordOther},
	{
		expression:  regexp.MustCompile(`(public|private|protected|internal)?\s*(partial|static|abstract|sealed)?\s*class\s+\w+`),
		patternType: keyword,
	},

	// modifiers
	{
		expression:  regexp.MustCompile(`\b(extern|override|sealed|readonly|virtual|volatile|partial)\b`),
		patternType: keywordOther,
	},

	{expression: regexp.MustCompile(`^\s*namespace\s+[\w]+(\.[\w]+)*\s*$`), patternType: keyword},

	// regions
	{expression: regexp.MustCompile(`^\s*#(region|endregion)(\s.*)?`), patternType: sectionScope},

	{expression: regexp.MustCompile(`static\s+void\s+\w+\s*\(`), patternType: keyword},

	// functions
	{expression: regexp.MustCompile(`\b(public|private|protected|internal)\s+`), patternType: keywordVisibility},

	{expression: regexp.MustCompile(`\bclass\s+\w+`), patternType: keyword},
	{expression: regexp.MustCompile(`(else\s+)?if\s*\(.+\)`), patternType: keywordControl},
	{expression: regexp.MustCompile(`\bwhile\s*\([^)]+\)\s*$`), patternType: keywordControl},
	{
		expression:  regexp.MustCompile(`\b(const\s+)?(sbyte|byte|short|ushort|int|uint|long|ulong|float|double|decimal|bool|char|string|var|object|dynamic)(\[\])?\s+\w+\s*(=|;)`),
		patternType: constantType,
	},

	// lists
	{
		expression:  regexp.MustCompile(`\b(new\s+)?(List|Dictionary|HashSet|Queue|Stack|IEnumerable|ICollection)<\w+>`),
		patternType: constantDictionary,
	},

	// array initialization style
	{
		expression:  regexp.MustCompile(`\b(int|string|char|byte|bool|double|float)\[\]\s+\w+\s*=\s*\{`),
		patternType: constantArray,
	},

	{expression: regexp.MustCompile(`#define\s.*`), patternType: macro},

	// PascalCase
	{
		expression:  regexp.MustCompile(`\s([A-Z]([A-Z0-9]*[a-z][a-z0-9]*[A-Z]|[a-z0-9]*[A-Z][A-Z0-9]*[a-z])[A-Za-z0-9]*)\s=`),
		patternType: macro,
	},

	// avoiding Java confusion
	{expression: regexp.MustCompile(`(extends|throws|@Attribute)`), patternType: not},
	{expression: regexp.MustCompile(`System\.(in|out)\.\w+`), patternType: not},

	// avoiding Ruby confusion
	{expression: regexp.MustCompile(`\bmodule\s\S`), patternType: not},

	// avoiding Dart confusion
	{expression: regexp.MustCompile(`^\s*import\s("|')dart:\w+("|')`), patternType: not},
}
