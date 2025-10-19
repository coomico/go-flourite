package flourite

var kotlin = []languagePattern{
	{expression: `fun main\((.*)?\) {`, patternType: keywordFunction},
	{
		expression:  `(inline|private|public|protected|override|operator(\s+))?fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)({|=)`,
		patternType: keywordFunction,
	},
	{expression: `println\((.*)\)(\n|;)`, patternType: keywordPrint},
	{expression: `(else )?if\s*\(.+\)`, patternType: keywordControl},
	{expression: `while\s+\(.+\)`, patternType: keywordControl},
	{expression: `(const)?(\s+)?val(\s+)(.*)(:(\s)(.*)(\?)?)?(\s+)=(\s+)`, patternType: keywordVariable},
	{expression: `^(\s+)?(inner|open|data)(\s+)class`, patternType: keyword},
	{expression: `^import(\s+)(.*)$`, patternType: metaImport, nearTop: true},
	{expression: `typealias(\s+)(.*)(\s+)=`, patternType: keywordControl},
	{expression: `companion(\s+)object`, patternType: keyword},
	{expression: `when(\s+)(\((.*)\)\s+)?{$`, patternType: keywordControl},
}
