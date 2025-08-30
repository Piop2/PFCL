package state

import (
	"github.com/piop2/pfcl/internal/parser/model"
	"github.com/piop2/pfcl/internal/parser/shared"
)

type TableState struct {
	ctx       *shared.Context
	name      string
	nameStack model.Stack[string]
}

func (s *TableState) SetContext(ctx *shared.Context) {
	s.ctx = ctx
	return
}

func (s *TableState) SetOnComplete(_ shared.OnCompleteCallback) {}

func (s *TableState) Process(token rune) (next shared.State, isProcessed bool, err shared.ErrPFCL) {
	// Ignore spaces
	if shared.IsSpace(token) {
		return s, true, nil
	}

	// commit
	// the end of table declaration
	if token == ']' {
		// table declaration closed without a name
		if s.name == "" {
			return nil, true, &shared.ErrSyntax{Message: "missing table name"}
		}

		err = s.Commit()
		if err != nil {
			return nil, true, err
		}

		next, _ = s.ctx.StateStack.Pop()
		return next, true, nil

	} else if token == '.' {
		if s.name == "" {
			return nil, true, &shared.ErrSyntax{Message: "invalid table result: consecutive dots"}
		}

		s.nameStack.Push(s.name)
		s.name = ""
		return s, true, nil
	}

	s.name += string(token)
	return s, true, nil
}
func (s *TableState) Commit() shared.ErrPFCL {
	table, err := s.ctx.TableAtCursor(s.nameStack.Data)
	if err != nil {
		return shared.ToErrPFCL(err)
	}

	// make new table
	table[s.name] = map[string]any{}

	// set cursor
	s.ctx.Cursor = append(s.nameStack.Data, s.name)
	return nil
}

func (s *TableState) Flush() (next shared.State, err shared.ErrPFCL) {
	next, _, err = s.Process(0) // give empty rune
	return
}

func (s *TableState) IsParsing() bool {
	return true
}
