package parser

type TableState struct {
	ctx       *Context
	name      string
	nameStack Stack[string]
}

func (s *TableState) SetContext(ctx *Context) {
	s.ctx = ctx
	return
}

func (s *TableState) SetOnComplete(_ onCompleteCallback) {}

func (s *TableState) Process(token rune) (next State, isProcessed bool, err ErrPFCL) {
	// Ignore spaces and newline characters
	if token == ' ' || token == '\n' {
		return s, true, nil
	}

	// the end of table declaration
	if token == ']' {
		// table declaration closed without a name
		if s.name == "" {
			return nil, true, &ErrSyntax{Message: "missing table name"}
		}

		if s.nameStack.data == nil {
			s.nameStack.data = []string{}
		}

		table, err := s.ctx.TableAtCursor(s.nameStack.data)
		if err != nil {
			return nil, true, ToErrPFCL(err)
		}

		// make new table
		table[s.name] = map[string]any{}

		// set cursor
		s.ctx.Cursor = append(s.nameStack.data, s.name)

		return &ReadyState{ctx: s.ctx}, true, nil

	} else if token == '.' {
		if s.name == "" {
			return nil, true, &ErrSyntax{Message: "invalid table key: consecutive dots"}
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
