package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("----------------")
	var tempNum = [][]int{
		{0, 9, 3, 3, 8, 2, 1, 4, 1, 7, 1, 2, 7},
		{6, 0, 2, 3, 3, 8, 5, 1, 9, 3, 2, 0, 7},
		{8, 4, 6, 0, 2, 6, 1, 5, 1, 0, 7, 2, 6},
		{1, 1, 9, 3, 9, 6, 5, 1, 1, 1, 1, 7, 2},
		{0, 0, 6, 3, 9, 4, 7, 5, 6, 0, 3, 7, 7},
		{5, 9, 7, 9, 6, 8, 1, 5, 3, 0, 3, 8, 3},
		{5, 1, 7, 4, 3, 9, 4, 9, 2, 6, 5, 0, 3},
	}
	setZeroes(tempNum)
	fmt.Println("----------------")
	fmt.Println("10" + "11")
}

// 矩阵 中等题 矩阵置零
// 方法1 暴力法 先遍历找到0的位置，然后第二次遍历根据条件置零
func setZeroes(matrix [][]int) {
	// 行数 列数
	row, column := len(matrix), len(matrix[0])
	// 记录所有0的位置
	zeroTemp := make([]string, 0)
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if matrix[i][j] == 0 {
				zeroTemp = append(zeroTemp, strconv.Itoa(i)+","+strconv.Itoa(j))
			}
		}
	}
	// 根据所有0的位置，置零
	for i := 0; i < len(zeroTemp); i++ {
		split := strings.Split(zeroTemp[i], ",")
		a, _ := strconv.Atoi(split[0])
		b, _ := strconv.Atoi(split[1])
		// 对a行 b列置零
		for m := 0; m < column; m++ {
			matrix[a][m] = 0
		}
		for n := 0; n < row; n++ {
			matrix[n][b] = 0
		}
	}
}

// 矩阵 中等题 矩阵置零
// 方法2 优化 空间复杂度
func setZeroes2(matrix [][]int) {
	row, column := len(matrix), len(matrix[0])
	zeroFirstRow := false
	zeroFirstCol := false
	// 检查第一行和第一列是否需要置零
	for j := 0; j < column; j++ {
		if matrix[0][j] == 0 {
			zeroFirstRow = true
			break
		}
	}
	for i := 0; i < row; i++ {
		if matrix[i][0] == 0 {
			zeroFirstCol = true
			break
		}
	}
	// 遍历矩阵，将需要置零的行和列的第一个元素置零
	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 根据第一行和第一列的标记，将需要置零的行和列置零
	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// 处理第一行和第一列
	if zeroFirstRow {
		for j := 0; j < column; j++ {
			matrix[0][j] = 0
		}
	}
	if zeroFirstCol {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}
}
