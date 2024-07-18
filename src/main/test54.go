package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rows := len(matrix)
	fmt.Println("Number of rows:", rows)

	result := spiralOrder(matrix)
	fmt.Println("Spiral order:", result)
}

// 矩阵 中等题 螺旋矩阵
// 方法2 按层模拟，根据数学方法，一圈圈、一层层遍历
func spiralOrder(matrix [][]int) []int {
	// 异常情况处理
	if len(matrix) == 0 {
		return []int{}
	}
	// 返回值
	resultNums := make([]int, len(matrix[0])*len(matrix))
	//  ，bottom层数
	top, bottom := 0, len(matrix)-1
	//  ，列数
	left, right := 0, len(matrix[0])-1
	//外层循环条件
	index := 0
	for top <= bottom && left <= right { // 外层循环条件 层数起点小于重点 列数起点小于终点
		// 沿着顶边 从左到右遍历
		for i := left; i <= right; i++ {
			resultNums[index] = matrix[top][i]
			index++
		}
		top++ //递归层数变化
		for i := top; i <= bottom; i++ {
			// 右边从上到下遍历
			resultNums[index] = matrix[i][right]
			index++
		}
		right-- //递归列数变化
		if top <= bottom {
			// 下边从右到左遍历
			for i := right; i >= left; i-- {
				resultNums[index] = matrix[bottom][i]
				index++
			}
			bottom-- //递归层数右边变化
		}
		if left <= right {
			// 左边从下到上遍历
			for i := bottom; i >= top; i-- {
				resultNums[index] = matrix[i][left]
				index++
			}
			left++ //递归列数变化
		}
	}
	return resultNums
}
