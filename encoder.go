package pfcl

import (
	"bytes"
	"io"

	"github.com/piop2/pfcl/internal/formatter"
)

type Encoder struct {
	writer io.Writer
	indent string
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{writer: w, indent: "    "}
}

func (e *Encoder) SetIndent(indent string) *Encoder {
	e.indent = indent
	return e
}

func (e *Encoder) Encode(v map[string]any) error {
	return formatter.Format(v, e.writer, e.indent)
}

// Marshal encodes a map[string]any into bytes using Encoder
func Marshal(v map[string]any) ([]byte, error) {
	var buf bytes.Buffer
	var writer io.Writer = &buf

	err := NewEncoder(writer).Encode(v)
	return buf.Bytes(), err
}
