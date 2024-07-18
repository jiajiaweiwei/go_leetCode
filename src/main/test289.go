package main

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(board)
}

// 矩阵 中等题 生命游戏
// 方法1 暴力法(屎山代码)
// 由于变化的元素可能会 影响 后变化的元素 所以 使用一个暂存矩阵 存储 元素的初始状态
func gameOfLife(board [][]int) {
	// 暂存矩阵
	// 创建 board 的副本
	// 这里的board是二维切片，所以，直接使用:=会导致tmepBoard和board底层引用的是同一个数组，所以要自己逐个复制
	tempBoard := make([][]int, len(board))
	for i := range board {
		tempBoard[i] = make([]int, len(board[i]))
		copy(tempBoard[i], board[i])
	}
	// 原始矩阵的行和列
	row, column := len(board), len(board[0])
	// 统某个节点周围八个位置活细胞的个数
	sumAround := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			// 获取统计值
			sumAround = getSumAround(tempBoard, i, j)
			// 根据统计值 判断 各个节点的值需要如何变化
			// 先判断是死细胞还是活细胞
			if tempBoard[i][j] == 0 {
				// 死细胞
				// 周围活细胞数sumAround = 3 个 	  死细胞复活（前提是死细胞）
				if sumAround == 3 {
					board[i][j] = 1
				}
			} else {
				// 活细胞
				// 周围活细胞数sumAround < 2 个 	  死亡置零
				// 周围活细胞数sumAround > 3 个 	  死亡置零
				if sumAround < 2 || sumAround > 3 {
					board[i][j] = 0
				}
				// 周围活细胞数sumAround = 2 或 3 个 细胞继续存活
			}
		}
	}
}

// 统计某元素周围八个元素存活个数 函数
func getSumAround(tempBoard [][]int, i, j int) int {
	// 获取临时矩阵 的 行数和列数
	row, colum := len(tempBoard), len(tempBoard[0])
	// 返回值
	resultNum := 0
	// 开始统计
	// 某节点i，j周围的节点
	// 先判断 下标是否存在
	// [i-1][j-1]
	if 0 <= i-1 && i-1 <= row-1 && 0 <= j-1 && j-1 <= colum-1 {
		resultNum += tempBoard[i-1][j-1]
	}
	//[i-1][j]
	if 0 <= i-1 && i-1 <= row-1 && 0 <= j && j <= colum-1 {
		resultNum += tempBoard[i-1][j]
	}
	// [i-1][j+1]
	if 0 <= i-1 && i-1 <= row-1 && 0 <= j+1 && j+1 <= colum-1 {
		resultNum += tempBoard[i-1][j+1]
	}
	// [i][j-1]
	if 0 <= i && i <= row-1 && 0 <= j-1 && j-1 <= colum-1 {
		resultNum += tempBoard[i][j-1]
	}
	// [i][j+1]
	if 0 <= i && i <= row-1 && 0 <= j+1 && j+1 <= colum-1 {
		resultNum += tempBoard[i][j+1]
	}
	// [i+1][j-1]
	if 0 <= i+1 && i+1 <= row-1 && 0 <= j-1 && j-1 <= colum-1 {
		resultNum += tempBoard[i+1][j-1]
	}
	// [i+1][j]
	if 0 <= i+1 && i+1 <= row-1 && 0 <= j && j <= colum-1 {
		resultNum += tempBoard[i+1][j]
	}
	// [i+1][j+1]
	if 0 <= i+1 && i+1 <= row-1 && 0 <= j+1 && j+1 <= colum-1 {
		resultNum += tempBoard[i+1][j+1]
	}
	return resultNum
}

// 方法2 使用 移位操作 简化空间复杂度 重点学习|=2的作用：将次地位置1，记录下一阶段的状态，同时&1也可以根据最低为状态得出初始状态是0还是1
func gameOfLife2(board [][]int) {
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return
	}
	rows := len(board)
	cols := len(board[0])
	// 第一轮遍历：计算每个细胞下一时刻状态
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			liveNum := getSumAround2(board, row, col)
			// 这一行是重点 为什么不直接使用 == 1来判断而是用 &
			if board[row][col]&1 == 1 { // 规则1、2、3，即判断活细胞下一时刻状态
				if liveNum == 2 || liveNum == 3 {
					board[row][col] |= 2 // 表示将次低位置1
				}
			} else { // 规则4，即判断死细胞下一时刻状态
				if liveNum == 3 {
					board[row][col] |= 2
				}
			}
		}
	}
	// 第二轮遍历：变更每个细胞的状态
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			board[row][col] >>= 1
		}
	}
}

// 统计某元素周围八个元素存活个数 函数 使用移位操作 优化时间复杂度
func getSumAround2(board [][]int, i, j int) int {
	rows := len(board)
	cols := len(board[0])
	resultNum := 0
	// 定义方向数组
	directions := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	for _, direction := range directions {
		x, y := i+direction[0], j+direction[1]
		if x >= 0 && x < rows && y >= 0 && y < cols {
			// 使用& 判断 2&1 = 0 说明原来是
			resultNum += board[x][y] & 1
		}
	}
	return resultNum
}
