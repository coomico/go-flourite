package flourite

import "regexp"

var typescript = []languagePattern{
	{
		expression:  regexp.MustCompile(`(Readonly<|ReadonlyArray<|Array<|Record<|Pick<|Omit<|Exclude<|Extract<|Required<|Partial<|ReturnType<|Parameters<)`),
		patternType: constantDictionary,
	},
	{expression: regexp.MustCompile(`\w+\[\]`), patternType: keywordOther},
	{expression: regexp.MustCompile(`:\s*\w+\s*<`), patternType: keywordControl},
	{
		expression:  regexp.MustCompile(`\??\s*:\s*(any|string|boolean|undefined|null|number|void|never|symbol|bigint)\s*`),
		patternType: constantType,
	},
	{
		expression:  regexp.MustCompile(`(let|const|var)*\s*\w*:\s*(any|string|boolean|undefined|null|number|void|never|symbol|bigint)\s*=`),
		patternType: constantType,
	},
	{expression: regexp.MustCompile(`console\.log\s*\(`), patternType: keywordPrint},
	{expression: regexp.MustCompile(`interface\s*(\w+)\s*\{`), patternType: constantType},
	{expression: regexp.MustCompile(`enum\s*\w+\s*=`), patternType: constantType},
	{expression: regexp.MustCompile(`type\s+\w+<[^>]+>\s*=\s*\{`), patternType: constantType},
	{expression: regexp.MustCompile(`function\s+\w+\(.*\):\s\w+\s*\{`), patternType: keywordFunction},

	// arrow function
	{expression: regexp.MustCompile(`\(.*\):\s\w+\s*=>\s*\{`), patternType: keywordFunction},

	{expression: regexp.MustCompile(`(typeof|declare)\s+`), patternType: keyword},
	{expression: regexp.MustCompile(`\b(keyof|infer)\s+\w+`), patternType: keyword},
	{expression: regexp.MustCompile(`\s+as\s+`), patternType: keyword},

	// Rust types
	{expression: regexp.MustCompile(`usize`), patternType: not},

	// Kotlin
	{expression: regexp.MustCompile(`Array<String>`), patternType: not},
}
