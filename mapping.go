package flourite

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

func getPatterns(lang LangKind) ([]languagePattern, bool) {
	switch lang {
	case C:
		return cLang, true
	case Clojure:
		return clojure, true
	case Cpp:
		return cpp, true
	case CS:
		return cs, true
	case CSS:
		return css, true
	case Dart:
		return dart, true
	case Dockerfile:
		return dockerfile, true
	case Elixir:
		return elixir, true
	case Go:
		return goLang, true
	case HTML:
		return html, true
	case Java:
		return java, true
	case JavaScript:
		return javascript, true
	case JSON:
		return json, true
	case Julia:
		return julia, true
	case Kotlin:
		return kotlin, true
	case Lua:
		return lua, true
	case Markdown:
		return markdown, true
	case Pascal:
		return pascal, true
	case PHP:
		return php, true
	case Python:
		return python, true
	case Ruby:
		return ruby, true
	case Rust:
		return rust, true
	case SQL:
		return sql, true
	case TypeScript:
		return typescript, true
	case YAML:
		return yaml, true
	default:
		return nil, false
	}
}
