package main

// 按顺时针的方向，从外到里打印矩阵的值。
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	step := 0
	size := len(matrix) * len(matrix[0])
	//	定义四个方向断点
	top, bottom, right, left := 0, len(matrix)-1, len(matrix[0])-1, 0
	nums := make([]int, size)
	for step < size {
		//	left -> right
		for i := left; i <= right && step < size; i++ {
			nums[step] = matrix[top][i]
			step++
		}
		top++
		// top -> bottom
		for i := top; i <= bottom && step < size; i++ {
			nums[step] = matrix[i][right]
			step++
		}
		right--
		// right -> left
		for i := right; i >= left && step < size; i-- {
			nums[step] = matrix[bottom][i]
			step++
		}
		bottom--
		//	bottom->top
		for i := bottom; i >= top && step < size; i-- {
			nums[step] = matrix[i][left]
			step++
		}
		left++
	}
	return nums
}
