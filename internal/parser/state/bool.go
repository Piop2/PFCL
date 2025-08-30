package state

import (
	"fmt"

	"github.com/piop2/pfcl/internal/parser/model"
	"github.com/piop2/pfcl/internal/parser/shared"
)

type BoolState struct {
	ctx        *shared.Context
	onComplete shared.OnCompleteCallback

	value bool

	isChecking bool
	targets    model.Queue[rune]
	buffer     rune
}

func (s *BoolState) SetContext(ctx *shared.Context) {
	s.ctx = ctx
	return
}

func (s *BoolState) SetOnComplete(f shared.OnCompleteCallback) {
	s.onComplete = f
	return
}

func (s *BoolState) Process(token rune) (next shared.State, isProcessed bool, err shared.ErrPFCL) {
	if !s.isChecking {
		s.isChecking = true

		if token == 't' { // start with 't' is true
			s.value = true

			s.targets.Enqueue('r')
			s.targets.Enqueue('u')
			s.targets.Enqueue('e')

		} else if token == 'f' { // start with 'f' is false
			s.value = false

			s.targets.Enqueue('a')
			s.targets.Enqueue('l')
			s.targets.Enqueue('s')
			s.targets.Enqueue('e')

		} else {
			// ERROR!
			err = &shared.ErrSyntax{
				Message: fmt.Sprintf(""),
			}
			return nil, true, err
		}

		s.buffer, _ = s.targets.Dequeue()
		return s, true, nil
	}

	if token != s.buffer {
		// ERROR!
		err = &shared.ErrSyntax{
			Message: "i thought it was bool",
		}
		return nil, true, err
	}

	var ok bool
	s.buffer, ok = s.targets.Dequeue()

	// commit
	if s.isChecking && !ok {
		s.onComplete(s.value)

		next, _ = s.ctx.StateStack.Pop()
		return next, true, nil
	}

	return s, true, nil
}

func (s *BoolState) Flush() (next shared.State, err shared.ErrPFCL) {
	next, _, err = s.Process(0) // give empty rune
	return
}

func (s *BoolState) IsParsing() bool {
	return true
}
