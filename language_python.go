package flourite

var python = []languagePattern{
	{expression: `def\s+\w+\(.*\)\s*:`, patternType: keywordFunction},
	{expression: `while (.+):`, patternType: keywordControl},
	{expression: `from [\w.]+ import (\w+|\*)`, patternType: metaImport},
	{expression: `class\s*\w+(\(\s*\w+\s*\))?\s*:`, patternType: keyword},
	{expression: `if\s+(.+)\s*:`, patternType: keywordControl},
	{expression: `elif\s+(.+)\s*:`, patternType: keywordControl},
	{expression: `else:`, patternType: keywordControl},
	{expression: `for (\w+|\(?\w+,\s*\w+\)?) in (.+):`, patternType: keywordControl},

	// python variable declaration
	{expression: `\w+\s*=\s*\w+(\s*)(\n|$)`, patternType: keyword},

	{expression: `import ([[^.]\w])+`, patternType: metaImport, nearTop: true},
	{expression: `print((\s*\(.+\))|\s+.+)`, patternType: keywordPrint},

	// &&/|| operators
	{expression: `(&{2}|\|{2})`, patternType: not},

	// avoiding Lua confusion
	{expression: `elseif`, patternType: not},
	{expression: `local\s(function|\w+)?\s=\s`, patternType: not},

	// avoiding Kotlin confusion
	{expression: `fun main\((.*)?\) {`, patternType: not},
	{expression: `(inline(\s+))?fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)({|=)`, patternType: not},
	{expression: `(const)?(\s+)?val(\s+)(.*)(:(\s)(.*)(\?)?)?(\s+)=(\s+)`, patternType: not},
}
