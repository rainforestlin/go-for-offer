package main

func maxSlideWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 || k > len(nums) {
		return nil
	}
	var maxNums []int
	max := -1
	for i := 0; i <= len(nums)-k; i++ {
		l := i
		r := i + k - 1
		if max == -1 || max == l-1 {
			max = getMax(nums, l, r)
		} else {
			if nums[r] > nums[max] {
				max = r
			}
		}
		maxNums = append(maxNums, nums[max])
	}
	return maxNums
}

func getMax(nums []int, left, right int) int {
	for left < right {
		if nums[left] > nums[right] {
			right--
		} else {
			left++
		}
	}
	return left
}
