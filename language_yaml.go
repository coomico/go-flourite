package flourite

import "regexp"

var yaml = []languagePattern{
	// regular key: value
	{expression: regexp.MustCompile(`^( )*([A-Za-z0-9_. ]+):( )?(.*)?$`), patternType: keyword},

	// regular array - key: value
	{expression: regexp.MustCompile(`^( )*-( )([A-Za-z0-9_. ]+):( )?(.*)?$`), patternType: keyword},

	// regular array - value
	{expression: regexp.MustCompile(`^( )*-( )(.*)$`), patternType: keyword},

	// binary tag
	{expression: regexp.MustCompile(`^( )*([A-Za-z0-9_. ]+):( )!!binary( )?(|)?$`), patternType: constantType},

	// literal multiline block
	{expression: regexp.MustCompile(`^( )*([A-Za-z0-9_. ]+):( )\|$`), patternType: keyword},

	// folded multiline style
	{expression: regexp.MustCompile(`^( )*([A-Za-z0-9_. ]+):( )>$`), patternType: keyword},

	// set types
	{expression: regexp.MustCompile(`^( )*\?( )(.*)$`), patternType: keyword},

	// complex key / multiline key
	{expression: regexp.MustCompile(`^( )*\?( )\|$`), patternType: constantType},

	// merge key
	{expression: regexp.MustCompile(`^( )*<<:( )(\*)(.*)?$`), patternType: constantType},

	// avoiding CSS confusion
	{expression: regexp.MustCompile(`^( )*([A-Za-z0-9_. ]+):(.*)?( )?{$`), patternType: not},
	{expression: regexp.MustCompile(`^( )*([A-Za-z0-9_. ]+):(.*)?( )?,$`), patternType: not},
}
