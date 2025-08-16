package pfcl

import (
	"bufio"
	"io"
	"os"

	"github.com/piop2/pfcl/internal/parser"
)

type Decoder struct {
	reader *bufio.Reader
}

func (decoder *Decoder) Decode() (map[string]any, error) {
	return parser.Parse(decoder.reader)
}

// NewDecoder returns a new Decoder that reads from r
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{reader: bufio.NewReader(r)}
}

func NewDecoderFromFile(f *os.File) *Decoder {
	return NewDecoder(f)
}
