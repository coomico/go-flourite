package flourite

import "regexp"

var json = []languagePattern{
	// object declaration on top
	{expression: regexp.MustCompile(`^\{$`), patternType: metaModule, nearTop: true},

	// normal data type
	{expression: regexp.MustCompile(`^\s*".+":\s*(".+"|[0-9]+|null|true|false)(,)?$`), patternType: keyword},

	// object and array
	{expression: regexp.MustCompile(`^\s*".+":\s*(\{|\[)$`), patternType: keyword},

	// inline key/value pair in object
	{expression: regexp.MustCompile(`^\s*".+":\s*\{(\s*".+":\s*(".+"|[0-9]+|null|true|false)(,)?\s*)+\}(,)?$`), patternType: keyword},

	// inline value in array
	{expression: regexp.MustCompile(`\s*".+"\s*\[\s*((".+"|[0-9]+|null|true|false)(,)?\s*)+\](,)?$`), patternType: keyword},
}
