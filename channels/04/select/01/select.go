package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := make(chan int)
	go genC(c)
	go genQ(q)
	receive(c, q)

	fmt.Println("about to exit")
}

func genC(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
}

func genQ(q chan<- int) {
	for j := 0; j < 100; j++ {
		q <- j
	}
}

func receive(c, q <-chan int) {
	for i := 0; i < 200; i++ {
		select {
		case val := <-c:
			fmt.Println("value returned from c", val)
		case val := <-q:
			fmt.Println("value returned from q", val)
		}
	}
	fmt.Println("Channels done")
}
