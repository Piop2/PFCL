package state

import (
	"fmt"
	"strconv"

	"github.com/piop2/pfcl/internal/errors"
	"github.com/piop2/pfcl/internal/parser/shared"
)

type FloatState struct {
	ctx        *shared.Context
	onComplete shared.OnCompleteCallback

	signType SignType
	result   string
}

func (s *FloatState) SetContext(ctx *shared.Context) {
	s.ctx = ctx
}

func (s *FloatState) SetOnComplete(f shared.OnCompleteCallback) {
	s.onComplete = f
}

func (s *FloatState) SetSignType(signType SignType) {
	s.signType = signType
}

func (s *FloatState) Process(token rune) (next shared.State, isProcessed bool, err errors.ErrPFCL) { //commit
	if shared.IsWhitespace(token) || token == ',' || token == '}' || token == 0 {
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

	return s, true, nil
}

func (s *FloatState) Commit() errors.ErrPFCL {
	// string to float64 convert
	result, convertErr := strconv.ParseFloat(s.result, 64)
	if s.signType == SignNegative {
		result *= -1
	}

	if convertErr != nil {
		return &errors.BaseErr{
			Message: fmt.Sprintf("failed to convert string to float64: %s", s.result),
		}
	}

	s.onComplete(result)
	return nil
}

func (s *FloatState) Flush() (next shared.State, err errors.ErrPFCL) {
	next, _, err = s.Process(0) // give empty rune
	return
}

func (s *FloatState) IsParsing() bool {
	return false
}
