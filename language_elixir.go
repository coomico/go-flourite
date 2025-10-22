package flourite

var elixir = []languagePattern{
	{expression: `^\s*defmodule\s+.+\s+do$`, patternType: metaModule},
	{expression: `\s*alias\s+.+as:.+`, patternType: keywordOther},
	{expression: `IO\.puts.+`, patternType: keywordPrint},

	// anonymous func
	{expression: `fn\s+[A-Za-z0-9_:<>()]+\s+->\s+.+(end)?$`, patternType: keywordFunction},

	{expression: `^\s*(def|defp)\s+.+\s+do$`, patternType: keywordFunction},
	{expression: `^\s*(if|unless|cond|case|try|defimpl|defprotocol)\s+.+\s+do$`, patternType: keywordControl},
	{expression: `^\s*defstruct\s+`, patternType: keyword},
	{expression: `^\s*@spec\s+.+::.+`, patternType: macro},
	{expression: `\{:.+,.+\}`, patternType: constantArray},
	{expression: `%\{(.+(=>|:).+(,)?)+\}`, patternType: constantDictionary},
}
