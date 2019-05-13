package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	bufferedChan := make(chan int, 1)

	// ANON func
	go func() {
		c <- 42
	}()

	// buffered channel won't block
	bufferedChan <- 100

	fmt.Println(<-c)
	fmt.Println(<-bufferedChan)
}
