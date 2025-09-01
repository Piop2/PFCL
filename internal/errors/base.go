package errors

import "fmt"

// ToErrPFCL converts a standard error into an ErrPFCL.
// Returns nil if input error is nil.
func ToErrPFCL(err error) ErrPFCL {
	if err == nil {
		return nil
	}
	return &BaseErr{Message: err.Error()}
}

// BaseErr is a basic implementation of ErrPFCL.
type BaseErr struct {
	Pos     [2]int
	Message string
}

func (e *BaseErr) SetMessage(msg string) { e.Message = msg }
func (e *BaseErr) SetPos(line, col int)  { e.Pos = [2]int{line, col} }
func (e *BaseErr) Error() string {
	return fmt.Sprintf("[%d:%d] %s", e.Pos[0], e.Pos[1], e.Message)
}
