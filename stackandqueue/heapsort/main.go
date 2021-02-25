package main

import (
	"fmt"
)

type heap struct {
	arr *[]int
	len int
}

type minHeap heap
type maxHeap heap

// 对下标为i的节点进行堆化，n表示最后一个节点的下标
// 左子节点2i+1，右子节点 2i+2
func maxHeapf(arr []int, n, i int) {
	for {
		maxPos := i
		if 2*i+1 < n && arr[2*i+1] > arr[i] {
			maxPos = 2*i + 1
		}
		if 2*i+2 < n && arr[2*i+2] > arr[maxPos] {
			maxPos = 2*i + 2
		}
		if maxPos == i {
			break
		}
		arr[i], arr[maxPos] = arr[maxPos], arr[i]
		i = maxPos
	}
}

func minHeapf(arr []int, n, i int) {
	for {
		maxPos := i
		if 2*i+1 < n && arr[2*i+1] < arr[i] {
			maxPos = 2*i + 1
		}
		if 2*i+2 < n && arr[2*i+2] < arr[maxPos] {
			maxPos = 2*i + 2
		}
		if maxPos == i {
			break
		}
		arr[i], arr[maxPos] = arr[maxPos], arr[i]
		i = maxPos
		fmt.Println(arr)
	}
}

// 建堆，就是在原切片上操作，形成堆结构
//只要按照顺序，把切片下标为n/2到1的节点依次堆化，最后就会把整个切片堆化
func buildMaxHeap(arr []int) *maxHeap {
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		maxHeapf(arr, n, i)
	}
	return &maxHeap{
		arr: &arr,
		len: n,
	}
}

func buildMinHeap(arr []int) *minHeap {
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		minHeapf(arr, n, i)
	}
	return &minHeap{
		&arr,
		n,
	}
}

func (h *minHeap) Pop() int {
	if h.len < 1 {
		return -1
	}
	out := (*h.arr)[0]
	(*h.arr)[0] = (*h.arr)[h.len-1] // 最后一个元素移到第一个
	*h.arr = (*h.arr)[1:]
	h.len--
	minHeapf(*h.arr, h.len-2, 0)
	return out

}

func (h *maxHeap) Pop() int {
	if h.len < 1 {
		return -1
	}
	out := (*h.arr)[0]
	(*h.arr)[0] = (*h.arr)[h.len-1] // 最后一个元素移到第一个
	*h.arr = (*h.arr)[1:]
	h.len--
	maxHeapf(*h.arr, h.len-2, 0)
	return out
}

func (h *minHeap) Push(data int) {
	h.len++
	*h.arr = append(*h.arr, data)
	fmt.Println(*h.arr)
	fmt.Println(len(*h.arr))
	fmt.Println(h.len)
	i := h.len - 1
	for i/2 >= 0 && (*h.arr)[i/2] > (*h.arr)[i] {
		(*h.arr)[i], (*h.arr)[i/2] = (*h.arr)[i/2], (*h.arr)[i]
		i /= 2
	}
}

func (h *maxHeap) Push(data int) {
	h.len++
	*h.arr = append(*h.arr, data)
	i := h.len - 1
	for i/2 >= 0 && (*h.arr)[i/2] < (*h.arr)[i] {
		(*h.arr)[i], (*h.arr)[i/2] = (*h.arr)[i/2], (*h.arr)[i]
		i /= 2
	}
}

func (h *maxHeap) Sort() {
	arrLen := h.len
	for i := arrLen - 1; i >= 0; i-- {
		(*h.arr)[0], (*h.arr)[i] = (*h.arr)[i], (*h.arr)[0]
		arrLen -= 1
		maxHeapf(*h.arr, arrLen, 0)
		fmt.Println(*h.arr)
	}
	fmt.Println(*h.arr)
}

func (h *minHeap) Sort() {
	arrLen := h.len
	for i := h.len - 1; i >= 0; i-- {
		(*h.arr)[0], (*h.arr)[i] = (*h.arr)[i], (*h.arr)[0]
		arrLen -= 1
		minHeapf(*h.arr, arrLen, 0)
		fmt.Println(*h.arr)
	}
	fmt.Println(*h.arr)
}

func main() {
	m := []int{0, 1, 2, 3, 4, 5, 6}
	heap := buildMinHeap(m)
	fmt.Println(heap.arr)
	heap.Pop()
	fmt.Println(heap.arr)
	heap.Push(0)
	fmt.Println(heap.arr)

}
