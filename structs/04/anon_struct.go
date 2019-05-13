package main

import "fmt"

func main() {
	jason := struct {
		first string
		last  string
	}{
		first: "Jason",
		last:  "Oscar",
	}

	fmt.Println(jason)
}
