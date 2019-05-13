package main

import (
	"fmt"
	"sort"
)

func main() {
	numList := []int{2, 1, 4, 3, 5}
	nameList := []string{"Zach", "June", "Jason", "Henry", "Josh"}
	sort.Ints(numList)
	sort.Strings(nameList)
	fmt.Println(numList)
	fmt.Println(nameList)
}
