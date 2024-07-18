package main

func main() {

}

// 一维动态规划 中等题 最长子序列(不需要连续)
// 方法1 暴力法
func lengthOfLIS(nums []int) int {
	// 特殊情况处理
	if len(nums) == 0 {
		return 0
	}
	// 返回值
	resultNum := 0
	// 动态规划变量
	dp := make([]int, len(nums)+1)
	dp[0] = 1
	// 转态转移方程 dp[i] = max(dp[j])+1,其中j<i
	for i := 1; i <= len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// 状态转移方程
				dp[i] = max300(dp[i], dp[j]+1)
			}
		}
		resultNum = max300(resultNum, dp[i])
	}
	return resultNum
}

func max300(a, b int) int {
	if a > b {
		return a
	}
	return b
}
