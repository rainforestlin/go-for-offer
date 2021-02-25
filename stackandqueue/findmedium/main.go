package main

import (
	"fmt"
	"go-for-offer/stackandqueue/findmedium/heap"
)

// 如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，
//那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，
//那么中位数就是所有数值排序之后中间两个数的平均值。

// 使用大顶堆和小顶堆进行

type MediumFinder struct {
	minHeap heap.MinHeap
	maxHeap heap.MaxHeap
	length  int
}

func Constructor() MediumFinder {
	return MediumFinder{
		minHeap: *heap.BuildMinHeap([]int{}),
		maxHeap: *heap.BuildMaxHeap([]int{}),
		length:  0,
	}
}

func (m *MediumFinder) AddNum(num int) {
	m.length++
	if m.minHeap.Len() == m.maxHeap.Len() {
		m.minHeap.Push(num)
		m.maxHeap.Push(m.minHeap.Pop())
	} else {
		m.maxHeap.Push(num)
		m.minHeap.Push(m.maxHeap.Pop())
	}

}
func (m *MediumFinder) FindMedian() float64 {
	if m.minHeap.Len() == m.maxHeap.Len() {
		return (float64(m.minHeap.Peek() + m.maxHeap.Peek())) / 2
	} else {
		return float64(m.maxHeap.Peek())
	}
}

func main() {
	h := Constructor()
	h.AddNum(1)
	h.AddNum(2)
	h.AddNum(3)
	fmt.Println(h.FindMedian())
}
