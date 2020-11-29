package main

import (
	"fmt"
	"go-for-offer/array/algorithm"
)

func main() {
	x := algorithm.InitHeap([]int{3, 1, 6, 7, 9, 2, 4, 5})
	fmt.Println(algorithm.HeapSort(x))
}
