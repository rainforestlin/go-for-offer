package main

func MergeSortUpToDown(s []int) {
	aux := make([]int, len(s)) //辅助切片
	mergeSortUpToDown(s, aux, 0, len(s)-1)
}

func mergeSortUpToDown(s, aux []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := (lo + hi) >> 1
	mergeSortUpToDown(s, aux, lo, mid)
	mergeSortUpToDown(s, aux, mid+1, hi)
	if s[mid] != s[mid+1] {
		Merge(s, aux, lo, mid, hi)
	}
}

func MergeSortDownToUp(s []int) {
	aux := make([]int, len(s))
	n := len(s)
	for length := 1; length < n; length *= 2 {
		for low := 0; low < n-length; low += 2 * length {
			Merge(s, aux, low, low+length-1, min(low+2*length-1, n-1))
		}
	}
}

//归并操作
func Merge(s, aux []int, low, mid, high int) {
	for k := low; k <= high; k++ {
		aux[k] = s[k]
	}
	i, j := low, mid+1
	for k := low; k <= high; k++ {
		//左半边取尽
		if i > mid {
			s[k] = aux[j]
			j++
		} else if j > high { // 右半边取尽
			s[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			s[k] = aux[j]
			j++
		} else {
			s[k] = aux[i]
			i++
		}
	}
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
