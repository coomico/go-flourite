package flourite

import "regexp"

var html = []languagePattern{
	{expression: regexp.MustCompile(`<!DOCTYPE (?:html|HTML PUBLIC .+)>`), patternType: metaModule, nearTop: true},

	// tags
	{expression: regexp.MustCompile(`<[a-z0-9]+(?:\s*[\w]+=(?:'|").+(?:'|")\s*)?>.+?</[a-z0-9]+>`), patternType: keyword},

	// comments
	{expression: regexp.MustCompile(`<!--(.*)(?:-->)?`), patternType: commentBlock},

	// properties
	{expression: regexp.MustCompile(`[a-z-]+=(?:"|').+(?:"|')`), patternType: keywordOther},

	// PHP tag
	{expression: regexp.MustCompile(`<\?php`), patternType: not},
}
