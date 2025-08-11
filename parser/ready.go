package parser

type ReadyState struct {
	ctx *Context
}

func (s *ReadyState) SetOnComplete(_ onCompleteCallback) {}

func (s *ReadyState) SetContext(ctx *Context) {
	s.ctx = ctx
	return
}

func (s *ReadyState) Update(token string) (State, error) {
	if token == " " || token == "\n" {
		return s, nil
	}

	var state State = s
	if token == "[" {
		state = &TableState{}
	}

	if token == "#" {
		state = &CommentState{}
	}

	state.SetContext(s.ctx)
	return state, nil
}

func (s *ReadyState) IsParsing() bool {
	return false
}
