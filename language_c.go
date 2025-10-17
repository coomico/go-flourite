package flourite

var c = []LanguagePattern{
	{Pattern: `(char|long|int|float|double)\s+\w+\s*=?`, Type: ConstantType},
	{Pattern: `malloc\(.+\)`, Type: KeywordFunction},
	{Pattern: `#include (<|")\w+\.h(>|")`, Type: MetaImport, NearTop: true},

	// pointer
	{Pattern: `(\w+)\s*\*\s*\w+`, Type: Keyword},

	// var declaration and/or initialisation
	{Pattern: `(\w+)\s+\w+(;|\s*=)`, Type: Macro},

	// array declaration
	{Pattern: `(\w+)\s+\w+\[.+\]`, Type: KeywordOther},

	// #define macro
	{Pattern: `#define\s+.+`, Type: Macro},

	{Pattern: `/NULL/`, Type: ConstantNull},
	{Pattern: `void`, Type: KeywordOther},
	{Pattern: `(printf|puts)\s*\(.+\)`, Type: KeywordPrint},

	// `new` keyword from C++
	{Pattern: `new \w+`, Type: Not},

	// `new` keyword from Java
	{Pattern: `new [A-Z]\w*\s*\(.+\)`, Type: Not},

	// single quote with multichar
	{Pattern: `'.{2,}'`, Type: Not},

	// JS variable declaration
	{Pattern: `var\s+\w+\s*=?`, Type: Not},

	// avoiding Ruby confusion
	{Pattern: `/def\s+\w+\s*(\(.+\))?\s*\n/`, Type: Not},
	{Pattern: `/puts\s+("|').+("|')/`, Type: Not},

	// avoiding C# confusion
	{Pattern: `Console\.(WriteLine|Write)(\s*)?\(`, Type: Not},
	{Pattern: `(using\s)?System(\..*)?(;)?`, Type: Not},
	{Pattern: `(public\s)?((partial|static|delegate)\s)?(class\s)`, Type: Not},
	{Pattern: `(public|private|protected|internal)`, Type: Not},
	{Pattern: `(new|this\s)?(List|IEnumerable)<(sbyte|byte|short|ushort|int|uint|long|ulong|float|double|decimal|bool|char|string)>`, Type: Not},

	// avoiding Lua confusion
	{Pattern: `local\s(function|\w+)?`, Type: Not},

	// avoiding Dart confusion
	{Pattern: `^(void\s)?main\(\)\s(async\s)?{`, Type: Not},
}
