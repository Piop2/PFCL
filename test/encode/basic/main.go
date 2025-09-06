package main

import (
	"fmt"

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
	b, err := pfcl.Marshal(MockData)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
