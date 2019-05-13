package main

import "fmt"

func main() {
	ch := place()

	for v := 0; v < 100; v++ {
		fmt.Println(<-ch)
	}
}

func place() <-chan int {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				ch <- j
			}
		}()
	}

	return ch
}
