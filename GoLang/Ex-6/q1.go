package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no_file.txt")
	if err != nil {
		fmt.Println("err:=", err)
		log.Println("err happened", err)
	}
}
