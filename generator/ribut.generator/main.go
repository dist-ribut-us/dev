package main

import (
	"fmt"
	"github.com/dist-ribut-us/dev/generator"
	"os"
)

func main() {
	g, err := generator.Read(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(g.Generate())
}
