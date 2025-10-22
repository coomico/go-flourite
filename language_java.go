package flourite

var java = []languagePattern{
	{expression: `System\.(in|out)\.\w+`, patternType: keywordPrint},

	// class variable declarations
	{expression: `(private|protected|public)\s*\w+\s*\w+(\s*=\s*[\w])?`, patternType: keyword},

	// method
	{expression: `(private|protected|public)\s*\w+\s*[\w]+\(.+\)`, patternType: keyword},

	// string class
	{expression: `(^|\s)(String)\s+[\w]+\s*=?`, patternType: keywordOther},

	// List/ArrayList
	{expression: `(List<\w+>|ArrayList<\w*>\s*\(.*\))(\s+[\w]+|;)`, patternType: keywordVariable},

	// class keyword
	{expression: `(public\s*)?class\b.*?\{`, patternType: keyword},

	// array declaration
	{expression: `(\w+)(\[\s*\])+\s+\w+`, patternType: constantArray},

	{expression: `final\s*\w+`, patternType: keywordOther},
	{expression: `\w+\.(get|set)\(.+\)`, patternType: keywordOther},
	{expression: `new [A-Z]\w*\s*\(.+\)`, patternType: keywordOther},
	{expression: `(^|\s)(char|long|int|float|double)\s+[\w]+\s*=?`, patternType: constantType},
	{expression: `(extends|implements)`, patternType: metaModule, nearTop: true},
	{expression: `null`, patternType: keywordOther},
	{expression: `(else )?if\s*\(.+\)`, patternType: keywordControl},
	{expression: `while\s+\(.+\)`, patternType: keywordControl},
	{expression: `void`, patternType: keywordOther},

	{expression: `const\s*\w+`, patternType: not},
	{expression: `(\w+)\s*\*\s*\w+`, patternType: not},

	// single quote multichar
	{expression: `'.{2,}'`, patternType: not},

	// C style include
	{expression: `#include\s*(<|")\w+(\.h)?(>|")`, patternType: not, nearTop: true},

	// avoiding Ruby confusion
	{expression: `def\s+\w+\s*(\(.+\))?\s*\n`, patternType: not},

	// avoiding C# confusion
	{expression: `\bnamespace\s.*(\s{)?`, patternType: not},
	{expression: `\[Attribute\]`, patternType: not},
	{expression: `Console\.(WriteLine|Write)(\s*)?\(`, patternType: not},
	{expression: `(#region(\s.*)?|#endregion\n)`, patternType: not},
	{expression: `using\sSystem(\..*)?(;)?`, patternType: not},

	// avoiding Kotlin confusion
	{expression: `fun main\((.*)?\) \{`, patternType: not},
	{expression: `(inline(\s+))?fun(\s+)([A-Za-z0-9_])(\s+)?\((.*)\)(\s+)(\{|=)`, patternType: not},
	{expression: `(const)?(\s+)?val(\s+)(.*)(:(\s)(.*)(\?)?)?(\s+)=(\s+)`, patternType: not},

	// avoiding Dart confusion
	{expression: `^(void\s)?main\(\)\s\{`, patternType: not},
}
