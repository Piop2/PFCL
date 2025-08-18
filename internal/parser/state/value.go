package state

import (
	"fmt"

	"github.com/piop2/pfcl/internal/parser/shared"
)

type ValueState struct {
	ctx        *shared.Context
	onComplete shared.OnCompleteCallback

	result any
}

func (s *ValueState) SetContext(ctx *shared.Context) {
	s.ctx = ctx
	return
}

func (s *ValueState) SetOnComplete(f shared.OnCompleteCallback) {
	s.onComplete = f
	return
}

func (s *ValueState) Process(token rune) (next shared.State, isProcessed bool, err shared.ErrPFCL) {
	// commit
	if s.result != nil {
		s.onComplete(s.result)

		next, _ = s.ctx.StateStack.Pop()
		return next, false, nil
	}

	// Ignore spaces
	if token == ' ' {
		return s, true, nil
	}

	// string value
	if token == '"' {
		next = &StringState{}
	} else {
		err = &shared.ErrSyntax{
			Message: fmt.Sprintf("unexpected token: %s", string(token)),
		}
		return nil, true, err
	}

	s.ctx.StateStack.Push(s)

	next.SetContext(s.ctx)
	next.SetOnComplete(func(result any) {
		s.result = result
		return
	})

	return next, true, nil
}

func (s *ValueState) IsParsing() bool {
	return true
}
