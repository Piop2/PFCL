package shared

import "errors"

type Context struct {
	Result     map[string]any
	StateStack Stack[State]
	Cursor     []string
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

func NewContext() *Context {
	return &Context{
		Result:     map[string]any{},
		StateStack: Stack[State]{},
		Cursor:     []string{},
	}
}
