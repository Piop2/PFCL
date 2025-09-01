package shared

import "github.com/piop2/pfcl/internal/errors"

type OnCompleteCallback func(result any)

// State represents a parser state.
type State interface {
	// SetContext assigns the parser context.
	SetContext(ctx *Context)

	// SetOnComplete sets a callback to be called on completion.
	SetOnComplete(f OnCompleteCallback)

	// Process handles a token and returns next state, whether it was processed, and any error.
	Process(token rune) (next State, isProcessed bool, err errors.ErrPFCL)

	// Commit finalizes the state and calls OnComplete with the result.
	Commit() errors.ErrPFCL

	// Flush finalizes the state at EOF
	Flush() (next State, err errors.ErrPFCL)

	// IsParsing returns true if the state is actively parsing.
	IsParsing() bool
}
