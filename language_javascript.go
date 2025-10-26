package flourite

import "regexp"

var javascript = []languagePattern{
	{expression: regexp.MustCompile(`undefined`), patternType: keyword},
	{expression: regexp.MustCompile(`window\.`), patternType: keyword},
	{expression: regexp.MustCompile(`console\.log\s*\(`), patternType: keywordPrint},
	{expression: regexp.MustCompile(`(var|const|let)\s+\w+\s*=?`), patternType: keywordVariable},

	// array/object declaration
	{expression: regexp.MustCompile(`(('|").+('|")\s*|\w+):\s*[{[]`), patternType: constantArray},

	{expression: regexp.MustCompile(`===`), patternType: keywordOperator},
	{expression: regexp.MustCompile(`!==`), patternType: keywordOperator},
	{
		expression:  regexp.MustCompile(`function\*?\s*([A-Za-z$_][\w$]*)?\s*\([^:;()]*\)\s*\{`),
		patternType: keywordFunction,
	},

	// arrow function
	{expression: regexp.MustCompile(`\(*\s*=>\s*\{`), patternType: keywordFunction},

	{expression: regexp.MustCompile(`null`), patternType: constantNull},

	// lambda expression
	{expression: regexp.MustCompile(`\(.*\)\s*=>\s*.+`), patternType: keywordControl},

	{expression: regexp.MustCompile(`(else )?if\s+\(.+\)`), patternType: keywordControl},
	{expression: regexp.MustCompile(`while\s+\(.+\)`), patternType: keywordControl},
	{expression: regexp.MustCompile(`for\s+\(.+\)`), patternType: keywordControl},

	// C style variable declaration
	{expression: regexp.MustCompile(`(^|\s)(char|long|int|float|double)\s+\w+\s*=?`), patternType: not},

	// pointer
	{expression: regexp.MustCompile(`\*\w+`), patternType: not},

	// HTML script tag
	{expression: regexp.MustCompile(`<(\/)?script( type=('|")text\/javascript('|"))?>`), patternType: not},
	{expression: regexp.MustCompile(`fn\s[A-Za-z0-9<>,]+\(.*\)\s->\s\w+(\s\{|)`), patternType: not},

	// avoiding C# confusion
	{expression: regexp.MustCompile(`Console\.(WriteLine|Write)(\s*)?\(`), patternType: not},
	{expression: regexp.MustCompile(`(using\s)?System(\..*)?(;)?`), patternType: not},
	{expression: regexp.MustCompile(`(func|fn)\s`), patternType: not},
	{expression: regexp.MustCompile(`(begin|end)\n`), patternType: not},

	// avoiding Lua confusion
	{expression: regexp.MustCompile(`local\s(function|(\w+)\s=)`), patternType: not},

	// avoiding Kotlin confusion
	{expression: regexp.MustCompile(`fun main\((.*)?\) \{`), patternType: not},
	{expression: regexp.MustCompile(`(inline(\s+))?fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)(\{|=)`), patternType: not},
	{expression: regexp.MustCompile(`(const)?(\s+)?val(\s+)(.*)(:(\s)(.*)(\?)?)?(\s+)=(\s+)`), patternType: not},

	// avoiding Dart confusion
	{expression: regexp.MustCompile(`^(void\s)?main\(\)\s\{`), patternType: not},
}
