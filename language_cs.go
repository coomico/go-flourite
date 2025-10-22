package flourite

import "regexp"

var cs = []languagePattern{
	{expression: regexp.MustCompile(`using\sSystem(\..*)?(;)?`), patternType: metaImport},
	{expression: regexp.MustCompile(`Console\.(WriteLine|Write)(\s*)?\(`), patternType: keywordPrint},
	{expression: regexp.MustCompile(`Console\.ReadLine\(\)`), patternType: keywordOther},
	{expression: regexp.MustCompile(`(public\s)?((partial|static|delegate)\s)?class\s`), patternType: keyword},

	// modifiers
	{expression: regexp.MustCompile(`(extern|override|sealed|readonly|virtual|volatile)`), patternType: keywordOther},
	{expression: regexp.MustCompile(`namespace\s(.*)(\.(.*))?(\s{)?`), patternType: keyword},

	// regions
	{expression: regexp.MustCompile(`(#region(\s.*)?|#endregion\n)`), patternType: sectionScope},

	// functions
	{expression: regexp.MustCompile(`(public|private|protected|internal)\s`), patternType: keywordVisibility},

	{expression: regexp.MustCompile(`\bclass\s+\w+`), patternType: keyword},
	{expression: regexp.MustCompile(`(else )?if\s*\(.+\)`), patternType: keywordControl},
	{expression: regexp.MustCompile(`\bwhile\s+\(.+\)`), patternType: keywordControl},
	{
		expression:  regexp.MustCompile(`(const\s)?(sbyte|byte|short|ushort|int|uint|long|ulong|float|double|decimal|bool|char|string)(\[\])?\s(.*)\s=\s(.*);`),
		patternType: constantType,
	},

	// lists
	{
		expression:  regexp.MustCompile(`(new|this\s)?(List|IEnumerable)<(sbyte|byte|short|ushort|int|uint|long|ulong|float|double|decimal|bool|char|string)>`),
		patternType: constantDictionary,
	},

	{expression: regexp.MustCompile(`#define\s(.*)`), patternType: macro},

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
