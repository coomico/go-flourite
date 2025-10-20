package flourite

var sql = []languagePattern{
	{expression: `(?i)CREATE (TABLE|DATABASE)`, patternType: keyword, nearTop: true},
	{expression: `(?i)DROP (TABLE|DATABASE)`, patternType: keyword, nearTop: true},
	{expression: `(?i)SHOW DATABASES`, patternType: keyword, nearTop: true},
	{expression: `(?i)INSERT INTO`, patternType: keyword},
	{expression: `(?i)(SELECT|SELECT DISTINCT)\s`, patternType: keyword},
	{expression: `(?i)INNER JOIN`, patternType: keyword},
	{expression: `(?i)(GROUP|ORDER) BY`, patternType: keyword},
	{expression: `(?i)(END;|COMMIT;)`, patternType: keyword},
	{expression: `(?i)UPDATE\s+\w+\sSET`, patternType: keyword},
	{expression: `(?i)VALUES+(\s+\(\w|\(\w)`, patternType: keyword},
	{expression: `--\s\w`, patternType: commentLine},

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
