package flourite

var yaml = []languagePattern{
	// regular key: value
	{expression: `^( )*([A-Za-z0-9_. ]+):( )?(.*)?$`, patternType: keyword},

	// regular array - key: value
	{expression: `^( )*-( )([A-Za-z0-9_. ]+):( )?(.*)?$`, patternType: keyword},

	// regular array - value
	{expression: `^( )*-( )(.*)$`, patternType: keyword},

	// binary tag
	{expression: `^( )*([A-Za-z0-9_. ]+):( )!!binary( )?(|)?$`, patternType: constantType},

	// literal multiline block
	{expression: `^( )*([A-Za-z0-9_. ]+):( )\|$`, patternType: keyword},

	// folded multiline style
	{expression: `^( )*([A-Za-z0-9_. ]+):( )>$`, patternType: keyword},

	// set types
	{expression: `^( )*\?( )(.*)$`, patternType: keyword},

	// complex key / multiline key
	{expression: `^( )*\?( )\|$`, patternType: constantType},

	// merge key
	{expression: `^( )*<<:( )(\*)(.*)?$`, patternType: constantType},

	// avoiding CSS confusion
	{expression: `^( )*([A-Za-z0-9_. ]+):(.*)?( )?{$`, patternType: not},
	{expression: `^( )*([A-Za-z0-9_. ]+):(.*)?( )?,$`, patternType: not},
}
