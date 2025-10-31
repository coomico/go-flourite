package flourite

// A LangKind represents a programming languages. The zero LangKind represent an Unknown language
type LangKind uint

const (
	Unknown LangKind = iota
	// should be sorted alphabetically
	C
	Clojure
	Cpp
	CS
	CSS
	Dart
	Dockerfile
	Elixir
	Go
	HTML
	Java
	JavaScript
	JSON
	Julia
	Kotlin
	Lua
	Markdown
	Pascal
	Perl
	PHP
	Python
	Ruby
	Rust
	Shell
	SQL
	TypeScript
	YAML
)

// String return the name of l (programming language).
func (l LangKind) String() string {
	if uint(l) < uint(len(langNames)) {
		return langNames[l]
	}
	return ""
}

var langNames = []string{
	Unknown:    "Unknown",
	C:          "C",
	Clojure:    "Clojure",
	Cpp:        "C++",
	CS:         "C#",
	CSS:        "CSS",
	Dart:       "Dart",
	Dockerfile: "Dockerfile",
	Elixir:     "Elixir",
	Go:         "Go",
	HTML:       "HTML",
	Java:       "Java",
	JavaScript: "JavaScript",
	JSON:       "JSON",
	Julia:      "Julia",
	Kotlin:     "Kotlin",
	Lua:        "Lua",
	Markdown:   "Markdown",
	Pascal:     "Pascal",
	Perl:       "Perl",
	PHP:        "PHP",
	Python:     "Python",
	Ruby:       "Ruby",
	Rust:       "Rust",
	Shell:      "Shell",
	SQL:        "SQL",
	TypeScript: "TypeScript",
	YAML:       "YAML",
}

// extracted from https://github.com/github-linguist/linguist/blob/main/lib/linguist/languages.yml
var interpreterMap = map[string]LangKind{
	"tcc":         C,
	"elixir":      Elixir,
	"chakra":      JavaScript,
	"d8":          JavaScript,
	"gjs":         JavaScript,
	"js":          JavaScript,
	"jsc":         JavaScript,
	"node":        JavaScript,
	"nodejs":      JavaScript,
	"rhino":       JavaScript,
	"qjs":         JavaScript,
	"v8":          JavaScript,
	"v8-shell":    JavaScript,
	"julia":       Julia,
	"instantfpc":  Pascal,
	"perl":        Perl,
	"php":         PHP,
	"python":      Python,
	"py":          Python,
	"pypy":        Python,
	"jruby":       Ruby,
	"macruby":     Ruby,
	"rake":        Ruby,
	"rbx":         Ruby,
	"ruby":        Ruby,
	"rust-script": Rust,
	"bash":        Shell,
	"ash":         Shell,
	"ksh":         Shell,
	"sh":          Shell,
	"zsh":         Shell,
	"deno":        TypeScript,
	"ts-node":     TypeScript,
	"tsx":         TypeScript,
}

var patternMap = map[LangKind][]languagePattern{
	C:          cLang,
	Clojure:    clojure,
	Cpp:        cpp,
	CS:         cs,
	CSS:        css,
	Dart:       dart,
	Dockerfile: dockerfile,
	Elixir:     elixir,
	Go:         goLang,
	HTML:       html,
	Java:       java,
	JavaScript: javascript,
	JSON:       json,
	Julia:      julia,
	Kotlin:     kotlin,
	Lua:        lua,
	Markdown:   markdown,
	Pascal:     pascal,
	PHP:        php,
	Python:     python,
	Ruby:       ruby,
	Rust:       rust,
	SQL:        sql,
	TypeScript: typescript,
	YAML:       yaml,
}
