package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Our main routine")

	wg.Add(2)

	go func() {
		fmt.Println("This is running concurrently.")
		wg.Done()
	}()

	go func() {
		fmt.Println("This is also running concurrently.")
		wg.Done()
	}()

	fmt.Println(runtime.NumGoroutine())
	fmt.Println("End of main routine")
	wg.Wait()
}
