package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

type Runner struct {
	Analyzer *analysis.Analyzer
	cPass    *analysis.Pass
}

func (r *Runner) Run(p *analysis.Pass) (any, error) {
	for _, f := range p.Files {
		r.cPass = p
		ast.Inspect(f, r.inspect)
	}
	return nil, nil
}

func (r *Runner) inspect(node ast.Node) bool {
	commentGroup, ok := node.(*ast.CommentGroup)
	if !ok {
		return true
	}

	if !strings.HasPrefix(commentGroup.List[0].Text, todoMatchComment) {
		return true
	}

	err := parseTodo(commentGroup)
	if err != nil {
		r.cPass.Reportf(node.Pos(), "%s",
			err.Error(),
		)
		return false
	}
	return true
}
