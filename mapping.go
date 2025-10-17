package flourite

type LangKind uint

const (
	Unknown LangKind = iota
	Go
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
}

func getPatterns(lang LangKind) ([]LanguagePattern, bool) {
	switch lang {
	case Go:
		return goLang, true
	default:
		return nil, false
	}
}
