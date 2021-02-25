package main

import (
	"fmt"
	"sort"
)

// 题目描述
// 在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内。
// 数组中某些数字是重复的，但不知道有几个数字是重复的，也不知道每个数字重复几次。
// 请找出数组中任意一个重复的数字。
// Input:
// {2, 3, 1, 0, 2, 5}
//
// Output:
// 2

// 由于只需要找出数组中任意一个重复的数字，因此遍历数组，遇到重复的数字即返回。
// 为了判断一个数字是否重复遇到，使用 map 存储已经遇到的数字，如果遇到的一个数字已经在 map 中，则当前的数字是重复数字。
func findRepeatNumByMap(arr []int) int {
	m := make(map[int]int)
	for _, num := range arr {
		if m[num] != 0 {
			return num
		}
		m[num] += 1
	}
	return -1
}

// 排序，看前后元素是否相等
func findRepeatNumBySort(arr []int) int {
	sort.Ints(arr)
	for i, numsSize := 0, len(arr); i < numsSize-1; i++ {
		if arr[i] == arr[i+1] {
			return arr[i]
		}
	}
	return -1
}

func findRepeatNum(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

func main() {
	res := findRepeatNum([]int{1, 2, 3, 4, 5, 1})
	fmt.Println(res)
}
