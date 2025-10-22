package flourite

var sql = []languagePattern{
	{expression: `(?i)CREATE (TABLE|DATABASE)`, patternType: keywordFunction, nearTop: true},
	{expression: `(?i)DROP (TABLE|DATABASE)`, patternType: keywordFunction, nearTop: true},
	{expression: `(?i)SHOW DATABASES`, patternType: keywordFunction, nearTop: true},
	{expression: `(?i)INSERT INTO`, patternType: keywordFunction},
	{expression: `(?i)(SELECT|SELECT DISTINCT)\s`, patternType: keywordFunction},
	{expression: `(?i)INNER JOIN`, patternType: keywordFunction},
	{expression: `(?i)(GROUP|ORDER) BY`, patternType: keywordFunction},
	{expression: `(?i)(END;|COMMIT;)`, patternType: keywordFunction},
	{expression: `(?i)UPDATE\s+\w+\sSET`, patternType: keywordFunction},
	{expression: `(?i)VALUES\s*\(\s*(?:\w+|'[^']*'|"[^"]*"|NULL)`, patternType: keywordFunction},
	{expression: `--\s*\w*`, patternType: commentLine},

	// data types
	{
		expression:  `(?i)(VARCHAR|CHAR|BINARY|VARBINARY|BLOB|TEXT|BIT|TINYINT|SMALLINT|MEDIUMINT|INT|INTEGER|BIGINT|DOUBLE)\([0-9]+\)`,
		patternType: constantType,
	},
	{
		expression:  `(?i)(TINYBLOB|TINYTEXT|MEDIUMTEXT|MEDIUMBLOB|LONGTEXT|LONGBLOB|BOOLEAN|BOOL|DATE|YEAR)`,
		patternType: constantType,
	},

	// math
	{expression: `(?i)(EXP|SUM|SQRT|MIN|MAX)`, patternType: keywordOperator},

	// avoiding Lua confusion
	{expression: `local\s(function|\w+)?\s=\s`, patternType: not},
	{expression: `(require|dofile)\((.*)\)`, patternType: not},
}
