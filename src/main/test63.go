package main

func main() {

}

// 多维动态规划 中等题 不同路径2
// 方法1 动态规划 不使用滚动数组优化空间复杂度
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 获取行数和列数
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	// 如果起点或终点有障碍物，直接返回0
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	// 初始化动态规划矩阵
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 起点初始化
	dp[0][0] = 1
	// 填充动态规划矩阵
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果该点为障碍物
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				// 判断该点的上方
				if i > 0 {
					dp[i][j] += dp[i-1][j]
				}
				// 判断该点的左方
				if j > 0 {
					dp[i][j] += dp[i][j-1]
				}
			}
		}
	}
	// 返回终点的路径数量
	return dp[m-1][n-1]
}

// 方法1 动态规划 使用滚动数组优化空间复杂度 省略了上方路径数的存储，降低了空间复杂度一个维度
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	// 行数和列数
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	// 动态规划 滚动数组 大小和列数相同
	f := make([]int, m)
	// 初始化动态规划滚动数组 省略了上方路径数的存储，降低了空间复杂度一个维度
	if obstacleGrid[0][0] == 0 {
		f[0] = 1
	}
	// 遍历
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 如果该点有障碍物，置零跳过
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			// 没有障碍物就累加路径数
			if j-1 >= 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[m-1]
}
