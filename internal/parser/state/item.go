package state

import (
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

func (s *ItemState) Process(_ rune) (next shared.State, isProcessed bool, err shared.ErrPFCL) {
	if s.key != "" && s.value != nil {
		table, err := s.ctx.Table()
		if err != nil {
			return nil, true, shared.ToErrPFCL(err)
		}

		table[s.key] = s.value

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

func (s *ItemState) IsParsing() bool {
	return true
}
