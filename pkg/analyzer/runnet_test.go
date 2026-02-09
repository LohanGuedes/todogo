package analyzer

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

func TestAll(t *testing.T) {
	runner := Runner{
		Analyzer: &analysis.Analyzer{
			Name:     "todogo",
			Doc:      "Checks for TODOs comments and make sure they're being tracked in an issue-tracker",
			Requires: []*analysis.Analyzer{inspect.Analyzer},
		},
	}
	runner.Analyzer.Run = runner.Run

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get wd: %s", err.Error())
	}

	testdata := filepath.Join(
		filepath.Dir(filepath.Dir(wd)),
		"testdata",
	)

	analysistest.Run(t, testdata, runner.Analyzer, "td")
}
