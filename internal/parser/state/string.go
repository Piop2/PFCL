package state

import (
	"fmt"

	"github.com/piop2/pfcl/internal/errors"
	"github.com/piop2/pfcl/internal/parser/shared"
)

type StringState struct {
	ctx        *shared.Context
	onComplete shared.OnCompleteCallback

	// string value
	result string
}

func (s *StringState) SetContext(ctx *shared.Context) {
	s.ctx = ctx
	return
}

func (s *StringState) SetOnComplete(f shared.OnCompleteCallback) {
	s.onComplete = f
	return
}

func (s *StringState) Process(token rune) (next shared.State, isProcessed bool, err errors.ErrPFCL) {
	if shared.IsNewline(token) {
		err = &errors.ErrSyntax{
			Message: fmt.Sprintf("unexpected string token: %s", string(token)),
		}
		return nil, true, err
	}

	// commit
	if token == '"' {
		_ = s.Commit()

		next, _ = s.ctx.StateStack.Pop()
		return next, true, nil
	}

	s.result += string(token)
	return s, true, nil
}

func (s *StringState) Commit() errors.ErrPFCL {
	s.onComplete(s.result)
	return nil
}

func (s *StringState) Flush() (next shared.State, err errors.ErrPFCL) {
	next, _, err = s.Process(0) // give empty rune
	return
}

func (s *StringState) IsParsing() bool {
	return true
}
