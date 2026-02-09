package main

import (
	"github.com/lohanguedes/todogo/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	runner := analyzer.Runner{
		Analyzer: &analysis.Analyzer{
			Name: "todogo",
			Doc:  "Checks for TODOs comments and make sure they're being tracked in an issue-tracker",
		},
	}
	runner.Analyzer.Run = runner.Run

	singlechecker.Main(runner.Analyzer)
}
