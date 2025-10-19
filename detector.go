package flourite

import (
	"math"
	"path/filepath"
	"regexp"
	"strings"
)

type detector struct {
	IsUnknown bool
	Heuristic bool
}

func NewDetector() detector {
	return detector{
		Heuristic: true,
	}
}

func (d detector) Detect(snippet string) DetectedLanguages {
	snippet = regexp.MustCompile(`[\r\n]+`).ReplaceAllString(snippet, "\n")
	lines := strings.Split(snippet, "\n")

	if d.Heuristic && len(lines) > 500 {
		lines = heuristicOptimization(lines)
	}

	if strings.Contains(lines[0], "#!") {
		interpreter := interpreterCheck(lines[0])
		if lang, ok := interpreterMap[interpreter]; ok {
			return DetectedLanguages{
				{lang, 5},
			}
		}
	}

	results := make(DetectedLanguages, 0, len(langNames))
	if d.IsUnknown {
		results = append(results, LangPoint{Unknown, 1})
	}

	for i := 1; i < len(langNames); i++ {
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

func (dl DetectedLanguages) Best() LangPoint {
	var best LangPoint = dl[0]
	for i := 1; i < len(dl); i++ {
		if dl[i].Points >= best.Points {
			best = dl[i]
		}
	}
	return best
}

func (dl DetectedLanguages) String() string {
	s := strings.Builder{}
	for i, l := range dl {
		if i > 0 && i < len(dl) {
			s.WriteString(";")
		}

		s.WriteString(l.String())
	}

	return s.String()
}

// acknowledge:
// https://github.com/dayvonjersen/linguist/blob/c82f0abfd1c3a1d6b4c467489292d22ea1907a4f/linguist.go#L131
func interpreterCheck(s string) string {
	shebangExpr := regexp.MustCompile(`^#!\s*(\S+)(?:\s+(\S+))?.*`)
	shebang := shebangExpr.FindStringSubmatch(s)
	if shebang == nil || len(shebang) != 3 {
		return ""
	}

	base := filepath.Base(shebang[1])
	if base == "env" {
		if shebang[2] == "" {
			return ""
		}

		base = shebang[2]
	}

	versionExpr := regexp.MustCompile(`((?:\d+\.?)+)`)
	return versionExpr.ReplaceAllString(base, "")
}

func heuristicOptimization(lines []string) []string {
	length := len(lines)
	out := make([]string, 0, length)

	for i := 0; i < length; i++ {
		// position of the line must be near the top or divisible by ⌈length / 500⌉
		if !isNearTop(i, lines) && i%int(math.Ceil(float64(length)/500)) != 0 {
			continue
		}

		out = append(out, lines[i])
	}
	return out
}
