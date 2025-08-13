package parser

import (
	"bufio"
	"io"
)

func Parse(reader *bufio.Reader) (data map[string]any, err error) {
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

		next, isProcessed, err := state.Process(string(r))
		if err != nil {
			return nil, err
		}

		// unread the last rune when it was not processed
		if !isProcessed {
			_ = reader.UnreadRune()
		}
		state = next
	}

	// unexpected EOF during parsing
	if state.IsParsing() {
		return nil, ErrSyntax
	}

	return ctx.Result, nil
}
