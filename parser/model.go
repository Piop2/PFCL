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

func (c *Context) Value() (map[string]any, error) {
	return c.ValueAtCursor(c.Cursor)
}

func (c *Context) ValueAtCursor(cursor []string) (map[string]any, error) {
	value := c.Result
	for _, key := range cursor {
		if nested, ok := value[key].(map[string]any); ok {
			value = nested
		} else {
			return nil, errors.New("key error")
		}
	}
	return value, nil
}

type onCompleteCallback func(result any)

type State interface {
	SetContext(ctx *Context)
	SetOnComplete(f onCompleteCallback)

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
