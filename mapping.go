package flourite

type LangKind uint

const (
	Unknown LangKind = iota
	Go
	C
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
	Go:         "Go",
	C:          "C",
	JavaScript: "JavaScript",
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
