package main

func main() {

}

// 多维动态规划 中等题 最小路径和 （方向只能 右 和 下 ）
// 同理120题
// 方法1 有问题 因为当判断要向上走还是向左走的时候可能遇见相同的值，这时候不知道哪边走才是最小路径，因为后面路径的值都不清楚
// 方法1 从终点开始找 缺点 只能找到右下角的最小路径 其他点找不到
func minPathSum(grid [][]int) int {
	// 返回值
	resultNum := 0
	// 状态转移方程 grid[i][j] =  min(grid[i-1][j],grid[i][j-1])+grid[i][j]
	i, j := len(grid)-1, len(grid[0])-1
	frequency := i + j
	for frequency > 0 {
		if i == 0 && j > 0 {
			// 只能向左或者向上走
			//循环往左走
			for j >= 0 {
				resultNum += grid[i][j]
				j--
			}
			return resultNum
		} else if j == 0 && i > 0 {
			// 只能向左或者向上走
			//循环往上走
			for i >= 0 {
				resultNum += grid[i][j]
				i--
			}
			return resultNum
		} else if i == 0 && j == 0 {
			resultNum += grid[0][0]
			return resultNum
		} else if grid[i-1][j] > grid[i][j-1] {
			resultNum += grid[i][j-1] // 向左走
			j--
		} else {
			resultNum += grid[i-1][j] //向上走
			i--
		}
		frequency--
	}
	return resultNum
}

// 找某个节点四周最小的节点 返回四周最小节点的下标 省 因为方向只能 右 下 所以不需要遍历四周
// func getMin64(grid [][]int, i, j int) (a, b int) {
//
//		return 1,2
//	}
//
// 方法2 正确方法  使用动态规划 正向找 并且可以顾及所有路径 从而找到真正的最小路径
func minPathSum2(grid [][]int) int {
	// 异常情况处理
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, columns := len(grid), len(grid[0])
	// 动态规划矩阵
	dp := make([][]int, rows)
	// 初始化动态矩阵空间
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, columns)
	}
	dp[0][0] = grid[0][0]
	// 初始化第一行
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	//初始化第一列
	for j := 1; j < columns; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	// 开始正序找路径
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			// 状态转移方程
			dp[i][j] = min64(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	// 返回结果
	return dp[rows-1][columns-1]
}

func min64(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 使用倒叙的正确方法 同理正序的动态规划
func minPathSum3(grid [][]int) int {
	// 获取网格的行数和列数
	m := len(grid)
	n := len(grid[0])
	// 从右下角开始倒序计算最小路径和
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 跳过右下角的初始点
			if i == m-1 && j == n-1 {
				continue
			}
			// 如果在最后一行，只能从右边来
			if i == m-1 {
				grid[i][j] += grid[i][j+1]
				// 如果在最后一列，只能从下边来
			} else if j == n-1 {
				grid[i][j] += grid[i+1][j]
				// 其他情况，取右边和下边路径中的最小值
			} else {
				// 状态转移方程
				grid[i][j] += min64(grid[i+1][j], grid[i][j+1])
			}
		}
	}
	// 返回左上角的值
	return grid[0][0]
}
