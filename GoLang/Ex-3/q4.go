package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 40
	c <- 41
	fmt.Print(<-c)
}
