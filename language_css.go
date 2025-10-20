package flourite

var css = []languagePattern{
	// properties check
	// [!] the original one isn't compatible with Golang regexp std library
	{expression: `[a-z-]+:\s*[^:;]+;`, patternType: keyword},

	// <style> tag from HTML
	{expression: `<(\/)?style>`, patternType: not},
}
