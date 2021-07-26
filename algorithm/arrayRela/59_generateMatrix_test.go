package arrayRela

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	// 上下行的边界
	top, bottom := 0, n-1
	// 左右列的边界
	left, right := 0, n-1
	// 填充的数值及最终的数值
	num := 1
	tar := n * n
	// 模拟顺时针填充数值
	for num <= tar {
		// 模拟从左往右填充上行
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++ // 更新上行的索引
		// 模拟从上往下填充右列
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right-- // 更新右列的索引
		// 模拟从右往左填充下行
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom-- // 更新下行的索引
		// 模拟从下往上填充左列
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++ // 更新左列的索引
	}
	return matrix
}
