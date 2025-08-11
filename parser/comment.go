package parser

type CommentState struct {
	ctx *Context
}

func (s *CommentState) SetContext(ctx *Context) {
	s.ctx = ctx
	return
}

func (s *CommentState) SetOnComplete(_ func(result any)) {}

func (s *CommentState) Update(token string) (State, error) {
	if token != "\n" {
		return s, nil
	}
	return &ReadyState{ctx: s.ctx}, nil
}

func (s *CommentState) IsParsing() bool {
	return true
}
