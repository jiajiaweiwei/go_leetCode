package main

import "fmt"

func main() {
	weight := hammingWeight(11)
	fmt.Println(weight)
}

// 位运算 简单题 位1的个数（求二进制中1的个数）
// 方法1 &1 然后移位操作，循环统计
func hammingWeight(n int) int {
	// 返回值
	resultNum := 0
	// &运算
	for i := 0; i < 64; i++ {
		if (n & 1) == 1 {
			resultNum++
		}
		n >>= 1
	}
	return resultNum
}
