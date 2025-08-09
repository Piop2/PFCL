package pfcl

import (
	"bufio"
	"os"

	"github.com/piop2/pfcl/parser"
)

type Decoder struct {
	reader *bufio.Reader
}

func (decoder *Decoder) Decode() (map[string]any, error) {
	return parser.Parse(decoder.reader)
}

// NewDecoder returns a new Decoder that reads from r
func NewDecoder(r *bufio.Reader) *Decoder {
	return &Decoder{reader: r}
}

func NewDecoderFromFile(f *os.File) *Decoder {
	return NewDecoder(bufio.NewReader(f))
}
