package algorithm

import "container/heap"

// MedianFinder 查找中位数的大小堆
type MedianFinder struct {
	minHeap *HeapMin
	maxHeap *HeapMin
}

// Constructor 查找中位数的构造函数
/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: new(HeapMin),
		maxHeap: new(HeapMin),
	}
}

// AddNum 往堆里增加数字
func (m *MedianFinder) AddNum(num int) {

	if m.minHeap.Len() > 0 && num > (*m.minHeap)[0] {
		heap.Push(m.minHeap, num)
	} else {
		heap.Push(m.maxHeap, -num)
	}

	if m.minHeap.Len()-m.maxHeap.Len() == 2 {
		heap.Push(m.maxHeap, -(heap.Pop(m.minHeap)).(int))
	} else if m.maxHeap.Len()-m.minHeap.Len() == 2 {
		heap.Push(m.minHeap, -(heap.Pop(m.maxHeap)).(int))
	}
}

// FindMedian 查找中位数
func (m *MedianFinder) FindMedian() float64 {
	if m.minHeap.Len() > m.maxHeap.Len() {
		return float64((*m.minHeap)[0])
	} else if m.minHeap.Len() < m.maxHeap.Len() {
		return -float64((*m.maxHeap)[0])
	}
	return float64((*m.minHeap)[0]-(*m.maxHeap)[0]) / float64(2)
}

// HeapMin 最小堆
type HeapMin []int

func (h HeapMin) Len() int {
	return len(h)
}

func (h HeapMin) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *HeapMin) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Push 往堆里添加
func (h *HeapMin) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 从堆里删除
func (h *HeapMin) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}

// InitHeap 初始化堆
func InitHeap(a []int) *HeapMin {
	h := new(HeapMin)
	for _, v := range a {
		heap.Push(h, v)
	}
	return h
}

// HeapSort 堆排序
func HeapSort(h *HeapMin) []int {
	var result []int
	for range *h {
		result = append(result, heap.Pop(h).(int))
	}
	return result
}
