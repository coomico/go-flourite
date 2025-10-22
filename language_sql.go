package flourite

import "regexp"

var sql = []languagePattern{
	{expression: regexp.MustCompile(`(?i)CREATE (TABLE|DATABASE)`), patternType: keywordFunction, nearTop: true},
	{expression: regexp.MustCompile(`(?i)DROP (TABLE|DATABASE)`), patternType: keywordFunction, nearTop: true},
	{expression: regexp.MustCompile(`(?i)SHOW DATABASES`), patternType: keywordFunction, nearTop: true},
	{expression: regexp.MustCompile(`(?i)INSERT INTO`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(?i)(SELECT|SELECT DISTINCT)\s`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(?i)INNER JOIN`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(?i)(GROUP|ORDER) BY`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(?i)(END;|COMMIT;)`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(?i)UPDATE\s+\w+\sSET`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`(?i)VALUES\s*\(\s*(?:\w+|'[^']*'|"[^"]*"|NULL)`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`--\s*\w*`), patternType: commentLine},

	// data types
	{
		expression:  regexp.MustCompile(`(?i)(VARCHAR|CHAR|BINARY|VARBINARY|BLOB|TEXT|BIT|TINYINT|SMALLINT|MEDIUMINT|INT|INTEGER|BIGINT|DOUBLE)\([0-9]+\)`),
		patternType: constantType,
	},
	{
		expression:  regexp.MustCompile(`(?i)(TINYBLOB|TINYTEXT|MEDIUMTEXT|MEDIUMBLOB|LONGTEXT|LONGBLOB|BOOLEAN|BOOL|DATE|YEAR)`),
		patternType: constantType,
	},

	// math
	{expression: regexp.MustCompile(`(?i)(EXP|SUM|SQRT|MIN|MAX)`), patternType: keywordOperator},

	// avoiding Lua confusion
	{expression: regexp.MustCompile(`local\s(function|\w+)?\s=\s`), patternType: not},
	{expression: regexp.MustCompile(`(require|dofile)\((.*)\)`), patternType: not},
}
