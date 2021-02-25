package main

import "fmt"

// 输入整数数组 arr ，找出其中最小的 k 个数。
//例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
// 利用快速排序思想，partition函数每次会定位一个元素，并返回其下标。假设该元素为x，位于x之前的都小于等于x，位于x之后的元素都大于等于x。
//我们的目的是得到前k小的元素，那么只需利用partition函数定位到下标为k的元素，在该元素之前的就是我们的目标。
//
func topK(arr []int, k int) []int {
	if len(arr) == 0 || len(arr) <= k {
		return arr
	}
	left := 0
	right := len(arr) - 1
	index := partation(arr, left, right)
	for index != k-1 {
		if index > k-1 {
			right = index - 1
			index = partation(arr, left, right)
		} else {
			left = index + 1
			index = partation(arr, left, right)
		}
	}
	result := arr[:k]
	return result
}

func partation(arr []int, left, right int) int {
	if left == right {
		return right
	}
	if left < right {
		//	1.base
		base := arr[left]
		l := left
		r := right
		for {
			// 如果右侧大于base值，则右指针左移
			for arr[r] >= base && r > l {
				r--
			}
			// 如果左指针小于base，则左指针右移
			for arr[l] <= base && l < r {
				l++
			}

			if l >= r {
				break
			}
			arr[l], arr[r] = arr[r], arr[l]

		}
		arr[left] = arr[l]
		arr[l] = base
		return l
	}
	return -1
}

func main() {
	arr := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(topK(arr, 4))
}
