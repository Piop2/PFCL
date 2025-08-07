package parser

import (
	"bufio"
	"errors"
	"io"
)

type context struct {
	Result map[string]any
}

func newContext() *context {
	return &context{Result: map[string]any{}}
}

func Parse(reader *bufio.Reader) (map[string]any, error) {
	ctx := newContext()
	state := NewState(ctx)

	for {
		// 글자 단위로 읽기
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		state = state.Update(string(r))
	}

	if state.IsParsing() {
		return nil, errors.New("syntax error")
	}

	return ctx.Result, nil
}
