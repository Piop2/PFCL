package parser

import "fmt"

type SyntaxError struct {
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("Syntax error")
}
