package main

import "fmt"

func fcfunc(x int) func() int {
	return func() int {
		return x * x
	}
}

func main() {
	square4 := fcfunc(4)

	fmt.Println(square4())
}
