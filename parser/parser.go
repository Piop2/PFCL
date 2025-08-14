package parser

import (
	"bufio"
	"io"
)

func Parse(reader *bufio.Reader) (data map[string]any, err error) {
	ctx := NewContext()
	state := NewState(ctx)

	line, col := 1, 0
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if r == '\n' {
			line++
			col = 0
		} else {
			col++
		}

		next, isProcessed, stateErr := state.Process(r)
		if stateErr != nil {
			stateErr.SetPos(line, col)
			return nil, stateErr
		}

		// unread the last rune when it was not processed
		if !isProcessed {
			_ = reader.UnreadRune()
		}
		state = next
	}

	// unexpected EOF during parsing
	if state.IsParsing() {
		return nil, &ErrSyntax{Message: "unexpected EOF"}
	}

	return ctx.Result, nil
}
