// Package flourite is a go port of typescript flourite library (https://github.com/teknologi-umum/flourite),
// provides programming language detection from a given string.
package flourite

import (
	"math"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	reCRNewline = regexp.MustCompile(`[\r\n]+`)
	reEmptyLine = regexp.MustCompile(`^\s*$`)
	reVersion   = regexp.MustCompile(`((?:\d+\.?)+)`)

	// skip for optional and variable args, then find the path and interpreter
	reShebang = regexp.MustCompile(`^#!\s*(\S+)(?:\s+(?:(?:-[i0uCSv]*|--\S+|\S+=\S+)\s+)*(\S+))?.*`)
)

// A Strategy is detection strategy. Its default value ([DefaultStrategy])
// is a usable strategy that enable heuristic optimisation.
type Strategy struct {
	// IsUnknown is an option to enable [Unknown] as a fallback language
	// when no language is detected.
	IsUnknown bool

	// Heuristic is an option to enable heuristic optimization for better performance.
	// Only usefull when the number of lines of code is more than 500.
	Heuristic bool
}

// DefaultStrategy is the default [Strategy] and is used by [Detect].
var DefaultStrategy = Strategy{
	Heuristic: true,
}

// Detect using a given [Strategy] and returns the [DetectedLanguages] containing
// all languages with their points.
//
// If there is a shebang on the first lines of code, it will return [DetectedLanguages]
// containing only language detected by the given interpreter.
func (s Strategy) Detect(snippet string) DetectedLanguages {
	snippet = reCRNewline.ReplaceAllString(snippet, "\n")
	lines := strings.Split(snippet, "\n")

	if s.Heuristic && len(lines) > 500 {
		lines = heuristicOptimization(lines)
	}

	if len(lines) > 0 && strings.Contains(lines[0], "#!") {
		interpreter := getInterpreter(lines[0])
		if lang, ok := interpreterMap[interpreter]; ok {
			return DetectedLanguages{
				{lang, 5},
			}
		}
	}

	results := make(DetectedLanguages, 0, len(langNames))
	if s.IsUnknown {
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
			if reEmptyLine.MatchString(lines[i]) {
				continue
			}

			points += getPoints(lines[i], patterns, isNearTop(i, lines))
		}

		results = append(results, LangPoint{LangKind(i), points})
	}

	return results
}

// Detect is [Strategy.Detect] with [DefaultStrategy].
func Detect(s string) DetectedLanguages {
	return DefaultStrategy.Detect(s)
}

// A DetectedLanguages is an alias for the [LangPoint] list.
type DetectedLanguages []LangPoint

// Best return the [LangPoint] with the most points
func (dl DetectedLanguages) Best() LangPoint {
	best := dl[0]
	for i := 1; i < len(dl); i++ {
		if dl[i].Points > best.Points {
			best = dl[i]
		}
	}
	return best
}

// String return all languages along with their respective points.
func (dl DetectedLanguages) String() string {
	s := strings.Builder{}
	for i, l := range dl {
		if i > 0 && i < len(dl) {
			s.WriteString(" ")
		}

		s.WriteString(l.String())
	}

	return s.String()
}

// acknowledge:
// https://github.com/dayvonjersen/linguist/blob/c82f0abfd1c3a1d6b4c467489292d22ea1907a4f/linguist.go#L131
func getInterpreter(s string) string {
	shebang := reShebang.FindStringSubmatch(s)
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

	return reVersion.ReplaceAllString(base, "")
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
