package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0

	var wg sync.WaitGroup
	var mut sync.Mutex
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			mut.Lock()
			newCount := count
			newCount++
			count = newCount
			fmt.Println(count)
			mut.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}
