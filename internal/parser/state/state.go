package state

import (
	common2 "github.com/piop2/pfcl/internal/parser/shared"
)

func NewState(ctx *common2.Context) common2.State {
	s := ReadyState{}
	s.SetContext(ctx)
	return &s
}
