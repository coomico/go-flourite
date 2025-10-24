package flourite

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetect(t *testing.T) {
	langs, err := os.ReadDir("testdata/languages")
	if err != nil {
		t.Error(err)
	}

	detector := Detector{}

	for _, langDir := range langs {
		if langDir.IsDir() {
			langName := langDir.Name()
			dirname := filepath.Join("testdata/languages", langName)
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
							t.Errorf("detection error: got %s, want %s", got, langName)
							t.Error(detected)
						}
					},
				)
			}
		}
	}
}

func TestDetectWithHeuristicOpt(t *testing.T) {
	langs, err := os.ReadDir("testdata/languages")
	if err != nil {
		t.Error(err)
	}

	detector := Detector{
		Heuristic: true,
	}

	for _, langDir := range langs {
		if langDir.IsDir() {
			langName := langDir.Name()
			dirname := filepath.Join("testdata/languages", langName)
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
							t.Errorf("detection error: got %s, want %s", got, langName)
							t.Error(detected)
						}
					},
				)
			}
		}
	}
}

func TestDetectInterpreter(t *testing.T) {
	testcases := []struct {
		name    string
		snippet string
		lang    LangKind
	}{
		{"js-env-node", "#!/usr/bin/env node", JavaScript},
		{"julia-env", "#!/usr/bin/env julia", Julia},
		{"perl-local", "#!/usr/local/bin/perl", Perl},
		{"perl-w/o-path", "#! perl", Perl},
		{"perl", "#!/usr/bin/perl", Perl},
		{"php-env", "#!/usr/bin/env php", PHP},
		{"php", "#!/usr/bin/php", PHP},
		{"python-env", "#!/usr/bin/env python", Python},
		{"python-env-w-version", "#!/usr/bin/env python2.4", Python},
		{"python", "#!/usr/bin/python", Python},
		{"ruby-jruby", "#!/usr/bin/env jruby", Ruby},
		{"ruby-macruby", "#!/usr/bin/env macruby", Ruby},
		{"ruby-rake", "#!/usr/bin/env rake", Ruby},
		{"ruby-rbx", "#!/usr/bin/env rbx", Ruby},
		{"ruby-env", "#!/usr/bin/env ruby", Ruby},
		{"ruby-w-args", "#! /usr/bin/env ruby -w -Ilib:test", Ruby},
		{"ruby-w-args-2", "#!/usr/bin/env -vS ruby -w -Ilib:test", Ruby},
	}

	detector := Detector{}

	for _, tc := range testcases {
		t.Run(
			tc.name, func(t *testing.T) {
				res := detector.Detect(tc.snippet)
				got := res[0].Language
				if got != tc.lang {
					t.Errorf("detection error: got %s, want %s", got, tc.lang)
				}
			},
		)
	}
}
