package main

import "fmt"

var (
	x = 42
	y = "James Bond"
	z = true
)

func main() {
	str := fmt.Sprintf("%d %s %t", x, y, z)
	fmt.Printf(str)
}
