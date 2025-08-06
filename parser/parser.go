package parser

type Context struct {
	Data map[string]any
}

func NewContext() *Context {
	return &Context{Data: map[string]any{}}
}

type State interface {
	SetContext(ctx *Context)

	Update(
		token string,
	) State
}

func NewState(ctx *Context) *StartState {
	state := &StartState{}
	state.SetContext(ctx)
	return state
}

type StartState struct {
	ctx *Context
}

func (state *StartState) SetContext(ctx *Context) {
	state.ctx = ctx
}

func (state *StartState) Update(token string) State {
	//TODO implement me
	panic("implement me")
}

type EndState struct{}

func (state *EndState) SetContext(_ *Context) {}

func (state *EndState) Update(_ string) State {
	return state
}
