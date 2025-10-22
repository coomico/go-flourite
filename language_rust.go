package flourite

import "regexp"

var rust = []languagePattern{
	{expression: regexp.MustCompile(`fn\smain\(\)`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(pub\s)?fn\s[A-Za-z0-9<>,]+\(.*\)\s->\s\w+(\s\{|)`), patternType: keywordVisibility},
	{expression: regexp.MustCompile(`let\smut\s\w+(\s=|)`), patternType: keywordVariable},
	{expression: regexp.MustCompile(`(.*)!\(.*\)`), patternType: macro},
	{expression: regexp.MustCompile(`use\s\w+::.*`), patternType: metaImport},
	{expression: regexp.MustCompile(`\{:\?\}`), patternType: keywordOther},
	{expression: regexp.MustCompile(`loop \{`), patternType: keywordControl},
	{expression: regexp.MustCompile(`(impl|crate|extern|macro|box)`), patternType: keywordOther},
	{expression: regexp.MustCompile(`match\s\w+\s\{`), patternType: keywordControl},
	{expression: regexp.MustCompile(`\w+\.len\(\)`), patternType: keywordOther},
	{expression: regexp.MustCompile(`(&str|(i|u)(8|16|32|64|128|size))`), patternType: constantType},

	// vector
	{expression: regexp.MustCompile(`(Vec|Vec::new)|vec!`), patternType: constantType},

	// traits
	{expression: regexp.MustCompile(`(Ok|Err|Box|ToOwned|Clone)`), patternType: keywordOther},

	{expression: regexp.MustCompile(`panic!\(.*\)`), patternType: keywordFunction},

	// avoiding clash with C#
	{expression: regexp.MustCompile(`using\sSystem`), patternType: not},
	{expression: regexp.MustCompile(`Console\.WriteLine\s*\(`), patternType: not},
	{expression: regexp.MustCompile(`(public\s)?((partial|static)\s)?class\s`), patternType: not},
	{expression: regexp.MustCompile(`(function|func)\s`), patternType: not},
}
