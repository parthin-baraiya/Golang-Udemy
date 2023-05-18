package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no_file.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("deferred functions don't run, When os.Exit() is called.")
}
