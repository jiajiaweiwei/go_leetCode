package main

import "math"

func main() {

}

// 一维动态规划 中等题 完全平凡数
func numSquares(n int) int {
	// 动态规划变量
	dp := make([]int, n+1)

	// 状态转移方程 dp[n] = min(minn,dp[n-k^2])+1
	for i := 1; i <= n; i++ {
		minTemp := math.MaxInt64
		for j := 0; j*j <= i; j++ {
			minTemp = min279(minTemp, dp[i-j*j])
		}
		dp[i] = minTemp + 1
	}

	return dp[n]
}
func min279(a, b int) int {
	if a > b {
		return b
	}
	return a
}
