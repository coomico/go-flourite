package flourite

import (
	"math"
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

	// shebang check
	if strings.Contains(lines[0], "#!") {
		if strings.Contains(lines[0], "#!/bin/bash") {
			return DetectedLanguages{
				{Bash, 5},
			}
		}

		if strings.Contains(lines[0], "#!/usr/bin/env") {
			interpreter := strings.Split(lines[0], " ")[1]
			if lang, ok := shebangMap[interpreter]; ok {
				return DetectedLanguages{
					{lang, 5},
				}
			}

			// TODO: pass on the given interpreter name
			return DetectedLanguages{
				{Unknown, 1},
			}
		}
	}

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
		if i > 0 && i < len(lp) {
			s.WriteString(";")
		}

		s.WriteString(l.String())
	}

	return s.String()
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
