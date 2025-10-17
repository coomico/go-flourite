package flourite

type LangKind uint

const (
	Unknown LangKind = iota
	Go
	C
)

func (l LangKind) String() string {
	if uint(l) < uint(len(langNames)) {
		return langNames[l]
	}
	return ""
}

var langNames = []string{
	Unknown: "Unknown",
	Go:      "Go",
	C:       "C",
}

func getPatterns(lang LangKind) ([]LanguagePattern, bool) {
	switch lang {
	case Go:
		return goLang, true
	case C:
		return c, true
	default:
		return nil, false
	}
}
