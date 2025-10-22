package flourite

import "regexp"

var php = []languagePattern{
	// PHP tag
	{expression: regexp.MustCompile(`<\?php`), patternType: metaModule},

	{expression: regexp.MustCompile(`\$\w+`), patternType: keywordVariable},

	// use Something/Something;
	{expression: regexp.MustCompile(`use\s+\w+(\\\w+)+\s*;`), patternType: metaImport, nearTop: true},

	// arrow
	{expression: regexp.MustCompile(`\$\w+->\w+`), patternType: keyword},

	{expression: regexp.MustCompile(`(require|include)(_once)?\s*\(?\s*('|").+\.php('|")\s*\)?\s*;`), patternType: metaImport},
	{expression: regexp.MustCompile(`echo\s+('|").+('|")\s*;`), patternType: keywordPrint},
	{expression: regexp.MustCompile(`NULL`), patternType: constantNull},
	{expression: regexp.MustCompile(`new\s+((\\\w+)+|\w+)(\(.*\))?`), patternType: keyword},
	{expression: regexp.MustCompile(`function(\s+[$\w]+\(.*\)|\s*\(.*\))`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(else)?if\s+\(.+\)`), patternType: keywordControl},

	// scope operator
	{expression: regexp.MustCompile(`\w+::\w+`), patternType: keyword},

	{expression: regexp.MustCompile(`===`), patternType: keywordOperator},
	{expression: regexp.MustCompile(`!==`), patternType: keywordOperator},

	// C/JS style variable declaration
	{expression: regexp.MustCompile(`(^|\s)(var|char|long|int|float|double)\s+\w+\s*=?`), patternType: not},

	// JavaScript variable declaration
	{expression: regexp.MustCompile(`(var|const|let)\s+\w+\s*=?`), patternType: not},

	// avoiding Lua confusion
	{expression: regexp.MustCompile(`local\s(function|\w+)`), patternType: not},
}
