package parser

import (
	"errors"
)

type State interface {
	SetContext(ctx *Context)
	SetOnComplete(f func(result any))

	Update(
		token string,
	) (State, error)

	IsParsing() bool
}

func NewState(ctx *Context) State {
	s := ReadyState{}
	s.SetContext(ctx)
	return &s
}

type ReadyState struct {
	ctx *Context
}

func (s *ReadyState) SetOnComplete(_ func(result any)) {}

func (s *ReadyState) SetContext(ctx *Context) {
	s.ctx = ctx
	return
}

func (s *ReadyState) Update(token string) (State, error) {
	if token == " " || token == "\n" {
		return s, nil
	}

	var state State = s
	if token == "[" {
		state = &TableState{ctx: s.ctx}
	}

	return state, nil
}

func (s *ReadyState) IsParsing() bool {
	return false
}

type TableState struct {
	ctx      *Context
	key      string
	keyQueue []string
}

func (s *TableState) SetContext(ctx *Context) {
	s.ctx = ctx
	return
}

func (s *TableState) SetOnComplete(_ func(result any)) {}

func (s *TableState) Update(token string) (State, error) {
	if token == " " || token == "\n" {
		return nil, ErrSyntax
	}

	if token == "]" {
		if s.key == "" {
			return nil, ErrSyntax
		}

		tableAny, err := s.ctx.ValueAtCursor(s.keyQueue)
		if err != nil {
			return nil, err
		}

		table, ok := tableAny.(map[string]any)
		if !ok {
			return nil, errors.New("unexpected type")
		}

		if table[s.key] != nil {
			return nil, errors.New("duplicate error")
		}

		table[s.key] = map[string]any{}

		s.ctx.cursor = append(s.keyQueue, s.key)

		return &ReadyState{ctx: s.ctx}, nil
	}

	if token == "." {
		if s.key == "" {
			return nil, ErrSyntax
		}

		s.keyQueue = append(s.keyQueue, s.key)
		s.key = ""
		return s, nil
	}

	s.key += token
	return s, nil
}

func (s *TableState) IsParsing() bool {
	return true
}
