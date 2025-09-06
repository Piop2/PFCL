package main

import (
	"os"

	"github.com/piop2/pfcl"
)

var MockData = map[string]any{
	"int":   1,
	"float": -1.2,
	"str":   "BOINK",
	"bool":  true,
	"list":  []any{1, 2, []any{3, 4}},
	"server": map[string]any{
		"int": 1,
		"beta": map[string]any{
			"int": 1,
		},
	},
}

func main() {
	f, err := os.Create("save.pfcl")
	if err != nil {
		panic(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	encoder := pfcl.NewEncoder(f)
	err = encoder.Encode(MockData)
	if err != nil {
		panic(err)
	}
}
