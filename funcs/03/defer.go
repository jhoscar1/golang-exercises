package main

import "fmt"

func main() {

	defer deferred()

	fmt.Println("I'm run first")

}

func deferred() {
	fmt.Println("I'm run later!")
}
