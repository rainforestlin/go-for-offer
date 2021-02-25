package main

import "fmt"

//实现一个包含 min() 函数的栈，该方法返回当前栈中最小的值。
//栈
//
//特点：先进后出
//题目中需要实现的方法
//push()：在数组的末属添加一个或多个元素（压入栈中）
//top()：返回数组中的最后一个元素（栈顶元素），但不出栈
//pop()：删除数组中的最后一个元素（栈顶元素出栈）
//包含min函数的栈，其实就是比栈多了一个min方法，调用则返回栈中的最小值。
//
//所以只需要增加一个辅助栈来存储较小的值，调用min()返回辅助栈的栈top即为最小值。
type MinStack struct {
	nums []int // 存储栈
	min  []int // 辅助栈
}

func Constructor() MinStack {
	return MinStack{
		make([]int, 0, 5),
		make([]int, 0, 5),
	}
}

func (m *MinStack) Push(x int) {
	m.nums = append(m.nums, x)
	if len(m.min) == 0 {
		m.min = append(m.min, x)
	} else if m.min[len(m.min)-1] < x {
		m.min = append(m.min, m.min[len(m.min)-1])
	} else {
		m.min = append(m.min, x)
	}
}
func (m *MinStack) Pop() {
	m.nums = m.nums[:len(m.nums)-1]
	m.min = m.min[:len(m.min)-1]
}

func (m *MinStack) Top() int {
	return m.nums[len(m.nums)-1]
}

func (m *MinStack) Min() int {
	return m.min[len(m.min)-1]
}

func main() {
	stack := Constructor()
	stack.Push(4)
	fmt.Println(stack.min)
	stack.Push(2)
	fmt.Println(stack.min)
	stack.Push(3)
	fmt.Println(stack.min)
	fmt.Println(stack.nums)

}
