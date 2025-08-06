package pfcl

import (
	"io"

	"github.com/piop2/pfcl/parser"
)

type Decoder struct {
	reader io.Reader
}

func (receiver *Decoder) Decode() (map[string]any, error) {
	_ = parser.NewContext()

	//return map[string]any{}, errors.New("not implemented")
	panic("implement me")
}

// NewDecoder returns a new Decoder that reads from reader
func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{reader: reader}
}

//func Unmarshal(data []byte) (map[string]any, error) {
//	return make(map[string]any), errors.New("not implemented")
//}
