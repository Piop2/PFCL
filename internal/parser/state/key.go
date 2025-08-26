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
	if shared.IsWhitespace(token) {
		return s, true, nil
	}

	// allow only letter, digit, and '='
	if !shared.IsAsciiLetter(token) && !shared.IsAsciiDigit(token) && token != '=' {
		err = &shared.ErrSyntax{
			Message: fmt.Sprintf("unexpected result token: %s", string(token)),
		}
		return nil, true, err
	}

	// commit result
	if token == '=' {
		if s.result == "" {
			err = &shared.ErrSyntax{
				Message: "wat is this",
			}
			return nil, true, err
		}

		table, _ := s.ctx.Table()
		if table[s.result] != nil {
			err = &shared.ErrSyntax{
				Message: fmt.Sprintf("duplicate key \"%s\"", s.result),
			}
			return nil, true, err
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
