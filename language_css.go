package flourite

import "regexp"

var css = []languagePattern{
	// properties check
	// [!] the original one isn't compatible with Golang regexp std library
	{expression: regexp.MustCompile(`[a-z-]+:\s*[^:;]+;`), patternType: keyword},

	// <style> tag from HTML
	{expression: regexp.MustCompile(`<(\/)?style>`), patternType: not},
}
