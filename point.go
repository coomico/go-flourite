package flourite

import (
	"strconv"
)

// A LangPoint describes a single result in a [DetectedLanguages].
type LangPoint struct {
	Language LangKind // the specific language
	Points   int      // is the total points obtained by that language
}

// String return the language name along with its points.
func (lp LangPoint) String() string {
	return lp.Language.String() + ":" + strconv.Itoa(lp.Points)
}

func parsePoint(p patternType) int {
	switch p {
	case keywordPrint, metaImport, metaModule:
		return 5
	case keywordFunction, constantNull:
		return 4
	case constantType,
		constantString,
		constantNumeric,
		constantBoolean,
		constantDictionary,
		constantArray,
		keywordVariable:
		return 3
	case sectionScope,
		keywordOther,
		keywordOperator,
		keywordControl,
		keywordVisibility,
		keyword:
		return 2
	case commentBlock,
		commentLine,
		commentDocumentation,
		macro:
		return 1
	case not:
		return 0
	default:
		return -20
	}
}

func getPoints(line string, patterns []languagePattern, isNearTop bool) int {
	var points int
	for _, languagePattern := range patterns {
		if languagePattern.nearTop && !isNearTop {
			continue
		}

		if valid := languagePattern.expression.MatchString(line); valid {
			points += parsePoint(languagePattern.patternType)
		}
	}

	return points
}

func isNearTop(i int, lines []string) bool {
	if len(lines) <= 10 {
		return true
	}

	return i < len(lines)/10
}
