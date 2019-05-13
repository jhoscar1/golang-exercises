package main

import "fmt"

func main() {
	type person struct {
		first    string
		last     string
		iceCream []string
	}

	jason := person{
		first:    "Jason",
		last:     "Oscar",
		iceCream: []string{"chunky monkey", "phish food"},
	}

	josh := person{
		first:    "Josh",
		last:     "Herren",
		iceCream: []string{"vanilla", "chocolate"},
	}

	for _, val := range jason.iceCream {
		fmt.Println(val)
	}

	for _, val := range josh.iceCream {
		fmt.Println(val)
	}
}
