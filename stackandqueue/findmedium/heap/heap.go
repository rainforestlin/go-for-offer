package heap

type heap struct {
	arr *[]int
	len int
}

type MinHeap heap
type MaxHeap heap

// 对下标为i的节点进行堆化，n表示最后一个节点的下标
// 左子节点2i+1，右子节点 2i+2
func MaxHeapf(arr []int, n, i int) {
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

func MinHeapf(arr []int, n, i int) {
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
	}
}

// 建堆，就是在原切片上操作，形成堆结构
//只要按照顺序，把切片下标为n/2到1的节点依次堆化，最后就会把整个切片堆化
func BuildMaxHeap(arr []int) *MaxHeap {
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		MaxHeapf(arr, n, i)
	}
	return &MaxHeap{
		arr: &arr,
		len: n,
	}
}

func BuildMinHeap(arr []int) *MinHeap {
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		MinHeapf(arr, n, i)
	}
	return &MinHeap{
		&arr,
		n,
	}
}

func (h *MinHeap) Pop() int {
	if h.len < 1 {
		return -1
	}
	out := (*h.arr)[0]
	if h.len == 1 {
		h.len -= 1
		*h.arr = []int{}
		return out
	} else {
		(*h.arr)[0] = (*h.arr)[h.len-1] // 最后一个元素移到第一个
		*h.arr = (*h.arr)[1:]
		h.len--
		MinHeapf(*h.arr, h.len-2, 0)
	}

	return out

}

func (h *MaxHeap) Pop() int {
	if h.len < 1 {
		return -1
	}
	out := (*h.arr)[0]

	if h.len == 1 {
		h.len -= 1
		*h.arr = []int{}
		return out
	} else {
		(*h.arr)[0] = (*h.arr)[h.len-1] // 最后一个元素移到第一个
		*h.arr = (*h.arr)[1:]
		h.len--
		MaxHeapf(*h.arr, h.len-2, 0)
	}
	return out
}

func (h *MinHeap) Push(data int) {
	h.len++
	*h.arr = append(*h.arr, data)
	i := h.len - 1
	for i/2 >= 0 && (*h.arr)[i/2] > (*h.arr)[i] {
		(*h.arr)[i], (*h.arr)[i/2] = (*h.arr)[i/2], (*h.arr)[i]
		i /= 2
	}
}

func (h *MaxHeap) Push(data int) {
	h.len++
	*h.arr = append(*h.arr, data)
	i := h.len - 1
	for i/2 >= 0 && (*h.arr)[i/2] < (*h.arr)[i] {
		(*h.arr)[i], (*h.arr)[i/2] = (*h.arr)[i/2], (*h.arr)[i]
		i /= 2
	}
}

func (h *MaxHeap) Sort() {
	arrLen := h.len
	for i := arrLen - 1; i >= 0; i-- {
		(*h.arr)[0], (*h.arr)[i] = (*h.arr)[i], (*h.arr)[0]
		arrLen -= 1
		MaxHeapf(*h.arr, arrLen, 0)
	}
}

func (h *MinHeap) Sort() {
	arrLen := h.len
	for i := h.len - 1; i >= 0; i-- {
		(*h.arr)[0], (*h.arr)[i] = (*h.arr)[i], (*h.arr)[0]
		arrLen -= 1
		MinHeapf(*h.arr, arrLen, 0)
	}
}

func (h *MinHeap) Len() int {
	return h.len
}
func (h *MaxHeap) Len() int {
	return h.len
}

func (h *MinHeap) Peek() int {
	return (*h.arr)[0]
}
func (h *MaxHeap) Peek() int {
	return (*h.arr)[0]
}
