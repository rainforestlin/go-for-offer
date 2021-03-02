package main

//冒泡排序
//最差时间复杂度O(n^2)，最好时间复杂度O(n)，平均时间复杂度O(n^2)
//空间复杂度O(1)
//是一种稳定算法
func BubbleSort(arr []int) {
	swapped := true
	arrLen := len(arr)
	x := 0
	for swapped {
		x = x + 1
		swapped = false
		for i := 0; i < arrLen-1; i++ {
			//fmt.Printf("第%d轮-第%d次：%v", x, i, arr)
			//fmt.Println()
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}
}

//选择排序
//最好时间复杂度O(n^2)，最差时间复杂度O(n^2)，平均时间复杂度O(n^2)
//空间复杂度O(1)
//是一种不稳定的排序算法
func SelectionSort(arr []int) {
	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++ {
		minIndex := i
		for j := i + 1; j < arrLen; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
		//fmt.Printf("第%d次：%v", i, arr)
		//fmt.Println()
	}
}

//插入算法
//最好时间复杂度O(n)，最差时间复杂度O(n^2)，平均时间复杂度O(n^2)
//空间复杂度为O(1)
//插入算法是一种稳定的排序算法
func InsertionSort(arr []int) {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		selected := arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] > selected {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				arr[j+1] = selected
				break
			}
		}
	}
}

func InsertionSortWithoutExchange(arr []int) {
	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		selected := arr[i]
		var j int
		for j = i - 1; j >= 0 && arr[j] > selected; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = selected
	}
}

//希尔排序

func ShellSort(arr []int) {
	for step := len(arr) / 2; step > 0; step /= 2 {
		for i := step; i < len(arr); i++ {
			for j := i - step; j >= 0 && arr[j+step] < arr[j]; j -= step {
				arr[j], arr[j+step] = arr[j+step], arr[j]
			}
		}
	}
}

//归并排序
//最好时间复杂度O(nlogn)，最差时间复杂度O(nlogn)，平均时间复杂度O(nlogn)
//空间复杂度为O(n)
//归并算法是一种稳定的排序算法
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])
	return merge(left, right)
}
func merge(a, b []int) []int {
	aLen, bLen := len(a), len(b)
	result := make([]int, aLen+bLen)
	index := 0
	i, j := 0, 0
	for i < aLen && j < bLen {
		if a[i] < b[j] {
			result[index] = a[i]
			i++
		} else {
			result[index] = b[j]
			j++
		}
		index++
	}
	for i != aLen {
		result[index] = a[i]
		index++
		i++
	}
	for j != bLen {
		result[index] = b[j]
		index++
		j++
	}
	return result
}

//快速排序
//最好时间复杂度O(nlogn)，最差时间复杂度O(n^2)，平均时间复杂度O(nlogn)
//空间复杂度为O(logn)~O(n)
//快速排序是一种不稳定的排序算法
func NormalQuickSort(arr []int) {
	normalQuickSort(arr, 0, len(arr)-1)
}

func normalQuickSort(arr []int, begin, end int) {
	if begin < end {
		//	进行切分
		loc := partition(arr, begin, end)
		//	对左边进行快排
		normalQuickSort(arr, begin, loc-1)
		//	对右边进行快排
		normalQuickSort(arr, loc+1, end)
	}
}
func partition(arr []int, begin, end int) (loc int) {
	// 将arr[begin]作为基准数
	loc = begin + 1
	j := end
	for loc < j {
		// 与基准数进行比较
		if arr[loc] > arr[begin] {
			// 交换，将比基准数大的数往右边挪
			arr[loc], arr[j] = arr[j], arr[loc]
			j--
		} else {
			loc++
		}
	}
	// 当挑出循环的时候，loc+1:end比基准数大，begin+1:loc-1比基准数小
	if arr[loc] >= arr[begin] {
		loc--
	}
	arr[begin], arr[loc] = arr[loc], arr[begin]
	return
}
