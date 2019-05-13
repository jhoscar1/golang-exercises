package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, "and I'm", p.age)
}

func main() {

	jason := person{
		first: "Jason",
		last:  "Oscar",
		age:   27,
	}

	jason.speak()
}
