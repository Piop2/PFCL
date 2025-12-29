package state

import (
	"fmt"

	"github.com/piop2/pfcl/internal/errors"
	"github.com/piop2/pfcl/internal/parser/shared"
)

type NumberType int

type SignType int

const (
	SignPositive SignType = iota
	SignNegative
)

type NumberState struct {
	ctx        *shared.Context
	onComplete shared.OnCompleteCallback

	// result type: int64 | float64
	result any
}

func (s *NumberState) SetContext(ctx *shared.Context) {
	s.ctx = ctx
	return
}

func (s *NumberState) SetOnComplete(f shared.OnCompleteCallback) {
	s.onComplete = f
	return
}

func (s *NumberState) Process(token rune) (next shared.State, isProcessed bool, err errors.ErrPFCL) {
	// commit
	if s.result != nil {
		_ = s.Commit()

		next, _ = s.ctx.StateStack.Pop()
		return next, false, nil
	}

	// ignore spaces
	if shared.IsSpace(token) {
		return s, true, nil
	}

	// allowed characters: digits(0-9), and minus(-)
	if !shared.IsAsciiDigit(token) && token != '-' {
		err = &errors.ErrSyntax{
			Message: fmt.Sprintf("invalid numeric character: '%c'", token),
		}

		return nil, false, err
	}

	var signType SignType
	if token == '-' {
		signType = SignNegative
		isProcessed = true
	} else {
		signType = SignPositive
		isProcessed = false
	}

	s.ctx.StateStack.Push(s)

	next = &IntState{signType: signType}
	next.SetContext(s.ctx)
	next.SetOnComplete(func(result any) {
		s.result = result
	})

	return next, isProcessed, nil
}

func (s *NumberState) Commit() errors.ErrPFCL {
	s.onComplete(s.result)
	return nil
}

func (s *NumberState) Flush() (next shared.State, err errors.ErrPFCL) {
	next, _, err = s.Process(0) // give empty rune
	return
}

func (s *NumberState) IsParsing() bool {
	return false
}
