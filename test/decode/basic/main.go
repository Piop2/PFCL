package main

import (
	"fmt"
	"strings"

	"github.com/piop2/pfcl"
)

var MockData = "# This is Comment\n\n" +
	"str = \"BOINK\" # This is string\n" +
	"cool = true # ddd\n" +
	"list = {\n" +
	"    true,\n" +
	"    -2,\n" +
	"    \"d\",\n" +
	"    {1, 3.4},\n" +
	"}\n\n" +
	"[server] # This is Table\n" +
	"ip = \"127.0.0.1:3000\"\n" +
	"c = 1\n" +
	"cd = 1.2\n\n" +
	"[server.beta]\n" +
	"ip = \"127.0.0.1:5000\"\n\n" +
	"[server.alpha]\n" +
	"ip = \"127.0.0.1:8000\""

func main() {
	decoder := pfcl.NewDecoder(strings.NewReader(MockData))

	data := make(map[string]any)

	decodeErr := decoder.Decode(&data)
	if decodeErr != nil {
		panic(decodeErr)
	}

	fmt.Println(data)
}
