package state

import (
	"fmt"
	"strconv"

	"github.com/piop2/pfcl/internal/parser/shared"
)

type NumberType int

const (
	NumberInt64 NumberType = iota
	NumberFloat64
)

type SignType int

const (
	SignPositive SignType = iota
	SignNegative
)

type NumberState struct {
	ctx        *shared.Context
	onComplete shared.OnCompleteCallback

	// number value: int64 | float64
	result     string
	numberType NumberType
	signType   SignType
}

func (s *NumberState) SetContext(ctx *shared.Context) {
	s.ctx = ctx
	return
}

func (s *NumberState) SetOnComplete(f shared.OnCompleteCallback) {
	s.onComplete = f
	return
}

func (s *NumberState) Process(token rune) (next shared.State, isProcessed bool, err shared.ErrPFCL) {
	// commit
	if shared.IsWhitespace(token) ||
		token == ',' || token == '}' {
		if s.result == "" {
			err = &shared.ErrSyntax{
				Message: "runtime error",
			}
			return nil, true, err
		}

		_ = s.Commit()

		next, _ = s.ctx.StateStack.Pop()
		return next, false, nil
	}

	// allowed characters: digits(0-9), dot(.), and minus(-)
	if !shared.IsAsciiDigit(token) && token != '.' && token != '-' {
		err = &shared.ErrSyntax{
			Message: fmt.Sprintf("invalid numeric character: '%c'", token),
		}

		return nil, true, err
	}

	// float type
	if token == '.' {
		// THIS SHOULD NEVER HAPPEN
		//
		// invalid syntax
		//if s.result == "" {
		//	err = &shared.ErrSyntax{
		//		Message: "unexpected numeric character",
		//	}
		//	return nil, true, err
		//}

		// consecutive dots are not allowed (currently treated as syntax error).
		// TODO: this should be changed later because of range parsing
		if s.numberType == NumberFloat64 {
			err = &shared.ErrSyntax{
				Message: "unexpected numeric character",
			}
			return nil, true, err
		}

		s.numberType = NumberFloat64
	}

	// negative sign
	if token == '-' {
		if s.signType == SignNegative {
			err = &shared.ErrSyntax{
				Message: "unexpected sign in number",
			}
			return nil, true, err
		}

		if s.result != "" {
			err = &shared.ErrSyntax{
				Message: "unexpected sign in number",
			}
			return nil, true, err
		}

		s.signType = SignNegative

		return s, true, nil
	}

	s.result += string(token)

	return s, true, nil
}

func (s *NumberState) Commit() shared.ErrPFCL {
	// apply negative sign if needed
	if s.signType == SignNegative {
		s.result = "-" + s.result
	}

	var value any
	var convertErr error
	switch s.numberType {
	case NumberInt64:
		// string to int64 convert
		value, convertErr = strconv.ParseInt(s.result, 10, 64)
		if convertErr != nil {
			panic("failed to convert string to int64")
		}

	case NumberFloat64:
		// string to float64 convert
		value, convertErr = strconv.ParseFloat(s.result, 64)
		if convertErr != nil {
			panic("failed to convert string to float64")
		}

	default:
		panic("invalid number type")
	}

	s.onComplete(value)
	return nil
}

func (s *NumberState) Flush() (next shared.State, err shared.ErrPFCL) {
	next, _, err = s.Process(0) // give empty rune
	return
}

func (s *NumberState) IsParsing() bool {
	return true
}
