package flourite

type patternType uint8

const (
	// not is a pattern that do not exist in the current language
	not patternType = iota

	// macro is a pattern likes @println()
	macro

	// metaImport is an importing pattern
	// such as import(), require(), import "", #include <>
	metaImport

	// metaModule is a module or package pattern
	// such as package <name>, module <name>
	metaModule

	// sectionScope is a scope pattern
	// such as begin & end
	sectionScope

	// keyword is a universal keyword pattern
	// such as namespace and class
	keyword

	// keywordPrint is a print keyword pattern
	// such as println, echo, console.log(), System.WriteLine()
	keywordPrint

	// keywordFunction is a function keyword pattern
	// such as func name() {, function name() {, fn name() {, def name:
	keywordFunction

	// keywordVariable is a variable keyword pattern
	// such as var, let, const
	keywordVariable

	// keywordOperator is an operator keyword pattern
	// such as >, <, -, <<
	keywordOperator

	// keywordControl is a control keyword pattern
	// such as if, while, for, return, break, continue
	keywordControl

	// keywordVisibility is a visibility keyword pattern
	// such as public, private, protected
	keywordVisibility

	// keywordOther is another keyword pattern designed for specific purpose
	// such as async, await, crate, extern
	keywordOther

	// constantNull is a null constant pattern
	// such as null, undefined, nil
	constantNull

	// constantType is a type constant pattern
	constantType

	// constantString is a string constant pattern
	constantString

	// constantNumeric is a numeric constant pattern
	// such as int, uint, float, double
	constantNumeric

	// ConstantBoolean is a boolean constant pattern
	// such as true, True, false, False
	constantBoolean

	// constantDictionary is a dictionary constant pattern
	// such as dict, object, associate array
	constantDictionary

	// constantArray is an array constant pattern
	constantArray

	// commentBlock is multiline of comments
	commentBlock

	// commentLine is a single block of comment
	commentLine

	// commentDocumentation indicates documentation
	commentDocumentation
)

type languagePattern struct {
	patternType patternType
	expression  string
	nearTop     bool
}
