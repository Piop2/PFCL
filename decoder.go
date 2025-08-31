package pfcl

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"

	"github.com/piop2/pfcl/internal/parser"
)

type Decoder struct {
	reader *bufio.Reader
}

func (d *Decoder) Decode(v any) error {
	m, err := parser.Parse(d.reader)
	if err != nil {
		return err
	}

	switch target := (v).(type) {
	case *map[string]any:
		*target = m

	default:
		return errors.New("v must be pointer to struct or *map[string]any")
	}

	return nil
}

// NewDecoder returns a new Decoder that reads from r
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{reader: bufio.NewReader(r)}
}

func NewDecoderFromFile(f *os.File) *Decoder {
	return NewDecoder(f)
}

func Unmarshal(data []byte, v *any) error {
	return NewDecoder(bytes.NewReader(data)).Decode(v)
}
