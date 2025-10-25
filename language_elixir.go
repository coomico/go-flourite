package flourite

import "regexp"

var elixir = []languagePattern{
	{expression: regexp.MustCompile(`^\s*defmodule\s+\S+(\s+do)?$`), patternType: metaModule},
	{expression: regexp.MustCompile(`\s*alias\s+.+as:.+`), patternType: keywordOther},
	{expression: regexp.MustCompile(`IO\.(puts|inspect|write)`), patternType: keywordPrint},

	// anonymous func
	{expression: regexp.MustCompile(`fn\s+.+\s+->`), patternType: keywordFunction},

	{expression: regexp.MustCompile(`^\s*(def|defp)\s+\w+`), patternType: keywordFunction},
	{
		expression:  regexp.MustCompile(`^\s*(if|unless|cond|case|try|defimpl|defprotocol)\s+`),
		patternType: keywordControl,
	},
	{expression: regexp.MustCompile(`^\s*defstruct\s+`), patternType: keyword},
	{expression: regexp.MustCompile(`^\s*@spec\s+.+::.+`), patternType: macro},
	{expression: regexp.MustCompile(`\{[^}]*,.*\}`), patternType: constantArray},
	{expression: regexp.MustCompile(`%\{(.+(=>|:).+(,)?)+\}`), patternType: constantDictionary},
	{expression: regexp.MustCompile(`^\s*end\s*$`), patternType: keyword},

	// pipe operator
	{expression: regexp.MustCompile(`\|>`), patternType: keywordOperator},
}
