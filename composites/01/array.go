package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for ind := range arr {
		fmt.Println(ind)
	}

	fmt.Printf("%T", arr)
}
