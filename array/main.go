package main

import (
	"fmt"
	"go-for-offer/array/algorithm"
)

func main() {
	var num int
	num = algorithm.FindSmallestNumInRotate([]int{3, 4, 5, 0, 2})
	fmt.Println(num)
}
