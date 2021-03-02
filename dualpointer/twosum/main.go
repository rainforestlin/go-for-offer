package main

// 在有序数组中找出两个数，使得和为给定的数 S。如果有多对数字的和等于 S，
//输出两个数的乘积最小的。

func twoSumByMap(nums []int, target int) []int {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{target - v, v}
		} else {
			m[v] = struct{}{}
		}
	}
	return nil
}

func twoSumByDualPointer(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{nums[left], nums[right]}
		}
	}
	return nil
}
