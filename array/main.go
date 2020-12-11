package main

import (
	"fmt"
	"go-for-offer/array/algorithm"
)

func main() {
	x := algorithm.SearchMatrix([][]int{{1,2,3,4},{5,6,7,8}}, 9)
	fmt.Println(x)
}
