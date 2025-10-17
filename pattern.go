package flourite

type PatternType uint8

const (
	// Not is a pattern that do not exist in the current language
	Not PatternType = iota

	// Macro is a pattern likes @println()
	Macro

	// MetaImport is a importing pattern
	// such as import(), require(), import "", #include <>
	MetaImport

	// MetaModule is a module or package pattern
	// such as package <name>, module <name>
	MetaModule

	// SectionScope is a scope pattern
	// such as begin & end
	SectionScope

	// Keyword is a universal keyword pattern
	// such as namespace and class
	Keyword

	// KeywordPrint is a print keyword pattern
	// such as println, echo, console.log(), System.WriteLine()
	KeywordPrint

	// KeywordFunction is a function keyword pattern
	// such as func name() {, function name() {, fn name() {, def name:
	KeywordFunction

	// KeywordVariable is a variable keyword pattern
	// such as var, let, const
	KeywordVariable

	// KeywordOperator is an operator keyword pattern
	// such as >, <, -, <<
	KeywordOperator

	// KeywordControl is a control keyword pattern
	// such as if, while, for, return, break, continue
	KeywordControl

	// KeywordVisibility is a visibility keyword pattern
	// such as public, private, protected
	KeywordVisibility

	// KeywordOther is another keyword pattern designed for specific purpose
	// such as async, await, crate, extern
	KeywordOther

	// ConstantNull is a null constant pattern
	// such as null, undefined, nil
	ConstantNull

	// ConstantType is a type constant pattern
	ConstantType

	// ConstantString is a string constant pattern
	ConstantString

	// ConstantNumeric is a numeric constant pattern
	// such as int, uint, float, double
	ConstantNumeric

	// ConstantBoolean is a boolean constant pattern
	// such as true, True, false, False
	ConstantBoolean

	// ConstantDictionary is a dictionary constant pattern
	// such as dict, object, associate array
	ConstantDictionary

	// ConstantArray is an array constant pattern
	ConstantArray

	// CommentBlock is multiline of comments
	CommentBlock

	// CommentLine is a single block of comment
	CommentLine

	// CommentDocumentation indicates documentation
	CommentDocumentation
)

type LanguagePattern struct {
	Type    PatternType
	Pattern string
	NearTop bool
}
