package parser

type Context struct {
	Result    map[string]any
	viewStack []State
	keyStack  []string
}

func NewContext() *Context {
	return &Context{Result: map[string]any{}}
}
