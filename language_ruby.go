package flourite

import "regexp"

var ruby = []languagePattern{
	{expression: regexp.MustCompile(`(require|include)\s+'\w+(\.rb)?'`), patternType: metaImport, nearTop: true},
	{expression: regexp.MustCompile(`^\s*def\s+\w+\s*(\(.+\))?`), patternType: keywordFunction},

	// instance variable
	{expression: regexp.MustCompile(`@\w+`), patternType: keywordOther},

	// boolean property
	{expression: regexp.MustCompile(`\.\w+\?`), patternType: constantBoolean},

	// Ruby print
	{expression: regexp.MustCompile(`puts\s+("|').+("|')`), patternType: keywordPrint},

	// inheriting class
	{expression: regexp.MustCompile(`class\s+[A-Z]\w*\s*(?:<\s*[A-Z]\w*)?`), patternType: keyword},

	{expression: regexp.MustCompile(`attr_accessor\s+(:\w+(,\s*)?)+`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`\w+\.new\s+`), patternType: keyword},
	{expression: regexp.MustCompile(`elsif`), patternType: keywordControl},
	{expression: regexp.MustCompile(`^\s*module\s+[A-Z]\w*\s*$`), patternType: keywordOther},
	{expression: regexp.MustCompile(`\bBEGIN\s\{.*\}`), patternType: keywordOther},
	{expression: regexp.MustCompile(`\bEND\s\{.*\}`), patternType: keywordOther},
	{expression: regexp.MustCompile(`do\s*\|\w+(,\s*\w+)*\|`), patternType: keywordControl},
	{expression: regexp.MustCompile(`for (\w+|\(?\w+,\s*\w+\)?) in (.+)`), patternType: keywordControl},
	{expression: regexp.MustCompile(`nil`), patternType: constantNull},

	// scope operator
	{expression: regexp.MustCompile(`[A-Z]\w*::[A-Z]\w*`), patternType: macro},

	{expression: regexp.MustCompile(`\[\d+\.\.-?\d+\]`), patternType: keywordOperator},
	{expression: regexp.MustCompile(`\d+\.\.\.?[\d\(]+`), patternType: keywordOperator},

	// termination code blocks `end`
	{expression: regexp.MustCompile(`^\s*end\b\s*(?:#.*)?$`), patternType: keyword},

	// ruby-specific methods
	{expression: regexp.MustCompile(`\.(each_index|each|times|downto|upto)\b`), patternType: keywordControl},
	{expression: regexp.MustCompile(`\.(partition|select|reject|map|inject)\s*\{`), patternType: keywordControl},
	{expression: regexp.MustCompile(`\.(dup|sort|push|floor|ceil|size|length)\b`), patternType: keywordOther},
	{expression: regexp.MustCompile(`\.to_(i|f|s|a|h)\b`), patternType: keywordOther},
	{expression: regexp.MustCompile(`Array\.\[\]`), patternType: keyword},

	{expression: regexp.MustCompile(`^\s*p\s+\w+`), patternType: keywordPrint},

	// string interpolation
	{expression: regexp.MustCompile(`#\{[^}]+\}`), patternType: constantString},
}
