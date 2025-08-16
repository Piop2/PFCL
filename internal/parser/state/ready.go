package state

import (
	"fmt"

	common2 "github.com/piop2/pfcl/internal/parser/shared"
)

type ReadyState struct {
	ctx *common2.Context
}

func (s *ReadyState) SetContext(ctx *common2.Context) {
	s.ctx = ctx
	return
}

func (s *ReadyState) SetOnComplete(_ common2.OnCompleteCallback) {}

func (s *ReadyState) Process(token rune) (next common2.State, isProcessed bool, err common2.ErrPFCL) {
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

	} else if common2.IsAsciiLetter(token) {
		//next = &ItemState{} <--- Implement please
		isProcessed = false

	} else {
		// ERROR!
		err = &common2.ErrSyntax{
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
