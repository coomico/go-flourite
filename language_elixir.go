package flourite

import "regexp"

var elixir = []languagePattern{
	{expression: regexp.MustCompile(`^\s*defmodule\s+.+\s+do$`), patternType: metaModule},
	{expression: regexp.MustCompile(`\s*alias\s+.+as:.+`), patternType: keywordOther},
	{expression: regexp.MustCompile(`IO\.puts.+`), patternType: keywordPrint},

	// anonymous func
	{expression: regexp.MustCompile(`fn\s+[A-Za-z0-9_:<>()]+\s+->\s+.+(end)?$`), patternType: keywordFunction},

	{expression: regexp.MustCompile(`^\s*(def|defp)\s+.+\s+do$`), patternType: keywordFunction},
	{expression: regexp.MustCompile(`^\s*(if|unless|cond|case|try|defimpl|defprotocol)\s+.+\s+do$`), patternType: keywordControl},
	{expression: regexp.MustCompile(`^\s*defstruct\s+`), patternType: keyword},
	{expression: regexp.MustCompile(`^\s*@spec\s+.+::.+`), patternType: macro},
	{expression: regexp.MustCompile(`\{:.+,.+\}`), patternType: constantArray},
	{expression: regexp.MustCompile(`%\{(.+(=>|:).+(,)?)+\}`), patternType: constantDictionary},
}
