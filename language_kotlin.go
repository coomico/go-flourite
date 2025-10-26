package flourite

import "regexp"

var kotlin = []languagePattern{
	{expression: regexp.MustCompile(`fun main\([^)]*\)\s*\{`), patternType: keywordFunction},
	{
		expression:  regexp.MustCompile(`(inline|private|public|protected|override|operator\s+)?fun\s+[A-Za-z0-9_]+\s*\(.*\)\s*(\{|=|:)`),
		patternType: keywordFunction,
	},
	{expression: regexp.MustCompile(`println\([^)]*\)`), patternType: keywordPrint},
	{expression: regexp.MustCompile(`(else )?if\s*\(.+\)`), patternType: keywordControl},
	{expression: regexp.MustCompile(`while\s+\(.+\)`), patternType: keywordControl},
	{
		expression:  regexp.MustCompile(`(const\s+)?val\s+[A-Za-z0-9_]+(\s*:\s*[A-Za-z0-9_<>,\s\[\]?]+)?\s*=`),
		patternType: keywordVariable,
	},
	{expression: regexp.MustCompile(`^\s*?(inner|open|data)\s+class`), patternType: keyword},
	{expression: regexp.MustCompile(`^import\s+(.*)$`), patternType: metaImport, nearTop: true},
	{expression: regexp.MustCompile(`typealias\s+(.*)\s+=`), patternType: keywordControl},
	{expression: regexp.MustCompile(`companion\s+object`), patternType: keyword},
	{expression: regexp.MustCompile(`when\s*(\([^)]*\)\s*)?\{`), patternType: keywordControl},
}
