package main

import (
	"fmt"
	"sort"
)

// 给定一个二维数组，其每一行从左到右递增排序，从上到下也是递增排序。给定一个数，判断这个数是否在该二维数组中。
//Consider the following matrix:
//[
//[1,   4,  7, 11, 15],
//[2,   5,  8, 12, 19],
//[3,   6,  9, 16, 22],
//[10, 13, 14, 17, 24],
//[18, 21, 23, 26, 30]
//]
//
//Given target = 5, return true.
//Given target = 20, return false.

// 利用数组规律 Go：从左下角开始找，
// 利用递增关系，大于往右边找，小于往上找，超出返回false
func findNumberIn2DArray(matrix [][]int, target int) bool {
	y := len(matrix) - 1 // 获取右下角的y坐标
	x := 0               // 获取右下角x坐标
	for y > -1 {
		if x < len(matrix[y]) {
			if target < matrix[y][x] {
				y-- // 小于target，向上查找
			} else if target > matrix[y][x] {
				x++ // 大于 target 向右查找
			} else if target == matrix[y][x] {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

func findNumberIn2DArrayBySort(matrix [][]int, target int) bool {
	for _, nums := range matrix {
		i := sort.SearchInts(nums, target)
		if i < len(nums) && target == nums[i] {
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(findNumberIn2DArrayBySort(matrix, 31))
}
