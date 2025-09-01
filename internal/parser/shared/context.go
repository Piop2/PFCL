package shared

import (
	"errors"

	"github.com/piop2/pfcl/internal/model"
)

// Context holds parsing results, state stack, and the current cursor path.
type Context struct {
	Result     map[string]any     // Parsed Data as nested tables
	StateStack model.Stack[State] // Stack of active states
	Cursor     []string           // Current table path
}

// Table returns the table at the current cursor.
func (c *Context) Table() (table map[string]any, err error) {
	return c.TableAtCursor(c.Cursor)
}

// TableAtCursor returns the table at the specified cursor path.
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

// NewContext creates and initializes a new Context.
func NewContext() *Context {
	return &Context{
		Result:     map[string]any{},
		StateStack: model.Stack[State]{},
		Cursor:     []string{},
	}
}
