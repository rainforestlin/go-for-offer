package algorithm

import (
	"sort"
	"strconv"
)

// FindNumbersSumEqualN 输出最小乘积的两个数
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
			end--
		}
		if a[start]+a[end] < sum {
			start++
		}
	}
	return result
}

// FindContinuousSequence 输出多组连续正数和为N
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

// MaxSubArraySum 一个整数数组中的元素有正有负，在该数组中找出一个连续子数组，要求该连续子数组中各元素的和最大，这个连续子数组便被称作最大连续子数组。比如数组{2,4,-7,5,2,-1,2,-4,3}的最大连续子数组为{5,2,-1,2}，最大连续子数组的和为5+2-1+2=8。
func MaxSubArraySum(nums []int) (result int) {
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

// GetNumbersOfNFirst 输出前N个
func GetNumbersOfNFirst(num []int, n int) (times int) {
	if len(num) == 0 {
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

// GetNumbersOfNSecond 输出前N个数
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

// SingleNumber 输出只出现了一次的元素
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

// FindSmallestNumInRotate 找到翻转数组中的最小数
func FindSmallestNumInRotate(nums []int) int {
	/**
	找到翻转数组中的最小数
	*/
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

// InversePairs 逆序对总数
func InversePairs(nums []int) (count int) {
	/**
	在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数P。并将P对1000000007取模的结果输出。即输出P%1000000007。
	输入描述:
	题目保证输入的数组中没有的相同的数字。
	数据范围：
	对于%50的数据,size<=10^4
	对于%75的数据,size<=10^5
	对于%100的数据,size<=2*10^5
	示例：
	输入
	1,2,3,4,5,6,7,0
	输出
	7
	*/
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	var copySlice = make([]int, len(nums))
	copy(copySlice, nums)
	count = mergeSort(nums, copySlice, 0, length-1)
	return
}

func mergeSort(slice, copySlice []int, start, end int) int {
	if start == end {
		return 0
	}
	mid := (start + end) / 2
	leftCount := mergeSort(slice, copySlice, start, mid) % 1000000007
	rightCount := mergeSort(slice, copySlice, mid+1, end) % 1000000007
	var count int
	i := mid
	j := end
	locCopy := end
	for i > start && j > mid {
		if slice[i] > slice[j] {
			count += j - mid
			locCopy--
			i--

			copySlice[locCopy] = slice[i]
			if count >= 1000000007 {
				count %= 1000000007
			}
		} else {
			locCopy--
			j--
			copySlice[locCopy] = slice[i]
		}
	}
	for ; i >= start; i-- {
		locCopy--
		copySlice[locCopy] = slice[i]
	}
	for ; j > mid; j-- {
		locCopy--
		copySlice[locCopy] = slice[j]
	}
	for index := start; index <= end; index++ {
		slice[index] = copySlice[index]
	}
	return (leftCount + rightCount + count) % 1000000007
}

// GetLeastNumbersBySortAll 输入n个整数，找出其中最小的K个数
func GetLeastNumbersBySortAll(input []int, k int) []int {
	/**
	输入n个整数，找出其中最小的K个数。例如输入4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4,。
	全部排序版
	*/
	if len(input) == 0 || k < 0 {
		return nil
	}
	if k >= len(input) {
		return input
	}
	sort.Ints(input)
	return input[:k]
}

// GetLeastNumbersByPartialSort 找到前K个值
func GetLeastNumbersByPartialSort(input []int, k int) []int {
	/**
	因为只要求拿到前k个值，并没有要求顺序，先假设原切片前k个值为结果，然后选取该结果中的最大值和后续进行比较
	如果后续的值比较小，就交换入结果切片。最终结果切片中的值为确为最小k个值
	*/
	if len(input) == 0 || k < 0 {
		return nil
	}
	if k >= len(input) {
		return input
	}

	result := input[0:k]
	for i := k; i < len(input); i++ {
		max := findMax(result)
		if result[max] > input[i] {
			result[max], input[i] = input[i], result[max]
		}
	}
	return result
}

func findMax(data []int) (max int) {
	max = 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[max] {
			max = i
		}
	}
	return max
}

// MoreThanHalfNumByCount 查找数组中某个数是否超过一半
func MoreThanHalfNumByCount(input []int) int {
	time := 0
	result := 0
	for i := 0; i < len(input); i++ {
		if time == 0 {
			result = input[i]
			time++
		} else if result == input[i] {
			time++
		} else if result != input[i] {
			time--
		}
	}
	if time == 0 {
		return -1
	}
	count := 0
	for _, num := range input {
		if num == result {
			count++
		}
	}
	if count < len(input)/2 {
		return -1
	}
	return result
}

// PrintMinNum 输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。例如输入数组{3，32，321}，则打印出这三个数字能排成的最小数字为321323。
func PrintMinNum(numbers []int) string {
	/**
	  输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
	  	例如输入数组{3，32，321}，则打印出这三个数字能排成的最小数字为321323。
	*/
	length := len(numbers)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if compare(numbers[i], numbers[j]) {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	var str string
	for _, v := range numbers {
		str += strconv.Itoa(v)
	}
	return str

}

func compare(a, b int) bool {
	m := strconv.Itoa(a) + strconv.Itoa(b)
	n := strconv.Itoa(b) + strconv.Itoa(a)
	mm, _ := strconv.Atoi(m)
	nn, _ := strconv.Atoi(n)

	return mm > nn
}

// Duplicate 查重
func Duplicate(num []int) int {
	length := len(num)
	for _, x := range num {
		if x >= length {
			x = x - length
		}
		if num[x] >= length {
			return x
		}
		num[x] = num[x] + length
	}
	return -1
}

// MaxSlideWindow 滑动窗口中的最大值
func MaxSlideWindow(nums []int, k int) []int {
	/**
	  定一个数组和滑动窗口的大小，找出所有滑动窗口里数值的最大值。
	  例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}；
	  针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个：
	  {[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。
	*/
	length := len(nums)
	if length == 0 || k <= 0 || k > length {
		return nil
	}
	var result []int
	if k == 1 {
		return nums
	}
	for i := 0; i <= length-k; i++ {
		temp := nums[i]
		for j := i + 1; j < i+k; j++ {
			if nums[j] > temp {
				temp = nums[j]
			}
		}
		result = append(result, temp)
	}

	return result
}

// MaxSlideWindowByCompareIndex 滑动窗口中的最大值
func MaxSlideWindowByCompareIndex(nums []int, k int) []int {
	length := len(nums)
	if length == 0 || k <= 0 || k > length {
		return nil
	}
	var maxNums []int
	max := -1
	for i := 0; i <= length-k; i++ {
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

//双指针查找最大值的索引
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

// ConstructMultiply 构建数积
func ConstructMultiply(a []int) []int {
	/**
	给定一个数组A[0,1,…,n-1],请构建一个数组B[0,1,…,n-1],其中B中的元素B[i]=A[0] * A[1] * … * A[i-1] * A[i+1] * … * A[n-1]。
	*/
	size := len(a)
	if size <= 1 {
		return a
	}
	left, right := make([]int, size), make([]int, size)
	left[0] = 1
	right[size-1] = 1

	for i := 1; i < size; i++ {
		left[i] = left[i-1] * a[i-1]
	}

	for i := size - 2; i >= 0; i-- {
		right[i] = right[i+1] * a[i+1]
	}
	b := make([]int, size)
	for i := 0; i < size; i++ {
		b[i] = left[i] * right[i]
	}
	return b
}
