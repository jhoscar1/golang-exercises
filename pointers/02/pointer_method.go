package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	(*p).first = "new first name"
}

func main() {
	jason := person{"jason", "oscar"}
	fmt.Printf("%T\n", jason)
	changeMe(&jason)
	fmt.Println(jason)
}
