package main

import (
	"fmt"
)

func main() {
	b := 42
	fmt.Printf("decimal: %d, binary: %b, hexadecimal: %#x\n", b, b, b)
	c := b << 1
	fmt.Printf("decimal: %d, binary: %b, hexadecimal: %#x\n", c, c, c)
}
