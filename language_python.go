package flourite

import "regexp"

var python = []languagePattern{
	{expression: regexp.MustCompile(`def\s+\w+\(.*\)\s*:`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`while (.+):`), patternType: keywordControl},
	{expression: regexp.MustCompile(`from [\w.]+ import (\w+|\*)`), patternType: metaImport},
	{expression: regexp.MustCompile(`class\s*\w+(\(\s*\w+\s*\))?\s*:`), patternType: keyword},
	{expression: regexp.MustCompile(`if\s+(.+)\s*:`), patternType: keywordControl},
	{expression: regexp.MustCompile(`elif\s+(.+)\s*:`), patternType: keywordControl},
	{expression: regexp.MustCompile(`else:`), patternType: keywordControl},
	{expression: regexp.MustCompile(`for (\w+|\(?\w+,\s*\w+\)?) in (.+):`), patternType: keywordControl},

	// python variable declaration
	// [!] the original one isn't compatible with Golang regexp std library
	{expression: regexp.MustCompile(`\w+\s*=\s*[^:;\n=]+\s*?(\n|$)`), patternType: keyword},

	{expression: regexp.MustCompile(`import ([[^.]\w])+`), patternType: metaImport, nearTop: true},
	{expression: regexp.MustCompile(`print((\s*\(.+\))|\s+.+)`), patternType: keywordPrint},

	// &&/|| operators
	{expression: regexp.MustCompile(`(&{2}|\|{2})`), patternType: not},

	// avoiding Lua confusion
	{expression: regexp.MustCompile(`elseif`), patternType: not},
	{expression: regexp.MustCompile(`local\s(function|\w+)?\s=\s`), patternType: not},

	// avoiding Kotlin confusion
	{expression: regexp.MustCompile(`fun main\((.*)?\) \{`), patternType: not},
	{expression: regexp.MustCompile(`(inline(\s+))?fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)(\{|=)`), patternType: not},
	{expression: regexp.MustCompile(`(const)?(\s+)?val(\s+)(.*)(:(\s)(.*)(\?)?)?(\s+)=(\s+)`), patternType: not},
}
