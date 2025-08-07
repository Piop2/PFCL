package main

import (
	"os"

	"github.com/piop2/pfcl"
)

func main() {
	file, err := os.Open("test.pfcl")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	decoder := pfcl.NewDecoderFromFile(file)

	_, err = decoder.Decode()
	if err != nil {
		panic(err)
	}
}
