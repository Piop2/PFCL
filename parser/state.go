package parser

type State interface {
	SetContext(ctx *context)

	Update(
		token string,
	) State

	IsParsing() bool
}

func NewState(ctx *context) State {
	s := ReadyState{}
	s.SetContext(ctx)
	return &s
}

type ReadyState struct {
	ctx *context
}

func (s *ReadyState) SetContext(ctx *context) {
	s.ctx = ctx
	return
}

func (s *ReadyState) Update(token string) State {
	if token != " " {
		newState := &RunningState{}
		return newState
	}

	return s
}

func (s *ReadyState) IsParsing() bool {
	return false
}

type RunningState struct {
}

func (r *RunningState) SetContext(_ *context) {
	return
}

func (r *RunningState) Update(_ string) State {
	return r
}

func (r *RunningState) IsParsing() bool {
	return true
}
