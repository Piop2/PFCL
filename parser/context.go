package parser

import "errors"

type Context struct {
	Result     map[string]any
	stateStack []State
	cursor     []string
}

func NewContext() *Context {
	return &Context{
		Result:     map[string]any{},
		stateStack: []State{},
		cursor:     []string{},
	}
}

func (c *Context) Value() (any, error) {
	return c.ValueAtCursor(c.cursor)
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
