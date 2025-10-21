package flourite

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetector(t *testing.T) {
	langs, err := os.ReadDir("testdata/Languages")
	if err != nil {
		t.Error(err)
	}

	detector := Detector{}

	for _, langDir := range langs {
		if langDir.IsDir() {
			langName := langDir.Name()
			dirname := filepath.Join("testdata/Languages", langName)
			files, err := os.ReadDir(dirname)
			if err != nil {
				t.Error(err)
			}

			for _, file := range files {
				filename := file.Name()
				path := filepath.Join(dirname, filename)
				t.Run(
					langName+"/"+filename, func(t *testing.T) {
						t.Parallel()
						sample, err := os.ReadFile(path)
						if err != nil {
							t.Error(err)
						}

						detected := detector.Detect(string(sample))
						got := detected.Best().Language.String()
						if got != langName {
							t.Errorf("Language detected got %s, want %s", got, langName)
							t.Error(detected)
						}
					},
				)
			}
		}
	}
}
