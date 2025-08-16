package state

import (
	"github.com/piop2/pfcl/internal/parser/shared"
)

func NewState(ctx *shared.Context) shared.State {
	s := ReadyState{}
	s.SetContext(ctx)
	return &s
}
