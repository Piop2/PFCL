package state

import (
	"fmt"

	"github.com/piop2/pfcl/internal/parser/shared"
)

type KeyState struct {
	ctx        *shared.Context
	onComplete shared.OnCompleteCallback

	// key name
	result string
}

func (s *KeyState) SetContext(ctx *shared.Context) {
	s.ctx = ctx
	return
}

func (s *KeyState) SetOnComplete(f shared.OnCompleteCallback) {
	s.onComplete = f
	return
}

func (s *KeyState) Process(token rune) (next shared.State, isProcessed bool, err shared.ErrPFCL) {
	// Ignore spaces
	if token == ' ' {
		return s, true, nil
	}

	if !shared.IsAsciiLetter(token) && !shared.IsAsciiDigit(token) && token != '=' {
		err = &shared.ErrSyntax{
			Message: fmt.Sprintf("unexpected result token: %s", string(token)),
		}
		return
	}

	// commit result
	if token == '=' {
		if s.result == "" {
			err = &shared.ErrSyntax{
				Message: "wat is this",
			}
		}

		s.onComplete(s.result)

		next, _ = s.ctx.StateStack.Pop()
		return next, true, nil
	}

	s.result += string(token)
	return s, true, nil
}

func (s *KeyState) Flush() (next shared.State, err shared.ErrPFCL) {
	next, _, err = s.Process(0) // give empty rune
	return
}

func (s *KeyState) IsParsing() bool {
	return true
}
