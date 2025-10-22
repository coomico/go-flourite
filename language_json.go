package flourite

var json = []languagePattern{
	// object declaration on top
	{expression: `^\{$`, patternType: metaModule, nearTop: true},

	// normal data type
	{expression: `^\s*".+":\s*(".+"|[0-9]+|null|true|false)(,)?$`, patternType: keyword},

	// object and array
	{expression: `^\s*".+":\s*(\{|\[)$`, patternType: keyword},

	// inline key/value pair in object
	{expression: `^\s*".+":\s*\{(\s*".+":\s*(".+"|[0-9]+|null|true|false)(,)?\s*)+\}(,)?$`, patternType: keyword},

	// inline value in array
	{expression: `\s*".+"\s*\[\s*((".+"|[0-9]+|null|true|false)(,)?\s*)+\](,)?$`, patternType: keyword},
}
