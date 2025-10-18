package flourite

type LangKind uint

const (
	Unknown LangKind = iota
	// should be sorted alphabetically
	C
	Clojure
	Go
	JavaScript
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
	Go:         "Go",
	JavaScript: "JavaScript",
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
	case Go:
		return goLang, true
	case C:
		return cLang, true
	case Clojure:
		return clojure, true
	case JavaScript:
		return js, true
	default:
		return nil, false
	}
}
