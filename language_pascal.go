package flourite

var pascal = []languagePattern{
	{expression: `^program (.*);$`, patternType: metaModule, nearTop: true},
	{expression: `(?i)var$`, patternType: constantType, nearTop: true},
	{expression: `(?i)const$`, patternType: constantType, nearTop: true},
	{expression: `(?i)type$`, patternType: constantType, nearTop: true},
	{expression: `(?i)(write|writeln)(\s+)?(\((.*)\))?;`, patternType: keywordPrint},
	{expression: `(?i)^(\s*)?(function|procedure)(\s*)(.*)\((.*)\)(\s)?:(\s)?(.*);$`, patternType: keywordFunction},
	{expression: `(?i)end(\.|;)`, patternType: keywordControl},
	{expression: `(?i):(\s*)?(cardinal|shortint|smallint|word|extended|comp)(\s*);$`, patternType: constantType},
	{expression: `(?i)if(\s+)(.*)(\s+)then`, patternType: keywordControl},
	{expression: `(?i)for(\s+)(.*?):=(.*)(\s+)(downto|to)(\s+)(.*)(\s+)do`, patternType: keywordControl},
	{expression: `(?i)with(\s+)(.*)(\s+)do`, patternType: keywordControl},
	{expression: `repeat$`, patternType: keyword},
	{expression: `begin$`, patternType: keyword},
	{expression: `(?i)until(\s+)(.*);`, patternType: keywordControl},
	{expression: `(?i)\w+\s*:=\s*.+;$`, patternType: keywordVariable},
}
