package pfcl

import (
	"io"

	"github.com/piop2/pfcl/internal/formatter"
)

type Encoder struct {
	writer  io.Writer
	indent  string
	newline bool
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{writer: w, indent: "    ", newline: true}
}

func (e *Encoder) SetIndent(indent string) *Encoder {
	e.indent = indent
	return e
}

func (e *Encoder) SetNewline(enable bool) *Encoder {
	e.newline = enable
	return e
}

func (e *Encoder) Encode(v map[string]any) error {
	return formatter.Format(v, e.writer, e.indent, e.newline)
}
