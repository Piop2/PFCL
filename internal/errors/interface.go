package errors

// ErrPFCL is the common interface for all PFCL errors.
type ErrPFCL interface {
	SetMessage(message string)
	SetPos(line int, col int)
	Error() string
}
