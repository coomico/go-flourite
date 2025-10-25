package flourite

import "regexp"

var dart = []languagePattern{
	{
		expression:  regexp.MustCompile(`^\s*(const|final|var|dynamic|late)?\s*(int|double|String|bool|List<[A-Za-z \[\](),]+>|HashMap<[A-Za-z \[\](),]+>|Iterator<[A-Za-z \[\](),]+>|Set<[A-Za-z \[\](),]+>)?(\?)?\s\w+(\s=\s.+)?(;|,)$`),
		patternType: keywordVariable,
	},
	{expression: regexp.MustCompile(`\bstdout\.write\(.+\);`), patternType: keywordPrint},
	{expression: regexp.MustCompile(`\bprint\(.+\);`), patternType: keywordPrint},
	{expression: regexp.MustCompile(`^\s*import\s+['"]dart:\w+['"]`), patternType: metaImport, nearTop: true},
	{expression: regexp.MustCompile(`^\s*import\s+['"]package:[\w/]+['"]`), patternType: metaImport, nearTop: true},
	{expression: regexp.MustCompile(`^\s*library\s\w+;`), patternType: metaModule, nearTop: true},
	{
		expression:  regexp.MustCompile(`^\s*(Future<void>\s+|void\s+)?main\s*\(\s*\)\s*(async\s*)?\{?`),
		patternType: keywordFunction,
	},

	// function with type definition
	{
		expression:  regexp.MustCompile(`^\s*(List<[A-Za-z \[\](),]+>|HashMap<[A-Za-z \[\](),]+>|int|double|String|bool|void|Iterator<[A-Za-z \[\](),]+>|Set<[A-Za-z \[\](),]+>)\s\w+\(.+\)\s*\{$`),
		patternType: keywordFunction,
	},

	// arrow function
	{
		expression:  regexp.MustCompile(`^\s*(int|double|String|bool|List<[A-Za-z \[\](),]+>|HashMap<[A-Za-z \[\](),]+>|Iterator<[A-Za-z \[\](),]+>|Set<[A-Za-z \[\](),]+>)\s\w+\(.+\)\s=>`),
		patternType: keywordFunction,
	},

	{
		expression:  regexp.MustCompile(`\bnew\s+(List|Map|Iterator|HashMap|Set)(<[^>]+>)?(\.\w+)?\([^)]*\)`),
		patternType: keywordVariable,
	},
	{
		expression:  regexp.MustCompile(`^(abstract\s)?class\s\w+\s(extends\s\w+\s)?(with\s\w+\s)?(implements\s\w+\s)?\{(\})?$`),
		patternType: keywordControl,
	},
	{expression: regexp.MustCompile(`\bget\s\w+=>\w+`), patternType: keywordControl},
	{expression: regexp.MustCompile(`^\s*@override$`), patternType: keywordControl},
	{expression: regexp.MustCompile(`\bset\s\w+\(.+\)`), patternType: keywordControl},
	{expression: regexp.MustCompile(`^\s*Future<\w+>\s\w+\(.*\)\s+async`), patternType: keywordControl},
	{expression: regexp.MustCompile(`\bawait\s+for\s*\(`), patternType: keywordControl},
	{
		expression:  regexp.MustCompile(`\btypedef\s+\w+(<[^>]+>)?\s*=\s*\w+\s+Function\s*\([^)]*\)`),
		patternType: keywordControl,
	},
	{expression: regexp.MustCompile(`\bassert\s*\(.+\)`), patternType: keyword},

	// avoiding C confusion
	{expression: regexp.MustCompile(`\blong\s+long\b`), patternType: not},
	{expression: regexp.MustCompile(`\blong\s`), patternType: not},
	{expression: regexp.MustCompile(`\s*function\b`), patternType: not},

	// avoiding Java confusion
	{expression: regexp.MustCompile(`\bArrayList`), patternType: not},
}
