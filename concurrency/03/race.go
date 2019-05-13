package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count := 0

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			newCount := count
			runtime.Gosched()
			newCount++
			count = newCount
			fmt.Println(count)
			wg.Done()
		}()
	}
	wg.Wait()
}
