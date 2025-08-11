package parser

import (
	"errors"
)

type Context struct {
	Result     map[string]any
	StateStack []State
	Cursor     []string
}

func NewContext() *Context {
	return &Context{
		Result:     map[string]any{},
		StateStack: []State{},
		Cursor:     []string{},
	}
}

func (c *Context) Value() (any, error) {
	return c.ValueAtCursor(c.Cursor)
}

func (c *Context) ValueAtCursor(cursor []string) (any, error) {
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
