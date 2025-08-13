package parser

type CommentState struct {
	ctx *Context
}

func (s *CommentState) SetContext(ctx *Context) {
	s.ctx = ctx
	return
}

func (s *CommentState) SetOnComplete(_ onCompleteCallback) {}

func (s *CommentState) Process(token string) (next State, isProcessed bool, err error) {
	if token != "\n" {
		return s, true, nil
	}
	return &ReadyState{ctx: s.ctx}, true, nil
}

func (s *CommentState) IsParsing() bool {
	return true
}
