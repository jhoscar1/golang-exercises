package main

import "fmt"

func main() {
	s := 16
	eq := s == 16
	gt := s > 15
	gte := s >= 16
	lt := s < 16
	lte := s <= 17
	neq := s != 17

	fmt.Println(eq)
	fmt.Println(gt)
	fmt.Println(gte)
	fmt.Println(lt)
	fmt.Println(lte)
	fmt.Println(neq)
}
