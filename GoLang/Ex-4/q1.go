package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 40
	c <- 41

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("-----\n")
	fmt.Printf("%T\n", c)
}
