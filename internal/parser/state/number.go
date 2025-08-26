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

type NumberState struct {
	ctx        *shared.Context
	onComplete shared.OnCompleteCallback

	// number value: int64 | float64
	result     string
	resultType NumberType
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

		var result any
		var convertErr error

		switch s.resultType {
		case NumberInt64:
			// string to int64 convert
			result, convertErr = strconv.ParseInt(s.result, 10, 64)
			if convertErr != nil {
				panic("failed to convert string to int64")
			}

		case NumberFloat64:
			// string to float64 convert
			result, convertErr = strconv.ParseFloat(s.result, 64)
			if convertErr != nil {
				panic("failed to convert string to int64")
			}

		default:
			panic("invalid number type")
		}

		s.onComplete(result)

		next, _ = s.ctx.StateStack.Pop()
		return next, false, nil
	}

	if !shared.IsAsciiDigit(token) && token != '.' {
		err = &shared.ErrSyntax{
			Message: fmt.Sprintf("invalid numeric character: '%c'", token),
		}

		return nil, true, err
	}

	// float type
	if token == '.' {
		if s.result == "" {
			err = &shared.ErrSyntax{
				Message: "unexpected numeric character",
			}
			return nil, true, err
		}

		s.resultType = NumberFloat64
	}

	s.result += string(token)

	return s, true, nil
}

func (s *NumberState) Flush() (next shared.State, err shared.ErrPFCL) {
	next, _, err = s.Process(0) // give empty rune
	return
}

func (s *NumberState) IsParsing() bool {
	return true
}
