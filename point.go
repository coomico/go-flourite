package flourite

import (
	"regexp"
	"strconv"
)

type LangPoint struct {
	language LangKind
	points   int
}

func (lp LangPoint) String() string {
	return lp.language.String() + "=" + strconv.Itoa(lp.points)
}

func parsePoint(p PatternType) int {
	switch p {
	case KeywordPrint, MetaImport, MetaModule:
		return 5
	case KeywordFunction, ConstantNull:
		return 4
	case ConstantType,
		ConstantString,
		ConstantNumeric,
		ConstantBoolean,
		ConstantDictionary,
		ConstantArray,
		KeywordVariable:
		return 3
	case SectionScope,
		KeywordOther,
		KeywordOperator,
		KeywordControl,
		KeywordVisibility,
		Keyword:
		return 2
	case CommentBlock,
		CommentLine,
		CommentDocumentation,
		Macro:
		return 1
	case Not:
		return 0
	default:
		return -20
	}
}

func getPoints(line string, patterns []LanguagePattern, isTop bool) int {
	var points int
	for _, languagePattern := range patterns {
		if languagePattern.NearTop != isTop {
			continue
		}

		re := regexp.MustCompile(languagePattern.Pattern)
		if valid := re.Match([]byte(line)); valid {
			points += parsePoint(languagePattern.Type)
		}
	}

	return points
}

func nearTop(i int, lines []string) bool {
	if len(lines) <= 10 {
		return true
	}

	return i < len(lines)/10
}
