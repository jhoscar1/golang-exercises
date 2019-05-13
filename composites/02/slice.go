package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for ind := range sl {
		fmt.Println(ind)
	}

	fmt.Printf("%T", sl)
}
