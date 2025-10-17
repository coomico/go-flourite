package flourite

import (
	"regexp"
	"strings"
)

type Detector struct {
	IsUnknown bool
	Heuristic bool
}

func (d Detector) Detect(snippet string) DetectedLanguages {
	snippet = regexp.MustCompile(`[\r\n]+`).ReplaceAllString(snippet, "\n")
	lines := strings.Split(snippet, "\n")

	// TODO: check heuristic options with lines > 500
	// TODO: shebang check

	results := make(DetectedLanguages, 0, len(langNames))
	if d.IsUnknown {
		results = append(results, LangPoint{Unknown, 1})
	}

	for i := 0; i < len(langNames); i++ {
		patterns, ok := getPatterns(LangKind(i))
		if !ok {
			continue
		}

		var points int
		for i := 0; i < len(lines); i++ {
			// skip for empty line or contains only spaces
			if regexp.MustCompile(`^\s*$`).MatchString(lines[i]) {
				continue
			}

			points += getPoints(lines[i], patterns, isNearTop(i, lines))
		}

		results = append(results, LangPoint{LangKind(i), points})
	}

	return results
}

type DetectedLanguages []LangPoint

func (lp DetectedLanguages) Best() LangPoint {
	var best LangPoint = lp[0]
	for i := 1; i < len(lp); i++ {
		if lp[i].Points >= best.Points {
			best = lp[i]
		}
	}
	return best
}

func (lp DetectedLanguages) String() string {
	s := strings.Builder{}
	for i, l := range lp {
		if i > 0 {
			s.WriteString("\n")
		}

		s.WriteString(l.String())
	}

	return s.String()
}
