package main

import "fmt"

//从尾到头反过来打印出每个结点的值。

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	return appendData(head)
}

func appendData(head *ListNode) []int {
	if head.Next != nil {
		list := appendData(head.Next)
		list = append(list, head.Val)
		return list
	}

	return []int{head.Val}
}

func insert(tail *ListNode, next *ListNode) {
	tail.Next = next
	tail = next
}

func main() {
	var head = new(ListNode)
	var tail *ListNode
	tail = head
	for i := 1; i < 10; i++ {
		var node = ListNode{Val: i}
		tail.Next = &node
		tail = &node
	}
	currentNode := head
	for {

		if currentNode == nil {
			break
		}
		fmt.Println(currentNode.Val)
		currentNode = currentNode.Next
	}
}
