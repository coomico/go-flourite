package flourite

var js = []languagePattern{
	{expression: `undefined`, patternType: keyword},
	{expression: `window\.`, patternType: keyword},
	{expression: `console\.log\s*\(`, patternType: keywordPrint},
	{expression: `(var|const|let)\s+\w+\s*=?`, patternType: keywordVariable},

	// array/object declaration
	{expression: `(('|").+('|")\s*|\w+):\s*[{[]`, patternType: constantArray},

	{expression: `===`, patternType: keywordOperator},
	{expression: `!==`, patternType: keywordOperator},
	{expression: `function\*?\s*([A-Za-z$_][\w$]*)?\s*[(][^:;()]*[)]\s*{`, patternType: keywordFunction},

	// arrow function
	{expression: `\(* => {`, patternType: keywordFunction},

	{expression: `null`, patternType: constantNull},

	// lambda expression
	{expression: `\(.*\)\s*=>\s*.+`, patternType: keywordControl},

	{expression: `(else )?if\s+\(.+\)`, patternType: keywordControl},
	{expression: `while\s+\(.+\)`, patternType: keywordControl},

	// C style variable declaration
	{expression: `(^|\s)(char|long|int|float|double)\s+\w+\s*=?`, patternType: not},

	// pointer
	{expression: `\*\w+`, patternType: not},

	// HTML script tag
	{expression: `<(\/)?script( type=('|")text\/javascript('|"))?>`, patternType: not},
	{expression: `fn\s[A-Za-z0-9<>,]+\(.*\)\s->\s\w+(\s\{|)`, patternType: not},

	// avoiding C# confusion
	{expression: `Console\.(WriteLine|Write)(\s*)?\(`, patternType: not},
	{expression: `(using\s)?System(\..*)?(;)?`, patternType: not},
	{expression: `(func|fn)\s`, patternType: not},
	{expression: `(begin|end)\n`, patternType: not},

	// avoiding Lua confusion
	{expression: `local\s(function|(\w+)\s=)`, patternType: not},

	// avoiding Kotlin confusion
	{expression: `fun main\((.*)?\) {`, patternType: not},
	{expression: `(inline(\s+))?fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)({|=)`, patternType: not},
	{expression: `(const)?(\s+)?val(\s+)(.*)(:(\s)(.*)(\?)?)?(\s+)=(\s+)`, patternType: not},

	// avoiding Dart confusion
	{expression: `^(void\s)?main()\s{`, patternType: not},
}
