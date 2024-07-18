package main

import (
	"fmt"
)

func main() {
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i, v := range matrix {
		fmt.Println(i, "----")
		fmt.Println(v, "----")
	}
}

// 数学 困难题 直线上最多的点数 求最多有多少个点在同一条直线上
//
//	寄了 放弃
func maxPoints(points [][]int) int {
	// 特殊情况处理
	if len(points) == 0 {
		return 0
	}
	if len(points) == 1 {
		return 1
	}
	// 存储斜率
	mostPoints := 0
	// 记录斜率
	//tempMap := make(map[float64]int)

	return mostPoints
}

func max149(a, b int) int {
	if a > b {
		return a
	}
	return b
}
