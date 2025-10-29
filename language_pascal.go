package flourite

import "regexp"

var pascal = []languagePattern{
	{expression: regexp.MustCompile(`^\s*program\s+\w+\s*;?\s*$`), patternType: metaModule, nearTop: true},
	{expression: regexp.MustCompile(`(?i)^\s*var\s*$`), patternType: constantType},
	{expression: regexp.MustCompile(`(?i)^\s*const\s*$`), patternType: constantType, nearTop: true},
	{expression: regexp.MustCompile(`(?i)^\s*type\s*$`), patternType: constantType, nearTop: true},
	{expression: regexp.MustCompile(`(?i)\b(write|writeln|read|readln)\s*(\([^)]*\))?\s*;`), patternType: keywordPrint},
	{
		expression:  regexp.MustCompile(`(?i)^\s*(function|procedure)\s+\w+\s*(\([^)]*\))?\s*(:\s*\w+)?\s*;`),
		patternType: keywordFunction,
	},

	// built-in functions
	{
		expression:  regexp.MustCompile(`(?i)\b(length|copy|reverse|chr|ord|inc|dec|succ|pred)\s*\(`),
		patternType: keywordFunction,
	},

	{expression: regexp.MustCompile(`(?i)\s*begin\s*$`), patternType: keywordControl},
	{expression: regexp.MustCompile(`(?i)end[\.;\s\n]`), patternType: keywordControl},
	{expression: regexp.MustCompile(`^\s*\{[^}]*\}\s*$`), patternType: commentLine},
	{
		expression:  regexp.MustCompile(`(?i):\s*(boolean|integer|real|char|string|cardinal|shortint|smallint|word|extended|comp|byte|longint)\s*;?\s*$`),
		patternType: constantType,
	},
	{expression: regexp.MustCompile(`(?i)\bif\b.*\bthen\b`), patternType: keywordControl},
	{expression: regexp.MustCompile(`(?i)^\s*else\b`), patternType: keywordControl},
	{
		expression:  regexp.MustCompile(`(?i)for(\s+)(.*?):=(.*)(\s+)(downto|to)(\s+)(.*)(\s+)do`),
		patternType: keywordControl,
	},
	{expression: regexp.MustCompile(`(?i)\bwhile\b.*\bdo\b`), patternType: keywordControl},
	{expression: regexp.MustCompile(`(?i)with(\s+)(.*)(\s+)do`), patternType: keywordControl},
	{expression: regexp.MustCompile(`(?i)\brepeat\b\s*$`), patternType: keyword},
	{expression: regexp.MustCompile(`(?i)until(\s+)(.*);`), patternType: keywordControl},
	{expression: regexp.MustCompile(`(?i)\w+\s*:=\s*.+;$`), patternType: keywordVariable},
}
