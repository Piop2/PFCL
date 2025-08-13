package parser

type ReadyState struct {
	ctx *Context
}

func (s *ReadyState) SetOnComplete(_ onCompleteCallback) {}

func (s *ReadyState) SetContext(ctx *Context) {
	s.ctx = ctx
	return
}

func (s *ReadyState) Process(token string) (next State, isProcessed bool, err error) {
	// Ignore spaces and newline characters
	if token == " " || token == "\n" || token == "\r" {
		return s, true, nil
	}

	//next = s
	if token == "[" {
		next = &TableState{}
		isProcessed = true

	} else if token == "#" {
		next = &CommentState{}
		isProcessed = true
	}

	next.SetContext(s.ctx)
	return
}

func (s *ReadyState) IsParsing() bool {
	return false
}
