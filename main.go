package main

import (
	"0xdive/internal/input"
	"fmt"
)

func main() {
	p := input.NewParser("     d     45346      ")
	fmt.Println(p.Parse())
}
