package flourite

var html = []languagePattern{
	{expression: `<!DOCTYPE (?:html|HTML PUBLIC .+)>`, patternType: metaModule, nearTop: true},

	// tags
	{expression: `<[a-z0-9]+(?:\s*[\w]+=(?:'|").+(?:'|")\s*)?>.+?</[a-z0-9]+>`, patternType: keyword},

	// comments
	{expression: `<!--(.*)(?:-->)?`, patternType: commentBlock},

	// properties
	{expression: `[a-z-]+=(?:"|').+(?:"|')`, patternType: keywordOther},

	// PHP tag
	{expression: `<\?php`, patternType: not},
}
