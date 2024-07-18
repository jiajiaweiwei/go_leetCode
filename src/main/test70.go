package main

func main() {

}

// 动态规划 简单题 爬楼梯
// 动态规划：状态转移方程
func climbStairs(n int) int {
	// 特殊情况处理
	if n == 2 {
		return 2
	}
	// 返回值
	resultNum := 1
	// 状态转移方程 climbStairs(n) = climbStairs(n-1) + climbStairs(n-2)
	// 动态规划变量
	a, b := 1, 2
	for i := 0; i < n-2; i++ {
		resultNum = a + b
		a = b
		b = resultNum
	}
	return resultNum
}
