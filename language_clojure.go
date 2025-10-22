package flourite

import "regexp"

var clojure = []languagePattern{
	{expression: regexp.MustCompile(`^(\s+)?\(ns(\s+)(.*)(\))?$`), patternType: metaModule},
	{expression: regexp.MustCompile(`^(\s+)?\(print(ln)?(\s+)(.*)(\))$`), patternType: keywordPrint},
	{expression: regexp.MustCompile(`^(\s+)?\((de)?fn(-)?(\s+)(.*)(\))?$`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`^(\s+)?\((let|def)(\s+)(.*)(\))?$`), patternType: keywordVariable},
	{expression: regexp.MustCompile(`^(\s+)?\((do|if|loop|cond|when|or|and|condp|case)`), patternType: keywordControl},

	// collections and sequences
	{
		expression:  regexp.MustCompile(`^(\s+)?\((class|coll\?|seq\?|range|cons|conj|concat|map|filter|reduce)(\s+)(.*)(\))?$`),
		patternType: keyword,
	},

	// threading macro
	{expression: regexp.MustCompile(`^(\s+)?\((as)?->(>)?`), patternType: macro},

	// modules
	{expression: regexp.MustCompile(`^(\s+)?\((use|require|import|:import)(\s+)(.*)(\))?$`), patternType: metaModule},
}
