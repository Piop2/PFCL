package parser

import "fmt"

type ReadyState struct {
	ctx *Context
}

func (s *ReadyState) SetOnComplete(_ onCompleteCallback) {}

func (s *ReadyState) SetContext(ctx *Context) {
	s.ctx = ctx
	return
}

func (s *ReadyState) Process(token rune) (next State, isProcessed bool, err ErrPFCL) {
	// Ignore spaces and newline characters
	if token == ' ' || token == '\n' || token == '\r' {
		return s, true, nil
	}

	if token == '[' {
		next = &TableState{}
		isProcessed = true

	} else if token == '#' {
		next = &CommentState{}
		isProcessed = true

	} else if isAsciiLetter(token) {
		//next =
		isProcessed = false

	} else {
		// ERROR!
		err = &ErrSyntax{
			Message: fmt.Sprintf("unexpected character: \"%s\"", string(token)),
		}
		return
	}

	next.SetContext(s.ctx)
	return
}

func (s *ReadyState) IsParsing() bool {
	return false
}
