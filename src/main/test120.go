package main

import "fmt"

func main() {
	fmt.Println(-1 > -2)
}

// 多维动态规划 中等题 三角形最小路径和
// 方法1 错误，没有求每条路径的大小，只考虑了当前路径最小，但不是每条最小
func minimumTotal1(triangle [][]int) int {
	// 特殊情况处理
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	// 返回值
	resultNum := 0
	// 多维动态规划变量
	dp := make([]int, 0)
	// 初始化第一层
	dp = append(dp, triangle[0][0])
	// 状态转移方程 dp = append(dp,min(triangle[i][dp[i-1]],triangle[i][dp[i-1]+1]))
	// resultNum += min(triangle[i],triangle[i+1])
	minIndex := 0
	for i := 1; i < len(triangle); i++ {
		if triangle[i][dp[i-1]] < triangle[i][dp[i-1]+1] {
			resultNum += triangle[i][dp[i-1]]
			minIndex = dp[i-1]
		} else {
			minIndex = dp[i-1] + 1
			resultNum += triangle[i][dp[i-1]+1]
		}
		//状态转移方程
		dp = append(dp, minIndex)
	}
	return resultNum
}

func minimumTotal_my_true(triangle [][]int) int {
	// 特殊情况处理
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	// 初始化 dp 数组
	dp := make([]int, len(triangle))
	// 存放最小路径和
	dp[0] = triangle[0][0]
	// 自顶向下进行状态转移
	for i := 1; i < len(triangle); i++ {
		// 从右向左更新 dp 数组
		for j := i; j >= 0; j-- {
			if j == i {
				dp[j] = dp[j-1] + triangle[i][j]
			} else if j == 0 {
				dp[j] = dp[j] + triangle[i][j]
			} else {
				dp[j] = min120(dp[j-1], dp[j]) + triangle[i][j]
			}
		}
	}
	// 找到 dp 数组中的最小值
	resultNum := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] < resultNum {
			resultNum = dp[i]
		}
	}
	return resultNum
}

// 正确方法
func minimumTotal(triangle [][]int) int {
	// 边界检查
	if len(triangle) == 0 {
		return 0
	}
	// 初始化 dp 数组为三角形最后一层的元素
	dp := make([]int, len(triangle[len(triangle)-1]))
	copy(dp, triangle[len(triangle)-1])
	// 从倒数第二层开始，逐层向上计算最小路径和
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			// 状态转移方程
			dp[j] = triangle[i][j] + min120(dp[j], dp[j+1])
		}
	}
	// 返回三角形顶部的最小路径和
	return dp[0]
}
func min120(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 自顶向下
func minimumTotal2(triangle [][]int) int {
	// 边界检查
	if len(triangle) == 0 {
		return 0
	}
	// 初始化 dp 数组
	dp := make([]int, len(triangle))
	dp[0] = triangle[0][0]
	// 自顶向下进行状态转移
	for i := 1; i < len(triangle); i++ {
		// 从右向左更新 dp 数组
		for j := i; j >= 0; j-- {
			if j == i {
				dp[j] = dp[j-1] + triangle[i][j]
			} else if j == 0 {
				dp[j] = dp[j] + triangle[i][j]
			} else {
				dp[j] = min120(dp[j-1], dp[j]) + triangle[i][j]
			}
		}
	}
	// 找到 dp 数组中的最小值
	resultNum := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] < resultNum {
			resultNum = dp[i]
		}
	}
	return resultNum
}

// 自底向上 ***** 重点方法
func minimumTotal3(triangle [][]int) int {
	// 特殊情况处理 省
	// 动态规划变量 省略
	// 状态转移方程 triangle[i][j] = min(triangle[i+1][j],triangle[i+1][j+1]) + triangle[i][j]
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			// 从下往上逐个计算最短路径
			triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}
