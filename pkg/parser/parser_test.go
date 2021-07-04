package parser

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	ch := ReadPorts("./test_sample/1.json")
	instances := 2
	seen := 0
	for p := range ch {
		fmt.Println(p)
		seen++
	}
	if instances != seen {
		t.Fatal("Expected ", instances, " but got ", seen)
	}
}
