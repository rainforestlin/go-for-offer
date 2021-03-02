package main

import "fmt"

//输出所有和为 S 的连续正数序列。例如和为 100 的连续序列有：
// [9, 10, 11, 12, 13, 14, 15, 16]
//[18, 19, 20, 21, 22]。

// 滑动窗口法
// 1、设定左右指针left/right，分别指向1和2，维护左右指针间的窗口和sum
// 2、如果sum等于target值，将左右指针间的数字作为一个数组加入最终结果，同时右移右指针
// 3、如果sum大于target值，缩小窗口，通过右移左指针，将左边的数字移出窗口
// 4、如果sum小于target值，增大窗口，通过右移右指针，将右边的数字加入窗口 5、左指针最多能移动到(target+1)/2的地方，因为至少需要两个数字，那么最大也就是中间的数字了

func findContinuousSequence(target int) [][]int {
	var result [][]int
	// 最小值为1+2=3
	if target < 3 {
		return result
	}
	left, right, mid := 1, 2, (target+1)/2
	sum := left + right
	for left < mid {
		if sum == target {
			var temp []int
			for i := left; i <= right; i++ {
				temp = append(temp, i)
			}
			result = append(result, temp)
			right++
			sum += right
		} else if sum < target {
			right++
			sum += right
		} else {
			sum -= left
			left++

		}
	}
	return result
}

func main() {
	fmt.Println(findContinuousSequence(9))
}
