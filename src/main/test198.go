package main

func main() {

}

// 动态规划 中等题 打家劫舍
func rob(nums []int) int {
	// 特殊情况处理
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return maxRob(nums[0], nums[1])
	}
	// 动态规划变量 同时也是返回值
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxRob(nums[0], nums[1])
	// 状态转移方程 dp[i] = max(dp[i-2]+num[i],dp[i-1])
	for i := 2; i < len(nums); i++ {
		dp[i] = maxRob(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func maxRob(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 优化时间复杂度
func rob2(nums []int) int {
	// 特殊情况处理
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return maxRob(nums[0], nums[1])
	}
	// 动态规划变量 同时也是返回值
	temp1 := nums[0]
	temp2 := maxRob(nums[0], nums[1])
	// 状态转移方程 dp[i] = max(dp[i-2]+num[i],dp[i-1])
	for i := 2; i < len(nums); i++ {
		temp := maxRob(temp1+nums[i], temp2)
		temp1 = temp2
		temp2 = temp
	}
	return temp2
}
