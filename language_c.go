package flourite

import "regexp"

var cLang = []languagePattern{
	{
		expression:  regexp.MustCompile(`\b(unsigned|signed)?\s*(char|short|int|long|long\s+long|float|double)\s+(\*+)?\w+\s*(\[.+\])?(\s*=)?`),
		patternType: constantType,
	},
	{expression: regexp.MustCompile(`\b(m|c|re)alloc\s*\(.+\)`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`\bfree\s*\(.+\)`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`#include\s*[<"].+\.h[">]`), patternType: metaImport, nearTop: true},

	// pointer
	{expression: regexp.MustCompile(`(\w+)\s*\*\s*\w+`), patternType: keyword},

	// var declaration and/or initialisation
	{expression: regexp.MustCompile(`(\w+)\s+\w+(;|\s*=)`), patternType: macro},

	// array declaration
	{expression: regexp.MustCompile(`(\w+)\s+\w+\[.+\]`), patternType: keywordOther},

	// #define macro
	{expression: regexp.MustCompile(`#define\s+.+`), patternType: macro},

	{expression: regexp.MustCompile(`NULL`), patternType: constantNull},
	{expression: regexp.MustCompile(`void`), patternType: keywordOther},
	{expression: regexp.MustCompile(`(printf|puts)\s*\(.+\)`), patternType: keywordPrint},

	// `new` keyword from C++
	{expression: regexp.MustCompile(`new\s+\w+`), patternType: not},

	// `new` keyword from Java
	{expression: regexp.MustCompile(`new\s+[A-Z]\w*\s*\(.+\)`), patternType: not},

	// single quote with multichar
	{expression: regexp.MustCompile(`'.{2,}'`), patternType: not},

	// JS variable declaration
	{expression: regexp.MustCompile(`var\s+\w+\s*=?`), patternType: not},

	// avoiding Ruby confusion
	{expression: regexp.MustCompile(`def\s+\w+\s*(\(.+\))?\s*\n`), patternType: not},
	{expression: regexp.MustCompile(`puts\s+("|').+("|')`), patternType: not},

	// avoiding C# confusion
	{expression: regexp.MustCompile(`Console\.(WriteLine|Write)(\s*)?\(`), patternType: not},
	{expression: regexp.MustCompile(`(using\s)?System(\..*)?(;)?`), patternType: not},
	{expression: regexp.MustCompile(`(public\s)?((partial|static|delegate)\s)?(class\s)?`), patternType: not},
	{expression: regexp.MustCompile(`(public|private|protected|internal)`), patternType: not},
	{
		expression:  regexp.MustCompile(`(new|this\s)?(List|IEnumerable)<(sbyte|byte|short|ushort|int|uint|long|ulong|float|double|decimal|bool|char|string)>`),
		patternType: not,
	},

	// avoiding Lua confusion
	{expression: regexp.MustCompile(`local\s(function|\w+)?`), patternType: not},

	// avoiding Dart confusion
	{expression: regexp.MustCompile(`^(void\s)?main\s*\(\s*\)\s*(async\s*)?\{\s*$`), patternType: not},
	{expression: regexp.MustCompile(`\btypedef\s+\w+<.+>=.+Function`), patternType: not},
}
