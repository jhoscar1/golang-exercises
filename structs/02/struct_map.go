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

	contacts := map[string]person{
		josh.last:  josh,
		jason.last: jason,
	}

	fmt.Println(contacts["Herren"])
	fmt.Println(contacts["Oscar"])

	for _, val := range contacts["Herren"].iceCream {
		fmt.Println(val)
	}

	for _, val := range contacts["Oscar"].iceCream {
		fmt.Println(val)
	}
}
