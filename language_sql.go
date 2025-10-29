package flourite

import "regexp"

var sql = []languagePattern{
	{
		expression:  regexp.MustCompile(`(?i)\bCREATE\s+(?:TABLE|DATABASE)\b`),
		patternType: keywordFunction,
		nearTop:     true,
	},
	{expression: regexp.MustCompile(`(?i)\bDROP\s+(?:TABLE|DATABASE)\b`), patternType: keywordFunction, nearTop: true},
	{expression: regexp.MustCompile(`(?i)\bSHOW\s+DATABASES\b`), patternType: keywordFunction, nearTop: true},
	{expression: regexp.MustCompile(`(?i)\bINSERT\s+INTO\b`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(?i)(SELECT|SELECT DISTINCT)\s`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(?i)\bFROM\s+\w+`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(?i)\bWHERE\s+\w+`), patternType: keywordFunction},
	{
		expression:  regexp.MustCompile(`(?i)\b(?:INNER|LEFT|RIGHT|FULL(?:\s+OUTER)?)\s+JOIN\b`),
		patternType: keywordFunction,
	},
	{expression: regexp.MustCompile(`(?i)\b(?:GROUP|ORDER)\s+BY\b`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(?i)\b(?:END|COMMIT)\b;`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(?i)\bUPDATE\s+\w+`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(?i)\bSET\s+\w+`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(?i)VALUES\s*\(\s*(?:\w+|'[^']*'|"[^"]*"|NULL)`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`--[^\r\n]*`), patternType: commentLine},

	// data types
	{
		expression:  regexp.MustCompile(`(?i)(VARCHAR|CHAR|BINARY|VARBINARY|BLOB|TEXT|BIT|TINYINT|SMALLINT|MEDIUMINT|INT|INTEGER|BIGINT|DOUBLE)\([0-9]+\)`),
		patternType: constantType,
	},
	{
		expression:  regexp.MustCompile(`(?i)(TINYBLOB|TINYTEXT|MEDIUMTEXT|MEDIUMBLOB|LONGTEXT|LONGBLOB|BOOLEAN|BOOL|DATE|YEAR)`),
		patternType: constantType,
	},

	// Oracle & Postgre
	{expression: regexp.MustCompile(`(?i)\bINTERVAL\s+'[^']+'`), patternType: keywordFunction},
	{
		expression:  regexp.MustCompile(`(?i)\b(TIMESTAMP|TIMESTAMP WITH TIME ZONE|TIMESTAMPTZ|INTERVAL|CLOB|NCLOB)\b`),
		patternType: constantType,
	},

	// math
	{expression: regexp.MustCompile(`(?i)(EXP|SUM|SQRT|MIN|MAX)`), patternType: keywordOperator},

	// avoiding Lua confusion
	{expression: regexp.MustCompile(`local\s(function|\w+)?\s=\s`), patternType: not},
	{expression: regexp.MustCompile(`(require|dofile)\((.*)\)`), patternType: not},
}
