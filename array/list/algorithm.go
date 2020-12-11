package list

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
