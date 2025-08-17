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
			return nil, false, shared.ToErrPFCL(err)
		}

		//table[s.key] = s.value
		table[s.key] = "BOINK"

		next, _ = s.ctx.StateStack.Pop()
		return next, false, nil
	}

	s.ctx.StateStack.Push(s)

	//Process KeyState ValueState
	if s.key == "" {
		next = &KeyState{}
		isProcessed = false
		next.SetOnComplete(func(result any) {
			key, ok := result.(string)
			if !ok {
				panic("BOINK")
			}

			s.key = key
			return
		})
	}

	next.SetContext(s.ctx)
	return
}

func (s *ItemState) IsParsing() bool {
	return true
}
