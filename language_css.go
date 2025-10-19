package flourite

var css = []languagePattern{
	// properties check
	// [!] the original one doesn't compatible with Golang regex std library
	{expression: `[a-z-]+:\s*[^:;]+;`, patternType: keyword},

	// <style> tag from HTML
	{expression: `<(\/)?style>`, patternType: not},
}
