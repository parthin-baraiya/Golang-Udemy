package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send(eve, odd, quit)
	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Even channel:", v)
		case v := <-o:
			fmt.Println("Odd channel:", v)
		case v := <-q:
			fmt.Println("Quit channel:", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}
