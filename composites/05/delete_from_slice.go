package main

import "fmt"

func main() {
	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(sl[:2], sl[6:]...)
	fmt.Println(y)
}
