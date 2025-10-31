package flourite

import (
	"strconv"
	"strings"
)

// A LangPoint describes a single result in a [DetectedLanguages].
type LangPoint struct {
	Language LangKind // the specific language
	Points   int      // is the total points obtained by that language
}

// String return the language name along with its points.
func (lp LangPoint) String() string {
	var sb strings.Builder
	sb.WriteString(lp.Language.String())
	sb.WriteByte(':')
	sb.WriteString(strconv.Itoa(lp.Points))
	return sb.String()
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

func getPoints(content string, isNearTop bool, patterns []languagePattern) int {
	var points int
	for _, pattern := range patterns {
		if pattern.nearTop && !isNearTop {
			continue
		}

		if valid := pattern.expression.MatchString(content); valid {
			points += parsePoint(pattern.patternType)
		}
	}

	return points
}

func isNearTop(i, length int) bool {
	if length <= 10 {
		return true
	}

	return i < length/10
}
