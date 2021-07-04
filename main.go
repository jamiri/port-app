package main

import (
	"fmt"
	"github.com/jamiri/port-app/pkg/parser"
)

func main() {

	c := parser.ReadPorts("./pkg/parser/ports.json")

	for p := range c {
		fmt.Println(p)
	}
}
