package state

import (
	common2 "github.com/piop2/pfcl/internal/parser/shared"
)

type TableState struct {
	ctx       *common2.Context
	name      string
	nameStack common2.Stack[string]
}

func (s *TableState) SetContext(ctx *common2.Context) {
	s.ctx = ctx
	return
}

func (s *TableState) SetOnComplete(_ common2.OnCompleteCallback) {}

func (s *TableState) Process(token rune) (next common2.State, isProcessed bool, err common2.ErrPFCL) {
	// Ignore spaces and newline characters
	if token == ' ' || token == '\n' {
		return s, true, nil
	}

	// the end of table declaration
	if token == ']' {
		// table declaration closed without a name
		if s.name == "" {
			return nil, true, &common2.ErrSyntax{Message: "missing table name"}
		}

		if s.nameStack.Data == nil {
			s.nameStack.Data = []string{}
		}

		table, err := s.ctx.TableAtCursor(s.nameStack.Data)
		if err != nil {
			return nil, true, common2.ToErrPFCL(err)
		}

		// make new table
		table[s.name] = map[string]any{}

		// set cursor
		s.ctx.Cursor = append(s.nameStack.Data, s.name)

		next, _ = s.ctx.StateStack.Pop()
		return next, true, nil

	} else if token == '.' {
		if s.name == "" {
			return nil, true, &common2.ErrSyntax{Message: "invalid table key: consecutive dots"}
		}

		s.nameStack.Push(s.name)
		s.name = ""
		return s, true, nil
	}

	s.name += string(token)
	return s, true, nil
}

func (s *TableState) IsParsing() bool {
	return true
}
