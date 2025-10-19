package flourite

var cs = []languagePattern{
	{expression: `using\sSystem(\..*)?(;)?`, patternType: metaImport},
	{expression: `Console\.(WriteLine|Write)(\s*)?\(`, patternType: keywordPrint},
	{expression: `Console\.ReadLine\(\)`, patternType: keywordOther},
	{expression: `(public\s)?((partial|static|delegate)\s)?class\s`, patternType: keyword},

	// modifiers
	{expression: `(extern|override|sealed|readonly|virtual|volatile)`, patternType: keywordOther},
	{expression: `namespace\s(.*)(\.(.*))?(\s{)?`, patternType: keyword},

	// regions
	{expression: `(#region(\s.*)?|#endregion\n)`, patternType: sectionScope},

	// functions
	{expression: `(public|private|protected|internal)\s`, patternType: keywordVisibility},

	{expression: `\bclass\s+\w+`, patternType: keyword},
	{expression: `(else )?if\s*\(.+\)`, patternType: keywordControl},
	{expression: `\bwhile\s+\(.+\)`, patternType: keywordControl},
	{
		expression:  `(const\s)?(sbyte|byte|short|ushort|int|uint|long|ulong|float|double|decimal|bool|char|string)(\[\])?\s(.*)\s=\s(.*);`,
		patternType: constantType,
	},

	// lists
	{
		expression:  `(new|this\s)?(List|IEnumerable)<(sbyte|byte|short|ushort|int|uint|long|ulong|float|double|decimal|bool|char|string)>`,
		patternType: constantDictionary,
	},

	{expression: `#define\s(.*)`, patternType: macro},

	// PascalCase
	{
		expression:  `\s([A-Z]([A-Z0-9]*[a-z][a-z0-9]*[A-Z]|[a-z0-9]*[A-Z][A-Z0-9]*[a-z])[A-Za-z0-9]*)\s=`,
		patternType: macro,
	},

	// avoiding Java confusion
	{expression: `(extends|throws|@Attribute)`, patternType: not},
	{expression: `System\.(in|out)\.\w+`, patternType: not},

	// avoiding Ruby confusion
	{expression: `\bmodule\s\S`, patternType: not},

	// avoiding Dart confusion
	{expression: `^\s*import\s("|')dart:\w+("|')`, patternType: not},
}
