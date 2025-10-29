package flourite

import "regexp"

var rust = []languagePattern{
	{expression: regexp.MustCompile(`fn\smain\(\)`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`\bfn\s+\w+`), patternType: keywordFunction},

	// Fn trait or closure or function ptr type
	{expression: regexp.MustCompile(`[fF]n\s*\(`), patternType: keyword},

	{expression: regexp.MustCompile(`\bpub\s+`), patternType: keywordVisibility},
	{expression: regexp.MustCompile(`let\s+(?:mut\s+)?\w+(\s=|)`), patternType: keywordVariable},
	{expression: regexp.MustCompile(`\w+!\s*\(`), patternType: macro},
	{expression: regexp.MustCompile(`use\s+\w+::.*`), patternType: metaImport},
	{expression: regexp.MustCompile(`\{:\?\}`), patternType: keywordOther},
	{expression: regexp.MustCompile(`loop\s*\{`), patternType: keywordControl},
	{expression: regexp.MustCompile(`for\s+\w+\s+in\s+`), patternType: keywordControl},

	// range expression
	{expression: regexp.MustCompile(`\d*\.\.=?[\d\w]*(?:\s*[\)\{;]|$|\s+)`), patternType: keywordControl},

	// tuple destructuring in match
	{expression: regexp.MustCompile(`\(\d+,\s*\d+\)\s*=>`), patternType: keywordControl},

	{expression: regexp.MustCompile(`match\s+.+\s*\{`), patternType: keywordControl},
	{
		expression:  regexp.MustCompile(`\b(impl|crate|extern|macro|box|mod|trait|enum|struct|const|static|where)\b`),
		patternType: keywordOther,
	},
	{expression: regexp.MustCompile(`\w+\.\s*(len|swap|take|eq|chars)\s*\(\)`), patternType: keywordOther},
	{
		expression:  regexp.MustCompile(`\b(&str|&\[T\]|&mut\s+\[T\]|(i|u|f)(8|16|32|64|128|size)|usize|bool)\b`),
		patternType: constantType,
	},
	{expression: regexp.MustCompile(`macro_rules!\s+\w+`), patternType: macro},

	// referencing and borrowing
	{expression: regexp.MustCompile(`&(mut\s+)?\w+`), patternType: keywordOperator},

	// function return type
	{
		expression:  regexp.MustCompile(`->\s*(?:&|\()?(usize|bool|i\d+|u\d+|f\d+|String|str|\w+)`),
		patternType: keywordFunction,
	},

	// vector
	{expression: regexp.MustCompile(`(Vec|Vec::new)|vec!`), patternType: constantType},

	// traits
	{
		expression:  regexp.MustCompile(`\b(Ok|Err|Box|ToOwned|Clone|Some|None|Result|Option)\b`),
		patternType: keywordOther,
	},

	{expression: regexp.MustCompile(`panic!\(.*\)`), patternType: keywordFunction},

	// avoiding clash with C#
	{expression: regexp.MustCompile(`using\sSystem`), patternType: not},
	{expression: regexp.MustCompile(`Console\.WriteLine\s*\(`), patternType: not},
	{expression: regexp.MustCompile(`(public\s)?((partial|static)\s)?class\s`), patternType: not},
	{expression: regexp.MustCompile(`(function|func)\s`), patternType: not},
}
