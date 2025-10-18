package flourite

var clojure = []languagePattern{
	{expression: `^(\s+)?\(ns(\s+)(.*)(\))?$`, patternType: metaModule},
	{expression: `^(\s+)?\(print(ln)?(\s+)(.*)(\))$`, patternType: keywordPrint},
	{expression: `^(\s+)?\((de)?fn(-)?(\s+)(.*)(\))?$`, patternType: keywordFunction},
	{expression: `^(\s+)?\((let|def)(\s+)(.*)(\))?$`, patternType: keywordVariable},
	{expression: `^(\s+)?\((do|if|loop|cond|when|or|and|condp|case)`, patternType: keywordControl},

	// collections and sequences
	{
		expression:  `^(\s+)?\((class|coll\?|seq\?|range|cons|conj|concat|map|filter|reduce)(\s+)(.*)(\))?$`,
		patternType: keyword,
	},

	// threading macro
	{expression: `^(\s+)?\((as)?->(>)?`, patternType: macro},

	// modules
	{expression: `^(\s+)?\((use|require|import|:import)(\s+)(.*)(\))?$`, patternType: metaModule},
}
