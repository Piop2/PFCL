package parser

type Context struct {
	Result     map[string]any
	stateStack []State
	keyStack   []string
}

func NewContext() *Context {
	return &Context{
		Result:     map[string]any{},
		stateStack: []State{},
		keyStack:   []string{},
	}
}
