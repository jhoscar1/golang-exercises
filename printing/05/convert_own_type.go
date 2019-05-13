package main

import "fmt"

type jason int

var age jason = 27
var newAge int

func main() {
	fmt.Println(age)
	fmt.Printf("%T\n", age)
	newAge = int(age)
	fmt.Printf("newAge has value %v and type %T age still has custom type %T", newAge, newAge, age)
}
