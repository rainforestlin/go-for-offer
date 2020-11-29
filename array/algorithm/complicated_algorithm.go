package algorithm

import "container/heap"

type MedianFinder struct {
	minHeap *heapMin
	maxHeap *heapMin
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: new(heapMin),
		maxHeap: new(heapMin),
	}
}

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

func (m *MedianFinder) FindMedian() float64 {
	if m.minHeap.Len() > m.maxHeap.Len() {
		return float64((*m.minHeap)[0])
	} else if m.minHeap.Len() < m.maxHeap.Len() {
		return -float64((*m.maxHeap)[0])
	}
	return float64((*m.minHeap)[0]-(*m.maxHeap)[0]) / float64(2)
}

type heapMin []int

func (h heapMin) Len() int {
	return len(h)
}

func (h heapMin) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *heapMin) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heapMin) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *heapMin) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}

func InitHeap(a []int) *heapMin {
	h := new(heapMin)
	for _, v := range a {
		heap.Push(h, v)
	}
	return h
}

func HeapSort(h *heapMin) []int {
	var result []int
	for range *h {
		result = append(result, heap.Pop(h).(int))
	}
	return result
}
