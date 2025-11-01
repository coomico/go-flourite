package flourite

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var blackhole string

func BenchmarkMultiline(b *testing.B) {
	samples, err := os.ReadDir("testdata/lines")
	if err != nil {
		b.Error(err)
	}

	for _, sample := range samples {
		filename := sample.Name()
		contents, err := os.ReadFile(filepath.Join("testdata/lines", filename))
		if err != nil {
			b.Error(err)
		}

		name := strings.Split(filename, ".")[0]
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				blackhole = Detect(string(contents)).Best().Language.String()
				if blackhole != langNames[Go] {
					b.Fatalf("detection error: got %s, want %s", blackhole, langNames[Go])
				}
			}
		})
	}
}
