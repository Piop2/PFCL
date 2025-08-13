package parser

import (
	"errors"
)

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

type Context struct {
	Result     map[string]any
	StateStack Stack[State]
	Cursor     []string
}

func NewContext() *Context {
	return &Context{
		Result:     map[string]any{},
		StateStack: Stack[State]{},
		Cursor:     []string{},
	}
}

func (c *Context) Table() (table map[string]any, err error) {
	return c.TableAtCursor(c.Cursor)
}

func (c *Context) TableAtCursor(cursor []string) (table map[string]any, err error) {
	table = c.Result
	for _, key := range cursor {
		if nested, ok := table[key].(map[string]any); ok {
			table = nested
		} else {
			return nil, errors.New("table name error")
		}
	}
	return
}

type onCompleteCallback func(result any)

type State interface {
	SetContext(ctx *Context)
	SetOnComplete(f onCompleteCallback)

	Process(
		token string,
	) (next State, isProcessed bool, err error)

	IsParsing() bool
}

func NewState(ctx *Context) State {
	s := ReadyState{}
	s.SetContext(ctx)
	return &s
}
