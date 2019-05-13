package main

import "fmt"

func main() {
	slint := []int{1, 2, 3, 4, 5}

	sum1 := sumInt(slint...)
	sum2 := sumSlice(slint)

	fmt.Println(sum1, sum2)
}

func sumInt(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func sumSlice(sliceNums []int) int {
	return sumInt(sliceNums...)
}
