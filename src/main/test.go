package main

import (
	. "MyLikou/src/MyStruct"
)

func invertTreeTest(matrix [][]int) []int {
	// 中等题 螺旋矩阵
	// 方法1 按层模拟 好理解
	// 返回值
	resultNum := make([]int, len(matrix)*len(matrix[0]))
	// 一圈一圈模拟，通过记录圈的复杂度来循环，内层算法都是一样的
	// 行数 顶层 和 底层 下标(上下边界)
	top, bottom := 0, len(matrix)-1
	// 列数  左右边界的下标（左右边界）
	left, right := 0, len(matrix[0])-1
	// 循环 重点搞清楚退出循环的条件是啥
	index := 0
	for top <= bottom && left <= right {
		// 左到右
		for i := left; i <= right; i++ {
			resultNum[index] = matrix[top][i]
			index++
		}
		// 更新 矩阵 规模（变小，遍历过的不用再遍历了）
		top++
		// 上到下
		for i := top; i <= bottom; i++ {
			resultNum[index] = matrix[i][right]
			index++
		}
		// 更新规模
		right--
		// 右到左
		if top <= bottom {
			for i := right; i >= left; i-- {
				resultNum[index] = matrix[bottom][i]
				index++
			}
			// 下临界更新
			bottom--
		}
		// 下到上
		if left <= right {
			for i := bottom; i >= top; i-- {
				resultNum[index] = matrix[i][left]
				index++
			}
			// 左临界更新
			left++
		}
	}
	return resultNum
}
