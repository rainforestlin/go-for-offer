package algorithm

func FindNumbersSumEqualN(a []int, sum int) []int {
	/**
	  输入一个递增排序的数组和一个数字S，在数组中查找两个数，使得他们的和正好是S，如果有多对数字的和等于S，输出两个数的乘积最小的。
	*/
	var result []int
	length := len(a)
	if length == 0 {
		return result
	}
	start := 0
	end := length - 1
	for start < end {
		if a[start]+a[end] == sum {
			// 两个数相隔越远，乘积越小
			result = append(result, start, end)
			break
		}
		if a[start]+a[end] > sum {
			end -= 1
		}
		if a[start]+a[end] < sum {
			start += 1
		}
	}
	return result
}

func FindContinuousSequence(sum int) (result [][]int) {
	/**
	小明很喜欢数学,有一天他在做数学作业时,要求计算出9~16的和,他马上就写出了正确答案是100。
	但是他并不满足于此,他在想究竟有多少种连续的正数序列的和为100(至少包括两个数)。
	没多久,他就得到另一组连续正数和为100的序列:18,19,20,21,22。

	使用等差数列求和公式
	*/
	low := 1
	high := 2
	for low < high {
		cur := (low + high) * (high - low + 1) / 2
		if cur == sum {
			var temp []int
			for i := low; i <= high; i++ {
				temp = append(temp, i)
			}
			result = append(result, temp)
			low++
		}
		if cur < sum {
			high++
		}
		if cur > sum {
			low++
		}
	}
	return
}

func MaxSubArraySum(nums []int) (result int) {
	/**
	一个整数数组中的元素有正有负，在该数组中找出一个连续子数组，要求该连续子数组中各元素的和最大，这个连续子数组便被称作最大连续子数组。
	比如数组{2,4,-7,5,2,-1,2,-4,3}的最大连续子数组为{5,2,-1,2}，最大连续子数组的和为5+2-1+2=8。
	*/
	result = nums[0]
	sum := 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		if result < sum {
			result = sum
		}
	}
	return
}

func GetNumbersOfNFirst(num []int, n int) (times int) {
	if num == nil || len(num) == 0 {
		return 0
	}
	start := 0
	end := len(num) - 1
	mid := 0
	for start <= end {
		mid = (end - start) / 2
		if n < num[mid] {
			end = mid - 1
		} else if n > num[mid] {
			start = mid + 1
		} else {
			times++
			break
		}
	}
	for start = mid - 1; (start >= 0) && (num[start] == n); start-- {
		times++
	}
	for start = mid + 1; (start < len(num)) && (num[start] == n); start++ {
		times++
	}
	return
}

func GetNumbersOfNSecond(nums []int, n int) int {
	if nums == nil {
		return 0
	}
	firstN := fetchFirstN(nums, n, 0, len(nums)-1)
	lastN := fetchLastN(nums, n, 0, len(nums)-1)
	if firstN != -1 && lastN != -1 {
		return lastN - firstN + 1
	}
	return 0
}

func fetchFirstN(elem []int, n, start, end int) int {
	// 使用递归
	if start > end {
		return -1
	}
	mid := (end + start) / 2
	if elem[mid] > n {
		return fetchFirstN(elem, n, start, mid-1)
	} else if elem[mid] < n {
		return fetchFirstN(elem, n, mid+1, end)
	} else if (mid-1 >= 0) && (elem[mid-1] == n) {
		return fetchFirstN(elem, n, start, mid-1)
	} else {
		return mid
	}
}

func fetchLastN(elem []int, n, start, end int) int {
	mid := (start + end) / 2
	for start <= end {
		if elem[mid] > n {
			end = mid - 1
		} else if elem[mid] < n {
			start = mid + 1
		} else if (mid+1 < len(elem)) && (elem[mid+1] == n) {
			start = mid + 1
		} else {
			return mid
		}
		mid = (start + end) / 2
	}
	return -1
}

func SingleNumber(nums []int) int {
	/**
	给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
	*/
	result := 0
	for _, x := range nums {
		result = x ^ result
	}
	return result
}

func FindSmallestNumInRotate(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	start := 0
	end := len(nums) - 1
	if nums[start] < nums[end] {
		return nums[0]
	}
	for start < end {
		mid := (start + end) / 2
		if nums[mid] > nums[end] {
			start = mid + 1
		}
		if nums[mid] == nums[end] {
			end--
		}
		if nums[mid] < nums[end] {
			end = mid
		}
	}
	return nums[start]
}
