package flourite

var cLang = []languagePattern{
	{expression: `(char|long|int|float|double)\s+\w+\s*=?`, patternType: constantType},
	{expression: `malloc\(.+\)`, patternType: keywordFunction},
	{expression: `#include (<|")\w+\.h(>|")`, patternType: metaImport, nearTop: true},

	// pointer
	{expression: `(\w+)\s*\*\s*\w+`, patternType: keyword},

	// var declaration and/or initialisation
	{expression: `(\w+)\s+\w+(;|\s*=)`, patternType: macro},

	// array declaration
	{expression: `(\w+)\s+\w+\[.+\]`, patternType: keywordOther},

	// #define macro
	{expression: `#define\s+.+`, patternType: macro},

	{expression: `/NULL/`, patternType: constantNull},
	{expression: `void`, patternType: keywordOther},
	{expression: `(printf|puts)\s*\(.+\)`, patternType: keywordPrint},

	// `new` keyword from C++
	{expression: `new \w+`, patternType: not},

	// `new` keyword from Java
	{expression: `new [A-Z]\w*\s*\(.+\)`, patternType: not},

	// single quote with multichar
	{expression: `'.{2,}'`, patternType: not},

	// JS variable declaration
	{expression: `var\s+\w+\s*=?`, patternType: not},

	// avoiding Ruby confusion
	{expression: `/def\s+\w+\s*(\(.+\))?\s*\n/`, patternType: not},
	{expression: `/puts\s+("|').+("|')/`, patternType: not},

	// avoiding C# confusion
	{expression: `Console\.(WriteLine|Write)(\s*)?\(`, patternType: not},
	{expression: `(using\s)?System(\..*)?(;)?`, patternType: not},
	{expression: `(public\s)?((partial|static|delegate)\s)?(class\s)`, patternType: not},
	{expression: `(public|private|protected|internal)`, patternType: not},
	{
		expression:  `(new|this\s)?(List|IEnumerable)<(sbyte|byte|short|ushort|int|uint|long|ulong|float|double|decimal|bool|char|string)>`,
		patternType: not,
	},

	// avoiding Lua confusion
	{expression: `local\s(function|\w+)?`, patternType: not},

	// avoiding Dart confusion
	{expression: `^(void\s)?main\(\)\s(async\s)?{`, patternType: not},
}
