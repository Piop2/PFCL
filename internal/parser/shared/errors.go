package shared

import "fmt"

type ErrPFCL interface {
	SetMessage(message string)
	SetPos(line int, col int)
	Error() string
}

func ToErrPFCL(err error) ErrPFCL {
	if err == nil {
		return nil
	}
	return &BaseErr{Message: err.Error()}
}

type BaseErr struct {
	Pos     [2]int
	Message string
}

func (e *BaseErr) SetMessage(msg string) { e.Message = msg }
func (e *BaseErr) SetPos(line, col int)  { e.Pos = [2]int{line, col} }
func (e *BaseErr) Error() string {
	return fmt.Sprintf("[%d:%d] %s", e.Pos[0], e.Pos[1], e.Message)
}

type ErrSyntax struct {
	Pos     [2]int
	Message string
}

func (e *ErrSyntax) SetMessage(msg string) { e.Message = msg }
func (e *ErrSyntax) SetPos(line, col int)  { e.Pos = [2]int{line, col} }
func (e *ErrSyntax) Error() string {
	return fmt.Sprintf("syntax error [%d:%d]: %s", e.Pos[0], e.Pos[1], e.Message)
}
