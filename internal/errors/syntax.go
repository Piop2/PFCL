package errors

import "fmt"

// ErrSyntax represents a syntax error in parsing.
type ErrSyntax struct {
	Pos     [2]int
	Message string
}

func (e *ErrSyntax) SetMessage(msg string) { e.Message = msg }
func (e *ErrSyntax) SetPos(line, col int)  { e.Pos = [2]int{line, col} }
func (e *ErrSyntax) Error() string {
	return fmt.Sprintf("[%d:%d] syntax error: %s", e.Pos[0], e.Pos[1], e.Message)
}
