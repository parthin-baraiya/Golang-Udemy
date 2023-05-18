package main

import (
	"os"
)

func main() {
	_, err := os.Open("no_file.txt")
	if err != nil {
		panic(err)
	}
}
