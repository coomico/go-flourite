package flourite

import "regexp"

var pascal = []languagePattern{
	{expression: regexp.MustCompile(`^program (.*);$`), patternType: metaModule, nearTop: true},
	{expression: regexp.MustCompile(`(?i)var$`), patternType: constantType, nearTop: true},
	{expression: regexp.MustCompile(`(?i)const$`), patternType: constantType, nearTop: true},
	{expression: regexp.MustCompile(`(?i)type$`), patternType: constantType, nearTop: true},
	{expression: regexp.MustCompile(`(?i)(write|writeln)(\s+)?(\((.*)\))?;`), patternType: keywordPrint},
	{expression: regexp.MustCompile(`(?i)^(\s*)?(function|procedure)(\s*)(.*)\((.*)\)(\s)?:(\s)?(.*);$`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(?i)end(\.|;)`), patternType: keywordControl},
	{expression: regexp.MustCompile(`(?i):(\s*)?(cardinal|shortint|smallint|word|extended|comp)(\s*);$`), patternType: constantType},
	{expression: regexp.MustCompile(`(?i)if(\s+)(.*)(\s+)then`), patternType: keywordControl},
	{expression: regexp.MustCompile(`(?i)for(\s+)(.*?):=(.*)(\s+)(downto|to)(\s+)(.*)(\s+)do`), patternType: keywordControl},
	{expression: regexp.MustCompile(`(?i)with(\s+)(.*)(\s+)do`), patternType: keywordControl},
	{expression: regexp.MustCompile(`repeat$`), patternType: keyword},
	{expression: regexp.MustCompile(`begin$`), patternType: keyword},
	{expression: regexp.MustCompile(`(?i)until(\s+)(.*);`), patternType: keywordControl},
	{expression: regexp.MustCompile(`(?i)\w+\s*:=\s*.+;$`), patternType: keywordVariable},
}
