package state

import (
	"github.com/piop2/pfcl/internal/errors"
	"github.com/piop2/pfcl/internal/parser/shared"
)

type ItemState struct {
	ctx *shared.Context

	key   string
	value any
}

func (s *ItemState) SetContext(ctx *shared.Context) {
	s.ctx = ctx
	return
}

func (s *ItemState) SetOnComplete(_ shared.OnCompleteCallback) {}

func (s *ItemState) Process(_ rune) (next shared.State, isProcessed bool, err errors.ErrPFCL) {
	// commit
	if s.key != "" && s.value != nil {
		err = s.Commit()
		if err != nil {
			return nil, true, err
		}

		next, _ = s.ctx.StateStack.Pop()
		return next, false, nil
	}

	s.ctx.StateStack.Push(s)

	if s.key == "" { // key
		next = &KeyState{}
		next.SetOnComplete(func(result any) {
			key, ok := result.(string)
			if !ok {
				panic("BOINK")
			}

			s.key = key
			return
		})
	} else if s.value == nil { // value
		next = &ValueState{}
		next.SetOnComplete(func(result any) {
			s.value = result
			return
		})
	}

	next.SetContext(s.ctx)
	return next, false, nil
}

func (s *ItemState) Commit() errors.ErrPFCL {
	table, err := s.ctx.Table()
	if err != nil {
		return errors.ToErrPFCL(err)
	}

	table[s.key] = s.value
	return nil
}

func (s *ItemState) Flush() (next shared.State, err errors.ErrPFCL) {
	next, _, err = s.Process(0) // give empty rune
	return
}

func (s *ItemState) IsParsing() bool {
	return false
}
