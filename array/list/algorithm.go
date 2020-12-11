package list

import (
	"container/list"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func DetectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	for {
		if fast.Next != nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}
	}
	// 快指针之前比慢指针多走了头指针到环入口节点的距离+一个环
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

// DeleteDulicatates 对有序链表进行去重
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	for p != nil {
		val := p.Val
		for p.Next != nil && p.Next.Val == val {
			p.Next = p.Next.Next
		}
		p = p.Next
	}
	return head
}

// FmtListNode 将一个链表从末尾往前输出
func FmtListNode(head *ListNode) {
	l := list.New()
	for ; head != nil; head = head.Next {
		l.PushFront(head.Val)
	}

	for item := l.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value)
	}
}

// FindFirstCommonNode 输入两个链表，找出它们的第一个公共节点。
func FindFirstCommonNode(a, b *ListNode) *ListNode {
	maps := make(map[*ListNode]int)
	for a != nil {
		maps[a] = a.Val
		a = a.Next
	}
	for b != nil {
		_, ok := maps[b]
		if ok {
			return b
		}
		b = b.Next
	}
	return nil
}
