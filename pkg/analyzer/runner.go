package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type Runner struct {
	Analyzer *analysis.Analyzer
	pass     *analysis.Pass
}

func (r *Runner) Run(p *analysis.Pass) (any, error) {
	inspector := p.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CommentGroup)(nil),
	}

	r.pass = p // To avoid using a closure.
	inspector.Preorder(nodeFilter, r.inspect)

	return nil, nil
}

func (r *Runner) inspect(node ast.Node) {
	commentGroup := node.(*ast.CommentGroup)

	if !strings.HasPrefix(commentGroup.List[0].Text, todoMatchComment) {
		return
	}

	err := parseTodo(commentGroup)
	if err != nil {
		r.pass.Reportf(node.Pos(), "%s",
			err.Error(),
		)
		return
	}
}
