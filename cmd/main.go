package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"net/url"
	"os"
	"strings"
)

const (
	todoMatchComment = "// TODO"
	todoValidComment = "// TODO:"
	todoJumpCut      = len(todoValidComment)
	ticketComment    = "// @ticket"
	ticketJumpCut    = len(ticketComment)
)

var (
	ErrEmptyTodoDescription         = errors.New("todos must have a small description")
	ErrUntrackedTodo                = errors.New("todos must have task-manager link")
	ErrInvalidLink                  = errors.New("ticket property invalid url link")
	ErrInvalidIssueTrackerLink      = errors.New("invalid task-manager url link")
	ErrMissingTicketPropertyComment = errors.New("TODO's must have \"@ticket\" comment-property")
	ErrMissingTicketLink            = errors.New("TODO's @ticket property must not be empty")
	ErrInvalidTODOComment           = errors.New("TODO's must have a ':' before its description")
)

// whats up
type visitor struct {
	fset *token.FileSet
}

// TODO: Must have the possibility of recieveing valid url for ticket (only allow for a specific domain)
func parseTodo(commentGroup *ast.CommentGroup) error {
	todoDesc := commentGroup.List[0]
	if len(todoDesc.Text) < todoJumpCut || todoDesc.Text[todoJumpCut-1] != ':' {
		return ErrInvalidTODOComment
	}

	text := todoDesc.Text[todoJumpCut:]
	if text == "" {
		return ErrEmptyTodoDescription
	}

	// Validate 2nd property: @ticket
	if len(commentGroup.List) < 2 {
		return ErrMissingTicketPropertyComment
	}
	todoTicket := commentGroup.List[1]

	if !strings.HasPrefix(todoTicket.Text, ticketComment) {
		// make this error a struct so that the function above can use errors.As and provide a better error message
		return ErrUntrackedTodo
	}
	rawURL := strings.Trim(todoTicket.Text[ticketJumpCut:], " ")
	if rawURL == "" {
		return ErrMissingTicketLink
	}
	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return ErrInvalidLink
	}

	return nil
}

func (v visitor) Visit(node ast.Node) ast.Visitor {
	commentGroup, ok := node.(*ast.CommentGroup)
	if !ok {
		return v
	}

	if !strings.HasPrefix(commentGroup.List[0].Text, todoMatchComment) {
		return v
	}

	err := parseTodo(commentGroup)
	if err != nil {
		fmt.Printf("%s: %s\n",
			v.fset.Position(node.Pos()),
			err.Error(),
		)
		return v
	}
	return v
}

func main() {
	v := visitor{fset: token.NewFileSet()}
	for _, filePath := range os.Args[1:] {
		if filePath == "--" {
			continue // Allows for: go run main.go -- input.go
		}

		f, err := parser.ParseFile(v.fset, filePath, nil, 0|parser.ParseComments)
		if err != nil {
			log.Fatalf("Failed to parse file %s: %s", filePath, err)
		}

		ast.Walk(&v, f)
	}
}
