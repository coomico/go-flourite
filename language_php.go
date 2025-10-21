package flourite

var php = []languagePattern{
	// PHP tag
	{expression: `<\?php`, patternType: metaModule},

	{expression: `\$\w+`, patternType: keywordVariable},

	// use Something/Something;
	{expression: `use\s+\w+(\\\w+)+\s*;`, patternType: metaImport, nearTop: true},

	// arrow
	{expression: `\$\w+->\w+`, patternType: keyword},

	{expression: `(require|include)(_once)?\s*\(?\s*('|").+\.php('|")\s*\)?\s*;`, patternType: metaImport},
	{expression: `echo\s+('|").+('|")\s*;`, patternType: keywordPrint},
	{expression: `NULL`, patternType: constantNull},
	{expression: `new\s+((\\\w+)+|\w+)(\(.*\))?`, patternType: keyword},
	{expression: `function(\s+[$\w]+\(.*\)|\s*\(.*\))`, patternType: keywordFunction},
	{expression: `(else)?if\s+\(.+\)`, patternType: keywordControl},

	// scope operator
	{expression: `\w+::\w+`, patternType: keyword},

	{expression: `===`, patternType: keywordOperator},
	{expression: `!==`, patternType: keywordOperator},

	// C/JS style variable declaration
	{expression: `(^|\s)(var|char|long|int|float|double)\s+\w+\s*=?`, patternType: not},

	// JavaScript variable declaration
	{expression: `(var|const|let)\s+\w+\s*=?`, patternType: not},

	// avoiding Lua confusion
	{expression: `local\s(function|\w+)`, patternType: not},
}
