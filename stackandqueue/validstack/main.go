package main

import "fmt"

func validStackSequences(pushed []int, popped []int) bool {
	// 辅助栈
	stack := make([]int, 0)
	i := 0
	for _, value := range pushed {
		stack = append(stack, value)
		for len(stack) != 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	poped := []int{4, 3, 5, 1, 2}
	fmt.Println(validStackSequences(pushed, poped))
}
