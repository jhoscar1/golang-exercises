package main

import "fmt"

type jason int

var age jason = 27

func main() {
	fmt.Println(age)
	fmt.Printf("%T\n", age)
	age = 42
	fmt.Println(age)
}
