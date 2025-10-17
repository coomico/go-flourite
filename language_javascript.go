package flourite

var js = []LanguagePattern{
	{Pattern: `undefined`, Type: Keyword},
	{Pattern: `window\.`, Type: Keyword},
	{Pattern: `console\.log\s*\(`, Type: KeywordPrint},
	{Pattern: `(var|const|let)\s+\w+\s*=?`, Type: KeywordVariable},

	// array/object declaration
	{Pattern: `(('|").+('|")\s*|\w+):\s*[{[]`, Type: ConstantArray},

	{Pattern: `===`, Type: KeywordOperator},
	{Pattern: `!==`, Type: KeywordOperator},
	{Pattern: `function\*?\s*([A-Za-z$_][\w$]*)?\s*[(][^:;()]*[)]\s*{`, Type: KeywordFunction},

	// arrow function
	{Pattern: `\(* => {`, Type: KeywordFunction},

	{Pattern: `null`, Type: ConstantNull},

	// lambda expression
	{Pattern: `\(.*\)\s*=>\s*.+`, Type: KeywordControl},

	{Pattern: `(else )?if\s+\(.+\)`, Type: KeywordControl},
	{Pattern: `while\s+\(.+\)`, Type: KeywordControl},

	// C style variable declaration
	{Pattern: `(^|\s)(char|long|int|float|double)\s+\w+\s*=?`, Type: Not},

	// pointer
	{Pattern: `\*\w+`, Type: Not},

	// HTML script tag
	{Pattern: `<(\/)?script( type=('|")text\/javascript('|"))?>`, Type: Not},
	{Pattern: `fn\s[A-Za-z0-9<>,]+\(.*\)\s->\s\w+(\s\{|)`, Type: Not},

	// avoiding C# confusion
	{Pattern: `Console\.(WriteLine|Write)(\s*)?\(`, Type: Not},
	{Pattern: `(using\s)?System(\..*)?(;)?`, Type: Not},
	{Pattern: `(func|fn)\s`, Type: Not},
	{Pattern: `(begin|end)\n`, Type: Not},

	// avoiding Lua confusion
	{Pattern: `local\s(function|(\w+)\s=)`, Type: Not},

	// avoiding Kotlin confusion
	{Pattern: `fun main\((.*)?\) {`, Type: Not},
	{Pattern: `(inline(\s+))?fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)({|=)`, Type: Not},
	{Pattern: `(const)?(\s+)?val(\s+)(.*)(:(\s)(.*)(\?)?)?(\s+)=(\s+)`, Type: Not},

	// avoiding Dart confusion
	{Pattern: `^(void\s)?main()\s{`, Type: Not},
}
