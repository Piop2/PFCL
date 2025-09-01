package state

import (
	"github.com/piop2/pfcl/internal/errors"
	"github.com/piop2/pfcl/internal/parser/shared"
)

type CommentState struct {
	ctx *shared.Context
}

func (s *CommentState) SetContext(ctx *shared.Context) {
	s.ctx = ctx
	return
}

func (s *CommentState) SetOnComplete(_ shared.OnCompleteCallback) {}

func (s *CommentState) Process(token rune) (next shared.State, isProcessed bool, err errors.ErrPFCL) {
	if !shared.IsNewline(token) {
		return s, true, nil
	}

	next, _ = s.ctx.StateStack.Pop()
	return next, true, nil
}

func (s *CommentState) Flush() (next shared.State, err errors.ErrPFCL) {
	next, _, err = s.Process(0) // give empty rune
	return
}

func (s *CommentState) Commit() errors.ErrPFCL {
	panic("very big freakin' panic")
}

func (s *CommentState) IsParsing() bool {
	return true
}
