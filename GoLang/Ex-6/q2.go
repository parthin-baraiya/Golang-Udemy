package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no_file.txt")
	if err != nil {
		log.Println("err:=", err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}
