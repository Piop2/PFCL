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
	if s.key != "" { // && s.value != nil
		table, err := s.ctx.Table()
		if err != nil {
			return nil, true, shared.ToErrPFCL(err)
		}

		//table[s.result] = s.value
		table[s.key] = "BOINK"

		next, _ = s.ctx.StateStack.Pop()
		return next, false, nil
	}

	s.ctx.StateStack.Push(s)

	// result
	if s.key == "" {
		next = &KeyState{}
		next.SetOnComplete(func(result any) {
			key, ok := result.(string)
			if !ok {
				panic("BOINK")
			}

			s.key = key
			return
		})
	}

	// value
	//if s.value == nil {
	//
	//}

	next.SetContext(s.ctx)
	return next, false, nil
}

func (s *ItemState) IsParsing() bool {
	return true
}
