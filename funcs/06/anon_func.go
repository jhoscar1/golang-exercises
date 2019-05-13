package main

import "fmt"

func main() {
	y := func(x int) int {
		return x + 1
	}(4)

	fmt.Println(y)
}
