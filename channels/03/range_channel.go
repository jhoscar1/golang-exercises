package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	go gen(c)
	receive(c)

	fmt.Println("about to exit")
}

func gen(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for val := range c {
		fmt.Println(val, "was received from the channel")
	}
}
