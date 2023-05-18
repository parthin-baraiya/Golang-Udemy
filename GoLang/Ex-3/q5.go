package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 40
	c <- 41

	fmt.Print(<-c)
	fmt.Print(<-c)
}
