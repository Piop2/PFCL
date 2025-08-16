package state

import (
	"fmt"

	"github.com/piop2/pfcl/internal/parser/shared"
)

type ReadyState struct {
	ctx *shared.Context
}

func (s *ReadyState) SetContext(ctx *shared.Context) {
	s.ctx = ctx
	return
}

func (s *ReadyState) SetOnComplete(_ shared.OnCompleteCallback) {}

func (s *ReadyState) Process(token rune) (next shared.State, isProcessed bool, err shared.ErrPFCL) {
	// Ignore spaces and newline characters
	if token == ' ' || token == '\n' || token == '\r' {
		return s, true, nil
	}

	s.ctx.StateStack.Push(s)

	if token == '[' {
		next = &TableState{}
		isProcessed = true

	} else if token == '#' {
		next = &CommentState{}
		isProcessed = true

	} else if shared.IsAsciiLetter(token) {
		//next = &ItemState{} <--- Implement please
		isProcessed = false

	} else {
		// ERROR!
		err = &shared.ErrSyntax{
			Message: fmt.Sprintf("unexpected character: \"%s\"", string(token)),
		}
		return
	}

	next.SetContext(s.ctx)
	return
}

func (s *ReadyState) IsParsing() bool {
	return false
}
