package state

import (
	"fmt"
	"strconv"

	"github.com/piop2/pfcl/internal/errors"
	"github.com/piop2/pfcl/internal/parser/shared"
)

type IntState struct {
	ctx        *shared.Context
	onComplete shared.OnCompleteCallback

	signType SignType
	result   string
}

func (s *IntState) SetContext(ctx *shared.Context) {
	s.ctx = ctx
}

func (s *IntState) SetOnComplete(f shared.OnCompleteCallback) {
	s.onComplete = f
}

func (s *IntState) Process(token rune) (next shared.State, isProcessed bool, err errors.ErrPFCL) {
	//commit
	if shared.IsWhitespace(token) || token == ',' || token == 0 {
		err = s.Commit()
		if err != nil {
			return nil, false, err
		}

		next, _ = s.ctx.StateStack.Pop()

		return next, false, nil
	}

	// allowed characters: digits(0-9)
	if !shared.IsAsciiDigit(token) {
		err = &errors.ErrSyntax{
			Message: fmt.Sprintf("invalid numeric character: %q", token),
		}

		return nil, true, err
	}

	s.result += string(token)

	// floating point
	//if token == '.' {
	//	next = &FloatState{signType: s.signType, result: s.result}
	//	next.SetContext(s.ctx)
	//	next.SetOnComplete(s.onComplete)
	//
	//	return next, true, nil
	//}

	return s, true, nil
}

func (s *IntState) Commit() errors.ErrPFCL {
	// string to int64 convert
	result, convertErr := strconv.ParseInt(s.result, 10, 64)
	if s.signType == SignNegative {
		result *= -1
	}

	// TODO
	if convertErr != nil {
		return &errors.BaseErr{
			Message: fmt.Sprintf("failed to convert string to int64: %s", s.result),
		}
	}

	s.onComplete(result)
	return nil
}

func (s *IntState) Flush() (next shared.State, err errors.ErrPFCL) {
	next, _, err = s.Process(0) // give empty rune
	return
}

func (s *IntState) IsParsing() bool {
	return true
}
