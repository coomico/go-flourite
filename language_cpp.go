package flourite

var cpp = []languagePattern{
	{expression: `(char|long|int|float|double)\s+\w+\s*=?`, patternType: constantType},
	{expression: `#include\s*(<|")\w+(\.h)?(>|")`, patternType: metaImport},
	{expression: `using\s+namespace\s+.+\s*;`, patternType: keyword},
	{expression: `template\s*<.*>`, patternType: keyword},
	{expression: `std::\w+`, patternType: keywordOther},
	{expression: `(cout|cin|endl)`, patternType: keywordPrint},
	{expression: `(public|protected|private):`, patternType: keywordVisibility},
	{expression: `nullptr`, patternType: keyword},
	{expression: `new \w+(\(.*\))?`, patternType: keyword},
	{expression: `#define\s+.+`, patternType: macro},

	// template usage
	{expression: `\w+<\w+>`, patternType: keywordOther},

	{expression: `class\s+\w+`, patternType: keyword},
	{expression: `void`, patternType: keyword},
	{expression: `(else )?if\s*\(.+\)`, patternType: keywordControl},
	{expression: `while\s+\(.+\)`, patternType: keywordControl},

	// scope operator
	{expression: `\w*::\w+`, patternType: macro},

	// single quote with multichar
	{expression: `'.{2,}'`, patternType: not},

	// Java List/ArrayList
	{expression: `(List<\w+>|ArrayList<\w*>\s*\(.*\))(\s+[\w]+|;)`, patternType: not},

	// avoiding Ruby confusion
	{expression: `def\s+\w+\s*(\(.+\))?\s*\n`, patternType: not},
	{expression: `puts\s+("|').+("|')`, patternType: not},
	{expression: `\bmodule\s\S`, patternType: not},

	// avoiding C# confusion
	{expression: `Console\.(WriteLine|Write)(\s*)?\(`, patternType: not},
	{expression: `(using\s)?System(\..*)?(;)?`, patternType: not},
	{expression: `static\s+\S+\s+Main\(.*\)`, patternType: not},
	{expression: `(public|private|protected|internal)\s`, patternType: not},

	// avoiding Kotlin confusion
	{expression: `fun main\((.*)?\) {`, patternType: not},
	{
		expression:  `(inline|private|public|protected|override|operator(\s+))?fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)({|=)`,
		patternType: not,
	},
	{expression: `(const)?(\s+)?val(\s+)(.*)(:(\s)(.*)(\?)?)?(\s+)=(\s+)`, patternType: not},

	// avoiding Dart confusion
	{expression: `^(void\s)?main\(\)\s(async\s)?{`, patternType: not},
}
