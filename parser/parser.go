package parser

import (
	"bufio"
	"errors"
	"io"
)

func Parse(reader *bufio.Reader) (map[string]any, error) {
	ctx := NewContext()
	state := NewState(ctx)

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		state, err = state.Update(string(r))
		if err != nil {
			return nil, err
		}
	}

	if state.IsParsing() {
		return nil, errors.New("syntax error")
	}

	return ctx.Result, nil
}
