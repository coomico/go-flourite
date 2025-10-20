package flourite

var rust = []languagePattern{
	{expression: `fn\smain()`, patternType: keywordFunction},
	{expression: `(pub\s)?fn\s[A-Za-z0-9<>,]+\(.*\)\s->\s\w+(\s\{|)`, patternType: keywordVisibility},
	{expression: `let\smut\s\w+(\s=|)`, patternType: keywordVariable},
	{expression: `(.*)!\(.*\)`, patternType: macro},
	{expression: `use\s\w+::.*`, patternType: metaImport},
	{expression: `\{:\?\}`, patternType: keywordOther},
	{expression: `loop \{`, patternType: keywordControl},
	{expression: `(impl|crate|extern|macro|box)`, patternType: keywordOther},
	{expression: `match\s\w+\s\{`, patternType: keywordControl},
	{expression: `\w+\.len\(\)`, patternType: keywordOther},
	{expression: `(&str|(i|u)(8|16|32|64|128|size))`, patternType: constantType},

	// vector
	{expression: `(Vec|Vec::new)|vec!`, patternType: constantType},

	// traits
	{expression: `(Ok|Err|Box|ToOwned|Clone)`, patternType: keywordOther},

	{expression: `panic!\(.*\)`, patternType: keywordFunction},

	// avoiding clash with C#
	{expression: `using\sSystem`, patternType: not},
	{expression: `Console\.WriteLine\s*\(`, patternType: not},
	{expression: `(public\s)?((partial|static)\s)?class\s`, patternType: not},
	{expression: `(function|func)\s`, patternType: not},
}
