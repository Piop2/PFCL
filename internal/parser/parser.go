package parser

import (
	"bufio"
	"io"

	"github.com/piop2/pfcl/internal/parser/shared"
	"github.com/piop2/pfcl/internal/parser/state"
)

func Parse(reader *bufio.Reader) (data map[string]any, err error) {
	ctx := shared.NewContext()
	currentState := state.NewState(ctx)

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

		next, isProcessed, stateErr := currentState.Process(r)
		if stateErr != nil {
			stateErr.SetPos(line, col)
			return nil, stateErr
		}

		// unread the last rune when it was not processed
		if !isProcessed {
			_ = reader.UnreadRune()
		}
		currentState = next
	}

	// unexpected EOF during parsing
	if currentState.IsParsing() {
		return nil, &shared.ErrSyntax{Message: "unexpected EOF"}
	}

	return ctx.Result, nil
}
