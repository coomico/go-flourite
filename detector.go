package flourite

import (
	"regexp"
	"strings"
)

type Detector struct {
	IsUnknown bool
	Heuristic bool
}

func (d Detector) Detect(snippet string) detectedLanguages {
	snippet = regexp.MustCompile(`[\r\n]+`).ReplaceAllString(snippet, "\n")
	lines := strings.Split(snippet, "\n")

	// TODO: check heuristic options with lines > 500
	// TODO: shebang check

	results := make(detectedLanguages, 0, len(langNames))
	if d.IsUnknown {
		results = append(results, LangPoint{Unknown, 1})
	}

	for lang, _ := range langNames {
		patterns, ok := getPatterns(LangKind(lang))
		if !ok {
			continue
		}

		var points int
		for i := 0; i < len(lines); i++ {
			// skip for empty line or contains only spaces
			if regexp.MustCompile(`^\s*$`).MatchString(lines[i]) {
				continue
			}

			points += getPoints(lines[i], patterns, nearTop(i, lines))
		}

		results = append(results, LangPoint{LangKind(lang), points})
	}

	return results
}

type detectedLanguages []LangPoint

func (lp detectedLanguages) Best() LangPoint {
	var best LangPoint = lp[0]
	for i := 1; i < len(lp); i++ {
		if lp[i].points >= best.points {
			best = lp[i]
		}
	}
	return best
}

func (lp detectedLanguages) String() string {
	s := strings.Builder{}
	for i, l := range lp {
		if i > 0 {
			s.WriteString("\n")
		}

		s.WriteString(l.String())
	}

	return s.String()
}
