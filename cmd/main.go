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

var (
	ErrEmptyTodo     = errors.New("todos must have a description and link")
	ErrUntrackedTodo = errors.New("todos must have task-manager link")
	ErrInvalidLink   = errors.New("invalid task-manager URL link")
)

// whats up
type visitor struct {
	fset *token.FileSet
}

const jumpCut = len("// TODO")

// TODO: Must have the possibility of recieveing valid url for ticket
func parseTodo(comment *ast.Comment) error {
	str := comment.Text[jumpCut:]
	text := strings.Trim(str, ": ")
	if text == "" {
		return ErrEmptyTodo
	}

	split := strings.Split(text, "@")
	if len(split) < 2 {
		return ErrUntrackedTodo
	}
	// fmt.Println(split)

	_, err := url.ParseRequestURI(split[1])
	if err != nil {
		return ErrInvalidLink
	}

	return nil
}

func (v visitor) Visit(node ast.Node) ast.Visitor {
	comment, ok := node.(*ast.Comment)
	if !ok {
		return v
	}

	if len(comment.Text) < 7 { // Cannot be a TODO comment
		return v
	}
	text := strings.Trim(comment.Text, " ")[3:]
	if text[:4] == "TODO" || text[:5] == "TODO:" {
		err := parseTodo(comment)
		if err != nil {
			fmt.Printf("%s: %s",
				v.fset.Position(node.Pos()),
				err.Error(),
			)
			return v
		}
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
