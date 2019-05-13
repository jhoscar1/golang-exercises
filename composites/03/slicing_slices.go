package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(sl[:5])
	fmt.Println(sl[5:])
	fmt.Println(sl[2:7])
	fmt.Println(sl[1:6])
}
