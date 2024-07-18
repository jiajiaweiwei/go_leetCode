package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}

// 矩阵 中等题 有效的数独
// 方法1 直接遍历每一个单元格 （bord长和宽都是确定的9*9）
func isValidSudoku(board [][]byte) bool {
	// 开始遍历
	// 遍历81次
	for i := 0; i < 9; i++ { //行
		for j := 0; j < 9; j++ { //列
			// 1.判断该 行 该元素是否重复
			num := board[i][j]
			if num == '.' {
				continue
			}
			for m := 0; m < 9; m++ {
				// 其他位置发生重复
				if m != j && num == board[i][m] {
					return false
				}
			}
			// 2.判断该 列 该元素是否重复
			for n := 0; n < 9; n++ {
				// 其他位置发生重复
				if n != i && num == board[n][j] {
					return false
				}
			}
			// 3.判断该 3*3 该元素是否重复
			// 根据数学关系 计算出 每个小单元格对应的子矩阵的起点位置
			a, b := (i/3)*3, (j/3)*3
			// 从起点开始判断 子矩阵中 9个元素 是否有重复的
			if minMatrixBool(a, b, board, num, i, j) == false {
				return false
			}
		}
	}
	return true
}
func minMatrixBool(a, b int, board [][]byte, num byte, i, j int) bool {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x+a == i && y+b == j {
				continue
			}
			if num == board[x+a][y+b] {
				return false
			}
		}
	}
	return true
}
