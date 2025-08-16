package state

import (
	common2 "github.com/piop2/pfcl/internal/parser/shared"
)

type CommentState struct {
	ctx *common2.Context
}

func (s *CommentState) SetContext(ctx *common2.Context) {
	s.ctx = ctx
	return
}

func (s *CommentState) SetOnComplete(_ common2.OnCompleteCallback) {}

func (s *CommentState) Process(token rune) (next common2.State, isProcessed bool, err common2.ErrPFCL) {
	if token != '\n' {
		return s, true, nil
	}

	next, _ = s.ctx.StateStack.Pop()
	return next, true, nil
}

func (s *CommentState) IsParsing() bool {
	return true
}
