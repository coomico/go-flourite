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
	Python
	Shell
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
	Python:     "Python",
	Shell:      "Shell",
}

var interpreterMap = map[string]LangKind{
	"node":   JavaScript,
	"jsc":    JavaScript,
	"python": Python,
	"bash":   Shell,
	"ash":    Shell,
	"zsh":    Shell,
	"sh":     Shell,
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
		return js, true
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
	default:
		return nil, false
	}
}
