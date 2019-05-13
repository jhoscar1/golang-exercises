package main

import "fmt"

type person struct {
	name string
	age  int
}

type speaker interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.name)
}

func saySomething(s speaker) {
	s.speak()
}

func main() {
	// p1 := person{"Jason", 1}
	p2 := &person{"June", 2}

	saySomething(p2)
	// saySomething(p1) this fails
}
