package flourite

import "regexp"

var ruby = []languagePattern{
	{expression: regexp.MustCompile(`(require|include)\s+'\w+(\.rb)?'`), patternType: metaImport, nearTop: true},
	{expression: regexp.MustCompile(`def\s+\w+\s*(\(.+\))?\s*\n`), patternType: keywordFunction},

	// instance variable
	{expression: regexp.MustCompile(`@\w+`), patternType: keywordOther},

	// boolean property
	{expression: regexp.MustCompile(`\.\w+\?`), patternType: constantBoolean},

	// Ruby print
	{expression: regexp.MustCompile(`puts\s+("|').+("|')`), patternType: keywordPrint},

	// inheriting class
	{expression: regexp.MustCompile(`class [A-Z]\w*\s*<\s*([A-Z]\w*(::)?)+`), patternType: keyword},

	{expression: regexp.MustCompile(`attr_accessor\s+(:\w+(,\s*)?)+`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`\w+\.new\s+`), patternType: keyword},
	{expression: regexp.MustCompile(`elsif`), patternType: keywordControl},
	{expression: regexp.MustCompile(`\bmodule\s\S`), patternType: keywordOther},
	{expression: regexp.MustCompile(`\bBEGIN\s\{.*\}`), patternType: keywordOther},
	{expression: regexp.MustCompile(`\bEND\s\{.*\}`), patternType: keywordOther},
	{expression: regexp.MustCompile(`do\s*\|\w+(,\s*\w+)*\|`), patternType: keywordControl},
	{expression: regexp.MustCompile(`for (\w+|\(?\w+,\s*\w+\)?) in (.+)`), patternType: keywordControl},
	{expression: regexp.MustCompile(`nil`), patternType: constantNull},

	// scope operator
	{expression: regexp.MustCompile(`[A-Z]\w*::[A-Z]\w*`), patternType: macro},

	// termination code blocks `end`
	{expression: regexp.MustCompile(`^\s*end\b\s*(?:#.*)?$`), patternType: keyword},
}
