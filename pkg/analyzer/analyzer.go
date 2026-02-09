package analyzer

import (
	"errors"
	"go/ast"
	"net/url"
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
