package flourite

import "regexp"

var cpp = []languagePattern{
	{expression: regexp.MustCompile(`(char|long|int|float|double)\s+\w+\s*=?`), patternType: constantType},
	{expression: regexp.MustCompile(`#include\s*(<|")\w+(\.h)?(>|")`), patternType: metaImport},
	{expression: regexp.MustCompile(`using\s+namespace\s+.+\s*;`), patternType: keyword},
	{expression: regexp.MustCompile(`template\s*<.*>`), patternType: keyword},
	{expression: regexp.MustCompile(`std::\w+`), patternType: keywordOther},
	{expression: regexp.MustCompile(`(cout|cin|endl)`), patternType: keywordPrint},
	{expression: regexp.MustCompile(`(public|protected|private):`), patternType: keywordVisibility},
	{expression: regexp.MustCompile(`nullptr`), patternType: keyword},
	{expression: regexp.MustCompile(`new \w+(\(.*\))?`), patternType: keyword},
	{expression: regexp.MustCompile(`#define\s+.+`), patternType: macro},

	// template usage
	{expression: regexp.MustCompile(`\w+<\w+>`), patternType: keywordOther},

	{expression: regexp.MustCompile(`class\s+\w+`), patternType: keyword},
	{expression: regexp.MustCompile(`void`), patternType: keyword},
	{expression: regexp.MustCompile(`(else )?if\s*\(.+\)`), patternType: keywordControl},
	{expression: regexp.MustCompile(`while\s+\(.+\)`), patternType: keywordControl},

	// scope operator
	{expression: regexp.MustCompile(`\w*::\w+`), patternType: macro},

	// single quote with multichar
	{expression: regexp.MustCompile(`'.{2,}'`), patternType: not},

	// Java List/ArrayList
	{expression: regexp.MustCompile(`(List<\w+>|ArrayList<\w*>\s*\(.*\))(\s+[\w]+|;)`), patternType: not},

	// avoiding Ruby confusion
	{expression: regexp.MustCompile(`def\s+\w+\s*(\(.+\))?\s*\n`), patternType: not},
	{expression: regexp.MustCompile(`puts\s+("|').+("|')`), patternType: not},
	{expression: regexp.MustCompile(`\bmodule\s\S`), patternType: not},

	// avoiding C# confusion
	{expression: regexp.MustCompile(`Console\.(WriteLine|Write)(\s*)?\(`), patternType: not},
	{expression: regexp.MustCompile(`(using\s)?System(\..*)?(;)?`), patternType: not},
	{expression: regexp.MustCompile(`static\s+\S+\s+Main\(.*\)`), patternType: not},
	{expression: regexp.MustCompile(`(public|private|protected|internal)\s`), patternType: not},

	// avoiding Kotlin confusion
	{expression: regexp.MustCompile(`fun main\((.*)?\) \{`), patternType: not},
	{
		expression:  regexp.MustCompile(`(inline|private|public|protected|override|operator(\s+))?fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)(\{|=)`),
		patternType: not,
	},
	{expression: regexp.MustCompile(`(const)?(\s+)?val(\s+)(.*)(:(\s)(.*)(\?)?)?(\s+)=(\s+)`), patternType: not},

	// avoiding Dart confusion
	{expression: regexp.MustCompile(`^(void\s)?main\(\)\s(async\s)?\{`), patternType: not},
}
