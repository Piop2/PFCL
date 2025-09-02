package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/piop2/pfcl"
)

var MockData = map[string]any{
	"int":   1,
	"float": -1.2,
	"str":   "BOINK",
	"bool":  true,
	"server": map[string]any{
		"int": 1,
		"beta": map[string]any{
			"int": 1,
		},
	},
}

func main() {

	var buf bytes.Buffer
	var writer io.Writer = &buf

	encoder := pfcl.NewEncoder(writer)

	err := encoder.Encode(MockData)
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
