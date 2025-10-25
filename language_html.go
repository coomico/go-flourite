package flourite

import "regexp"

var html = []languagePattern{
	{expression: regexp.MustCompile(`<!DOCTYPE (?:html|HTML PUBLIC .+)>`), patternType: metaModule, nearTop: true},

	// HTML-specific tags
	{expression: regexp.MustCompile(`<(html|head|body|title|meta|link)(\s|>)`), patternType: keyword},

	// tags
	{expression: regexp.MustCompile(`<[a-z][a-z0-9]*(?:\s+[\w-]+=(?:'[^']*'|"[^"]*"))*\s*/?>`), patternType: keyword},

	// closing tags separately
	{expression: regexp.MustCompile(`</[a-z][a-z0-9]*>`), patternType: keyword},

	// comments
	{expression: regexp.MustCompile(`<!--(.*)(?:-->)?`), patternType: commentBlock},

	// properties
	{expression: regexp.MustCompile(`[a-z-]+=(?:"|').+(?:"|')`), patternType: keywordOther},

	// PHP tag
	{expression: regexp.MustCompile(`<\?php`), patternType: not},
}
