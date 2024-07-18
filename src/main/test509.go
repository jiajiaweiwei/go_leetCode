package main

func main() {

}

// 动态规划 入门简单题 求第n个斐波那契数
func fib(n int) int {
	// 特殊情况处理
	if n == 1 {
		return 0
	}
	// 返回值
	resultNum := 0
	// 状态转移方程 fib(n) = fib(n-2)+fib(n-1)
	// 动态规划变量
	a, b := 0, 1
	for i := 0; i < n-1; i++ {
		resultNum = a + b
		a = b
		b = resultNum
	}

	return resultNum
}
