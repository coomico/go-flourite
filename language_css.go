package flourite

import "regexp"

var css = []languagePattern{
	// properties check
	// [!] the original one isn't compatible with Golang regexp std library
	{expression: regexp.MustCompile(`^\s*[a-z-]+:\s*[^:;{}\[\]]+;?`), patternType: keyword},

	// selector
	{expression: regexp.MustCompile(`^[a-zA-Z][\w-]*(\[[^\]]+\])?(\:[a-z-]+)?\s*[,{]`), patternType: keyword},

	// specific property:value
	{
		expression:  regexp.MustCompile(`(border|margin|padding|background|color|display|position|width|height|font|text)(-[a-z]+)?:\s*[^:;]+;?`),
		patternType: keyword,
	},

	{expression: regexp.MustCompile(`/\*[\s\S]*?\*/`), patternType: commentBlock},

	// <style> tag from HTML
	{expression: regexp.MustCompile(`<(\/)?style>`), patternType: not},
}
