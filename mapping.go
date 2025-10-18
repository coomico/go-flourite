package flourite

type LangKind uint

const (
	Unknown LangKind = iota
	// should be sorted alphabetically
	Bash
	C
	Go
	JavaScript
)

func (l LangKind) String() string {
	if uint(l) < uint(len(langNames)) {
		return langNames[l]
	}
	return ""
}

var langNames = []string{
	Unknown:    "Unknown",
	Bash:       "Bash",
	C:          "C",
	Go:         "Go",
	JavaScript: "JavaScript",
}

var shebangMap = map[string]LangKind{
	"node": JavaScript,
	"jsc":  JavaScript,
}

func getPatterns(lang LangKind) ([]languagePattern, bool) {
	switch lang {
	case Go:
		return goLang, true
	case C:
		return cLang, true
	case JavaScript:
		return js, true
	default:
		return nil, false
	}
}
