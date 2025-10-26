package flourite

import "regexp"

var java = []languagePattern{
	{expression: regexp.MustCompile(`^\s*import\s+[\w.]+(;|\s|$)`), patternType: metaImport, nearTop: true},
	{expression: regexp.MustCompile(`System\.(in|out)\.\w+`), patternType: keywordPrint},

	// class variable declarations
	{expression: regexp.MustCompile(`(private|protected|public)\s*\w+\s*\w+(\s*=\s*[\w])?`), patternType: keyword},

	// method
	{
		expression:  regexp.MustCompile(`\b(private|protected|public)\s+(static\s+)?(\w+(<.*>)?|\w+\[\])\s+\w+\s*\([^)]*\)\s*(throws\s+\w+)?\s*\{`),
		patternType: keyword,
	},

	// string class
	{expression: regexp.MustCompile(`(^|\s)(String)\s+[\w]+\s*=?`), patternType: keywordOther},

	// List/ArrayList
	{
		expression:  regexp.MustCompile(`\b(List|ArrayList|LinkedList|Vector|Stack)<[^>]+>\s+\w+`),
		patternType: keywordVariable,
	},

	{
		expression:  regexp.MustCompile(`\bnew\s+(ArrayList|HashMap|HashSet|LinkedList|TreeMap|TreeSet)<[^>]*>\s*\(`),
		patternType: keywordOther,
	},

	// class keyword
	{expression: regexp.MustCompile(`(public\s+)?class\s+\w+(\s+extends|\s+implements|\s*\{)`), patternType: keyword},

	// array declaration
	{expression: regexp.MustCompile(`(\w+)(\[\s*\])+\s+\w+`), patternType: constantArray},

	{expression: regexp.MustCompile(`final\s*\w+`), patternType: keywordOther},
	{expression: regexp.MustCompile(`\w+\.(get|set)\(.+\)`), patternType: keywordOther},
	{expression: regexp.MustCompile(`\bnew\s+[A-Z]\w*\s*\(`), patternType: keywordOther},
	{
		expression:  regexp.MustCompile(`\b(char|byte|short|long|int|float|double|boolean)\s+\w+\s*(=|;)`),
		patternType: constantType,
	},
	{expression: regexp.MustCompile(`\b(extends|implements)\s+\w+`), patternType: metaModule, nearTop: true},
	{expression: regexp.MustCompile(`null`), patternType: keywordOther},
	{expression: regexp.MustCompile(`(else )?if\s*\(.+\)`), patternType: keywordControl},
	{expression: regexp.MustCompile(`while\s+\(.+\)`), patternType: keywordControl},
	{expression: regexp.MustCompile(`void`), patternType: keywordOther},

	// Java annotations
	{
		expression:  regexp.MustCompile(`@(Override|Deprecated|SuppressWarnings|FunctionalInterface|SafeVarargs)\b`),
		patternType: keywordOther,
	},

	{expression: regexp.MustCompile(`const\s*\w+`), patternType: not},
	{expression: regexp.MustCompile(`(\w+)\s*\*\s*\w+`), patternType: not},

	// single quote multichar
	{expression: regexp.MustCompile(`'.{2,}'`), patternType: not},

	// C style include
	{expression: regexp.MustCompile(`#include\s*(<|")\w+(\.h)?(>|")`), patternType: not, nearTop: true},

	// avoiding Ruby confusion
	{expression: regexp.MustCompile(`def\s+\w+\s*(\(.+\))?\s*\n`), patternType: not},

	// avoiding C# confusion
	{expression: regexp.MustCompile(`\bnamespace\s.*(\s{)?`), patternType: not},
	{expression: regexp.MustCompile(`\[Attribute\]`), patternType: not},
	{expression: regexp.MustCompile(`Console\.(WriteLine|Write)(\s*)?\(`), patternType: not},
	{expression: regexp.MustCompile(`(#region(\s.*)?|#endregion\n)`), patternType: not},
	{expression: regexp.MustCompile(`using\sSystem(\..*)?(;)?`), patternType: not},

	// avoiding Kotlin confusion
	{expression: regexp.MustCompile(`fun main\((.*)?\) \{`), patternType: not},
	{expression: regexp.MustCompile(`(inline(\s+))?fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)(\{|=)`), patternType: not},
	{expression: regexp.MustCompile(`(const)?(\s+)?val(\s+)(.*)(:(\s)(.*)(\?)?)?(\s+)=(\s+)`), patternType: not},

	// avoiding Dart confusion
	{expression: regexp.MustCompile(`^(void\s)?main\(\)\s\{`), patternType: not},
}
