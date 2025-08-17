package state

import (
	"fmt"

	"github.com/piop2/pfcl/internal/parser/shared"
)

type KeyState struct {
	ctx        *shared.Context
	onComplete shared.OnCompleteCallback

	key string
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
			Message: fmt.Sprintf("unexpected key token: %s", string(token)),
		}
		return
	}

	// commit key
	if token == '=' {
		if s.key == "" {
			err = &shared.ErrSyntax{
				Message: "wat is this",
			}
		}

		s.onComplete(s.key)

		next, _ = s.ctx.StateStack.Pop()
		return next, true, nil
	}

	s.key += string(token)
	return s, true, nil
}

func (s *KeyState) IsParsing() bool {
	return true
}
