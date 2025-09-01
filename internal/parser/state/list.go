package state

import (
	"github.com/piop2/pfcl/internal/errors"
	"github.com/piop2/pfcl/internal/parser/shared"
)

type ListState struct {
	ctx        *shared.Context
	onComplete shared.OnCompleteCallback

	result []any
	buffer any
}

func (s *ListState) SetContext(ctx *shared.Context) {
	s.ctx = ctx
	return
}

func (s *ListState) SetOnComplete(f shared.OnCompleteCallback) {
	s.onComplete = f
	return
}

func (s *ListState) appendBuffer() errors.ErrPFCL {
	if s.buffer == nil {
		err := &errors.ErrSyntax{
			Message: "duplicated ',' in list",
		}
		return err
	}

	s.result = append(s.result, s.buffer)
	s.buffer = nil
	return nil
}

func (s *ListState) Process(token rune) (next shared.State, isProcessed bool, err errors.ErrPFCL) {
	// append result
	if token == ',' {
		err = s.appendBuffer()
		if err != nil {
			return nil, true, err
		}

		return s, true, nil
	}

	// ignore spaces
	if shared.IsSpace(token) || shared.IsNewline(token) {
		return s, true, nil
	}

	// commit result
	if token == '}' {

		// if buffer is nil     -> empty list
		// if buffer is not nil ->  append the last buffered element
		if s.buffer != nil {
			_ = s.appendBuffer()
		}

		_ = s.Commit()

		next, _ = s.ctx.StateStack.Pop()
		return next, true, nil
	}

	s.ctx.StateStack.Push(s)

	next = &ValueState{}
	next.SetContext(s.ctx)
	next.SetOnComplete(func(result any) {
		s.buffer = result
	})

	return next, false, nil
}

func (s *ListState) Commit() errors.ErrPFCL {
	s.onComplete(s.result)
	return nil
}

func (s *ListState) Flush() (next shared.State, err errors.ErrPFCL) {
	next, _, err = s.Process(0) // give empty rune
	return
}

func (s *ListState) IsParsing() bool {
	return true
}
