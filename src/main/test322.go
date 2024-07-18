package main

func main() {

}

// 一维动态规划 中等题 零钱兑换
// 方法1 动态规划
func coinChange(coins []int, amount int) int {
	// 初始化 dp 数组，长度为 amount + 1，初始值为 amount + 1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	// 动态规划，计算最小硬币数量
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				// 状态转移方程
				dp[i] = min322(dp[i], dp[i-coin]+1)
			}
		}
	}
	// 如果 dp[amount] 还是初始值，说明无法凑出 amount，返回 -1
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min322(a, b int) int {
	if a > b {
		return b
	}
	return a
}
